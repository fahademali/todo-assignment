package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/cadence/client"
	"go.uber.org/cadence/workflow"

	"worker/config"
	"worker/models"
)

func MakePOSTReq[E interface{}](payloadBytes []byte, url string) (E, error) {
	var successResponse E
	var errResponse models.ErrorResponse

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return successResponse, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return successResponse, err
	}
	defer response.Body.Close()

	//TODO: If the route is not found e.g sending post instead of get do inform the user properly

	if response.StatusCode < 200 || response.StatusCode > 299 {
		if err = json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			return successResponse, err
		}
		return successResponse, fmt.Errorf(errResponse.Error)
	}

	if err = json.NewDecoder(response.Body).Decode(&successResponse); err != nil {
		return successResponse, err
	}

	return successResponse, nil
}

func MakeGetReq[E interface{}](url string) (E, error) {
	var successResponse E
	var errResponse models.ErrorResponse

	response, err := http.Get(url)
	if err != nil {
		return successResponse, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		if err = json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			return successResponse, err
		}
		return successResponse, fmt.Errorf(errResponse.Error)
	}

	if err = json.NewDecoder(response.Body).Decode(&successResponse); err != nil {
		return successResponse, err
	}

	return successResponse, nil
}

func GetDefaultActivityOptions() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		TaskList:               config.AppConfig.TASK_LIST_NAME,
		ScheduleToCloseTimeout: time.Second * 60,
		ScheduleToStartTimeout: time.Second * 60,
		StartToCloseTimeout:    time.Second * 60,
		HeartbeatTimeout:       time.Second * 10,
		WaitForCancellation:    false,
	}
}

func GetDefaultBeginWorkflowOptions(cronSchedule string) client.StartWorkflowOptions {
	return client.StartWorkflowOptions{
		TaskList:                     config.AppConfig.TASK_LIST_NAME,
		ExecutionStartToCloseTimeout: 10 * time.Second,
		// CronSchedule:                 cronSchedule,
	}
}
