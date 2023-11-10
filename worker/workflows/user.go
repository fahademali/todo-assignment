package workflows

import (
	"fmt"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"

	"worker/activities"
	"worker/log"
	"worker/models"
	"worker/utils"
)

func RemindUsersForDueDateWorkflow(ctx workflow.Context) error {
	fmt.Println(",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,")
	emailSubject := `[Important] Todo Due Today`
	emailBody := `
	<html>
		<body>
			<p>Hello!</p>
			<p>You have pending todos which are due today open your app and get it down. No More procrastination</p>
		</body>
	</html>
	`

	activityOptions := utils.GetDefaultActivityOptions()
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	var todosDueTodayResponse models.TodosDueTodaySuccessResponse
	todosDueDateFuture := workflow.ExecuteActivity(ctx, activities.GetTodosDueTodayActivity)
	if err := todosDueDateFuture.Get(ctx, &todosDueTodayResponse); err != nil {
		return err
	}
	log.GetLog().Warn(todosDueTodayResponse)

	var userEmailsResponse models.GetEmailsByIDSuccessResponse
	userEmailsFuture := workflow.ExecuteActivity(ctx, activities.GetEmailsByUserIDActivity, todosDueTodayResponse)
	if err := userEmailsFuture.Get(ctx, &userEmailsResponse); err != nil {
		return err
	}
	log.GetLog().Warn(userEmailsResponse)

	var sendEmailResponse models.SendEmailSuccessResponse
	sendEmailFuture := workflow.ExecuteActivity(ctx, activities.SendEmailActivity, userEmailsResponse, emailSubject, emailBody)
	if err := sendEmailFuture.Get(ctx, &sendEmailResponse); err != nil {
		return err
	}

	log.GetLog().Info(sendEmailResponse)

	workflow.GetLogger(ctx).Info("Done", zap.String("result", "result"))
	return nil
}
