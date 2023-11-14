package activities

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"worker/config"
	"worker/models"
	"worker/utils"
)

func GetTodosDueTodayActivity(ctx context.Context) (models.TodosDueTodayResponse, error) {
	now := time.Now()

	params := url.Values{}

	params.Add("start_date", now.Format(models.DATE_FORMAT))
	params.Add("end_date", now.Format(models.DATE_FORMAT))

	url, _ := url.ParseRequestURI(config.AppConfig.BASEURL_TODO_SERVICE)
	url.Path = models.TODO_RESOURCE
	url.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", url)

	return utils.MakeGetReq[models.TodosDueTodayResponse](urlStr)
}
