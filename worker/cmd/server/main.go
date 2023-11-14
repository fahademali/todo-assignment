package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"worker/activities"
	"worker/config"
	"worker/workflows"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	startWorker(buildLogger(), buildCadenceClient())

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
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(config.AppConfig.CLIENT_NAME))
	if err != nil {
		panic("Failed to setup tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: config.AppConfig.CLIENT_NAME,
		Outbounds: yarpc.Outbounds{
			config.AppConfig.CADENCE_SERVICE: {Unary: ch.NewSingleOutbound(config.AppConfig.HOST_PORT)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(config.AppConfig.CADENCE_SERVICE))
}

func startWorker(logger *zap.Logger, service workflowserviceclient.Interface) {
	workerOptions := worker.Options{
		Logger:       logger,
		MetricsScope: tally.NewTestScope(config.AppConfig.TASK_LIST_NAME, map[string]string{}),
	}

	worker := worker.New(
		service,
		config.AppConfig.DOMAIN,
		config.AppConfig.TASK_LIST_NAME,
		workerOptions)

	worker.RegisterWorkflow(workflows.RemindUsersForDueDateWorkflow)
	worker.RegisterActivity(activities.GetTodosDueTodayActivity)
	worker.RegisterActivity(activities.GetEmailsByUserIDActivity)
	worker.RegisterActivity(activities.SendEmailActivity)
	err := worker.Start()
	if err != nil {
		panic("Failed to start worker")
	}

	logger.Info("Started Worker.", zap.String("worker", config.AppConfig.TASK_LIST_NAME))
}
