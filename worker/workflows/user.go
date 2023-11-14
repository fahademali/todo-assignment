package workflows

import (
	"fmt"
	"strings"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"

	"worker/activities"
	"worker/log"
	"worker/models"
	"worker/utils"
)

func RemindUsersForDueDateWorkflow(ctx workflow.Context) error {
	var userIDs []int64

	activityOptions := utils.GetDefaultActivityOptions()
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	var todosDueTodayResponse models.TodosDueTodayResponse
	todosDueDateFuture := workflow.ExecuteActivity(ctx, activities.GetTodosDueTodayActivity)
	if err := todosDueDateFuture.Get(ctx, &todosDueTodayResponse); err != nil {
		return err
	}

	for _, userTodos := range todosDueTodayResponse.Data {
		userIDs = append(userIDs, userTodos.UserID)
	}

	log.GetLog().Info(userIDs)
	var userEmailsResponse models.GetEmailsByIDResponse
	userEmailsFuture := workflow.ExecuteActivity(ctx, activities.GetEmailsByUserIDActivity, userIDs)
	if err := userEmailsFuture.Get(ctx, &userEmailsResponse); err != nil {
		return err
	}

	var userMap = make(map[int64]string)

	for _, user := range userEmailsResponse.Data {
		userMap[user.ID] = user.Email
	}

	var recepientDetails []models.RecepietDetails

	for _, todo := range todosDueTodayResponse.Data {
		emailBody := fmt.Sprintf(models.EMAIL_BODY_TEMPLATE, strings.Join(todo.TodoTitles, ","))
		emailSubject := fmt.Sprintf(models.EMAIL_SUBJECT_TEMPLATE, len(todo.TodoTitles))
		recepientDetails = append(recepientDetails, models.RecepietDetails{UserEmail: userMap[todo.UserID], Body: emailBody, Subject: emailSubject})
	}

	var sendEmailResponse models.SendEmailResponse
	sendEmailFuture := workflow.ExecuteActivity(ctx, activities.SendEmailActivity, recepientDetails)
	if err := sendEmailFuture.Get(ctx, &sendEmailResponse); err != nil {
		return err
	}

	workflow.GetLogger(ctx).Info("Done", zap.String("result", "result"))
	return nil
}
