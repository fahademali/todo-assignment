package activities

import (
	"context"
	"fmt"
	"net/url"
	"time"
	"user_service/log"

	"worker/config"
	"worker/models"
	"worker/utils"
)

func GetTodosDueTodayActivity(ctx context.Context) (models.TodosDueTodayResponse, error) {
	var todosDueToday models.TodosDueTodayResponse

	now := time.Now()

	params := url.Values{}

	params.Add("start_date", now.Format(models.DATE_FORMAT))
	params.Add("end_date", now.Format(models.DATE_FORMAT))

	url, _ := url.ParseRequestURI(config.AppConfig.BASEURL_TODO_SERVICE)
	url.Path = models.TODO_RESOURCE
	url.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", url)

	todosDueToday, err := utils.MakeGetReq[models.TodosDueTodayResponse](urlStr)
	if err != nil {
		return todosDueToday, err
	}

	log.GetLog().Info("todosDueToday")
	log.GetLog().Info(todosDueToday)

	return todosDueToday, nil
}
