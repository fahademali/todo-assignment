package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo_service/log"

	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"

	"github.com/uber-go/tally"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var HostPort = "127.0.0.1:7933"
var Domain = "test-domain"
var TaskListName = "test-list"
var ClientName = "test-client"
var CadenceService = "cadence-frontend"

func main() {
	startWorker(buildLogger(), buildCadenceClient())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start your application
	startWorker(buildLogger(), buildCadenceClient())

	// Wait for termination signals
	<-quit
}

func buildLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.InfoLevel)

	var err error
	logger, err := config.Build()
	if err != nil {
		panic("Failed to setup logger")
	}

	return logger
}

func buildCadenceClient() workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
	if err != nil {
		panic("Failed to setup tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			CadenceService: {Unary: ch.NewSingleOutbound(HostPort)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(CadenceService))
}

func SimpleWorkFlow(ctx workflow.Context) error {
	fmt.Println(",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,")
	ao := workflow.ActivityOptions{
		TaskList:               TaskListName,
		ScheduleToCloseTimeout: time.Second * 60,
		ScheduleToStartTimeout: time.Second * 60,
		StartToCloseTimeout:    time.Second * 60,
		HeartbeatTimeout:       time.Second * 10,
		WaitForCancellation:    false,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	future := workflow.ExecuteActivity(ctx, SimpleActivity)
	var result string
	if err := future.Get(ctx, &result); err != nil {
		return err
	}
	workflow.GetLogger(ctx).Info("Done", zap.String("result", result))
	return nil
}

type todoDueTodaySuccessResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    []int64 `json:"data"`
}
type GetRemailByUserIDSuccessResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}
type SendEmailSuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
type SuccessResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func SimpleActivity(ctx context.Context) error {
	fmt.Println(",,,,,,,,,,,,,,,activity,,,,,,,,,,,,,,,")

	now := time.Now()
	baseURL := "http://localhost:8081"
	resource := "/admin/todos"
	dateFormat := "2006-01-02"
	params := url.Values{}

	params.Add("due_date", now.Format(dateFormat))

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resp, err := http.Get(urlStr)
	if err != nil {
		log.GetLog().Error(err.Error())
		return err
	}

	var data todoDueTodaySuccessResponse
	var errResponse ErrorResponse

	if resp.StatusCode != http.StatusOK {
		if err = json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			return err
		}
		log.GetLog().Error(errResponse.Error)
		return fmt.Errorf(errResponse.Error)
	}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	// userIDs := ConvertSlice[float64](data.Data)

	payload := map[string][]int64{
		"user_ids": data.Data,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	emailResponse, err := makeHttpReq[GetRemailByUserIDSuccessResponse](payloadBytes, "http://localhost:8082/users")
	if err != nil {
		return err
	}

	// userEmailAddresses := ConvertSlice[string](successResponse.Data)

	payloadv2 := map[string][]string{
		"userEmailAddresses": emailResponse.Data,
	}
	payloadBytes, err = json.Marshal(payloadv2)
	if err != nil {
		return err
	}

	successResponse, err := makeHttpReq[SendEmailSuccessResponse](payloadBytes, "http://localhost:8082/send-email")
	if err != nil {
		return err
	}

	log.GetLog().Info(successResponse.Data)
	return nil
}

// func ConvertSlice[E any](in []any) (out []E) {
// 	out = make([]E, 0, len(in))
// 	for _, v := range in {
// 		out = append(out, v.(E))
// 	}
// 	return out
// }

func makeHttpReq[E interface{}](payloadBytes []byte, url string) (E, error) {
	var successResponse E
	var errResponse ErrorResponse

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return successResponse, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return successResponse, err
	}
	defer resp.Body.Close()

	log.GetLog().Info(resp.StatusCode)
	log.GetLog().Info(resp.Body)

	//TODO: If the route is not found like prev sending post instead of get do inform the user properly
	if resp.StatusCode != http.StatusOK {
		if err = json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			return successResponse, err
		}
		log.GetLog().Error(errResponse.Error)
		return successResponse, fmt.Errorf(errResponse.Error)
	}

	if err = json.NewDecoder(resp.Body).Decode(&successResponse); err != nil {
		return successResponse, err
	}

	return successResponse, nil
}

func startWorker(logger *zap.Logger, service workflowserviceclient.Interface) {
	// TaskListName identifies set of client workflows, activities, and workers.
	// It could be your group or client or application name.
	workerOptions := worker.Options{
		Logger:       logger,
		MetricsScope: tally.NewTestScope(TaskListName, map[string]string{}),
	}

	worker := worker.New(
		service,
		Domain,
		TaskListName,
		workerOptions)

	worker.RegisterWorkflow(SimpleWorkFlow)
	worker.RegisterActivity(SimpleActivity)
	err := worker.Start()
	if err != nil {
		panic("Failed to start worker")
	}

	logger.Info("Started Worker.", zap.String("worker", TaskListName))
}
