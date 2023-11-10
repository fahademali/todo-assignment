package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"worker/config"
	"worker/models"
)

func GetTodosDueTodayActivity(ctx context.Context) (models.TodosDueTodaySuccessResponse, error) {
	fmt.Println(",,,,,,,,,,,,,,,activity,,,,,,,,,,,,,,,")
	var todosDueToday models.TodosDueTodaySuccessResponse
	var errResponse models.ErrorResponse

	now := time.Now()
	resource := "/internal/todos"
	dateFormat := "2006-01-02"
	params := url.Values{}

	params.Add("due_date", now.Format(dateFormat))

	u, _ := url.ParseRequestURI(config.AppConfig.BASEURL_TODO_SERVICE)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resp, err := http.Get(urlStr)
	if err != nil {
		return todosDueToday, err
	}

	if resp.StatusCode != http.StatusOK {
		if err = json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			return todosDueToday, err
		}
		return todosDueToday, fmt.Errorf(errResponse.Error)
	}

	if err = json.NewDecoder(resp.Body).Decode(&todosDueToday); err != nil {
		return todosDueToday, err
	}

	return todosDueToday, nil
}
