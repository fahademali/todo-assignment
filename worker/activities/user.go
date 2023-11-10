package activities

import (
	"context"
	"encoding/json"
	"fmt"

	"worker/config"
	"worker/models"
	"worker/utils"
)

func GetEmailsByUserIDActivity(ctx context.Context, userIDs models.TodosDueTodaySuccessResponse) (models.GetEmailsByIDSuccessResponse, error) {
	var emailResponse models.GetEmailsByIDSuccessResponse
	var getEmailsByIDUrl = fmt.Sprintf("%s/internal/users", config.AppConfig.BASEURL_USER_SERVICE)

	payload := map[string][]int64{
		"user_ids": userIDs.Data,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return emailResponse, err
	}

	emailResponse, err = utils.MakeHttpReq[models.GetEmailsByIDSuccessResponse](payloadBytes, getEmailsByIDUrl)
	if err != nil {
		return emailResponse, err
	}

	return emailResponse, nil
}

func SendEmailActivity(ctx context.Context, emailResponse models.GetEmailsByIDSuccessResponse, subject string, body string) (models.SendEmailSuccessResponse, error) {
	var emailSuccessResponse models.SendEmailSuccessResponse
	var sendEmailUrl = fmt.Sprintf("%s/internal/send-email", config.AppConfig.BASEURL_USER_SERVICE)

	payload := models.SendEmailPayload{
		UserEmailAddresses: emailResponse.Data,
		Subject:            subject,
		Body:               body,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return emailSuccessResponse, err
	}

	emailSuccessResponse, err = utils.MakeHttpReq[models.SendEmailSuccessResponse](payloadBytes, sendEmailUrl)
	if err != nil {
		return emailSuccessResponse, err
	}

	return emailSuccessResponse, nil
}
