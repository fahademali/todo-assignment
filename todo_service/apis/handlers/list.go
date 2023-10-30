package handlers

import (
	"net/http"
	httpResponse "todo_service/http"
	"todo_service/models"
	"todo_service/services"

	"github.com/gin-gonic/gin"
)

type IListHandlers interface {
	HandleCreateList(ctx *gin.Context)
	HandleDeleteList(ctx *gin.Context)
	// ASK: shoudl i keep it ListName in below or should i remove list reason of keeping is updaing also means adding a todo.
	HandleUpdateListName(ctx *gin.Context)
}

type ListHandlers struct {
	listService services.IListService
}

func NewListHandlers(listService services.IListService) IListHandlers {
	return &ListHandlers{listService: listService}
}

func (lh *ListHandlers) HandleCreateList(ctx *gin.Context) {
	var requestBody models.CreateListRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	if err := lh.listService.CreateList(requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("list has been created"))
}

func (lh *ListHandlers) HandleDeleteList(ctx *gin.Context) {
	listID := ctx.Param("listID")

	if err := lh.listService.DeleteList(listID); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("list has been deleted"))
}

func (lh *ListHandlers) HandleUpdateListName(ctx *gin.Context) {
	var requestBody models.UpdateListRequest
	listID := ctx.Param("listID")

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := lh.listService.UpdateList(listID, requestBody.Name); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("list has been updated"))
}
