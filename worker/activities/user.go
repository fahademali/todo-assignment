package activities

import (
	"context"
	"encoding/json"
	"fmt"

	"worker/config"
	"worker/log"
	"worker/models"
	"worker/utils"
)

// TODO: either pass the todoIDs only or name it something different
func GetEmailsByUserIDActivity(ctx context.Context, userIDs []int64) (models.GetEmailsByIDSuccessResponse, error) {
	fmt.Println(",,,,,,,,,,,,,,,running GetEmailsByUserIDActivity,,,,,,,,,,,,,,,")

	var emailResponse models.GetEmailsByIDSuccessResponse
	var getEmailsByIDUrl = fmt.Sprintf("%s/internal/users", config.AppConfig.BASEURL_USER_SERVICE)

	payload := map[string][]int64{
		"userIDs": userIDs,
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

func SendEmailActivity(ctx context.Context, recepientDetails []models.RecepietDetails) (models.SendEmailSuccessResponse, error) {
	fmt.Println(",,,,,,,,,,,,,,,running SendEmailActivity,,,,,,,,,,,,,,,")
	var emailSuccessResponse models.SendEmailSuccessResponse
	var sendEmailUrl = fmt.Sprintf("%s/internal/send-email", config.AppConfig.BASEURL_USER_SERVICE)

	payload := models.SendEmailPayload{
		RecepietDetails: recepientDetails,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return emailSuccessResponse, err
	}

	emailSuccessResponse, err = utils.MakeHttpReq[models.SendEmailSuccessResponse](payloadBytes, sendEmailUrl)
	if err != nil {
		return emailSuccessResponse, err
	}

	log.GetLog().Info("email sent success fully ")
	log.GetLog().Info(emailSuccessResponse)

	return emailSuccessResponse, nil
}
