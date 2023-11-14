package activities

import (
	"context"
	"encoding/json"
	"fmt"

	"worker/config"
	"worker/models"
	"worker/utils"
)

func GetEmailsByUserIDActivity(ctx context.Context, userIDs []int64) (models.GetEmailsByIDResponse, error) {
	var emailResponse models.GetEmailsByIDResponse
	var getEmailsByIDUrl = fmt.Sprintf("%s/internal/users", config.AppConfig.BASEURL_USER_SERVICE)

	payload := models.GetEmailByUserIDPayload{
		UserIDs: userIDs,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return emailResponse, err
	}

	return utils.MakePOSTReq[models.GetEmailsByIDResponse](payloadBytes, getEmailsByIDUrl)
}

func SendEmailActivity(ctx context.Context, recepientDetails []models.RecepietDetails) (models.SendEmailResponse, error) {
	var emailSuccessResponse models.SendEmailResponse
	var sendEmailUrl = fmt.Sprintf("%s/internal/send-email", config.AppConfig.BASEURL_USER_SERVICE)

	payload := models.SendEmailPayload{
		RecepietDetails: recepientDetails,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return emailSuccessResponse, err
	}

	emailSuccessResponse, err = utils.MakePOSTReq[models.SendEmailResponse](payloadBytes, sendEmailUrl)
	if err != nil {
		return emailSuccessResponse, err
	}

	return emailSuccessResponse, nil
}
