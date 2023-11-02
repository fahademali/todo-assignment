package handlers

import (
	"net/http"
	"strconv"
	httpResponse "todo_service/http"
	"todo_service/models"
	"todo_service/services"

	"github.com/gin-gonic/gin"
)

type IListHandlers interface {
	HandleCreateList(ctx *gin.Context)
	HandleDeleteList(ctx *gin.Context)
	HandleUpdateList(ctx *gin.Context)
}

type ListHandlers struct {
	listService services.IListService
}

func NewListHandlers(listService services.IListService) IListHandlers {
	return &ListHandlers{listService: listService}
}

func (lh *ListHandlers) HandleCreateList(ctx *gin.Context) {
	var requestBody models.CreateListRequest
	user := ctx.MustGet("user").(models.UserInfo)

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	if err := lh.listService.CreateList(requestBody, user.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("list has been created"))
}

func (lh *ListHandlers) HandleDeleteList(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.UserInfo)
	listIDStr := ctx.Param("list_id")
	listID, err := strconv.ParseInt(listIDStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	//TODO: implemente with CASCADE
	if err := lh.listService.DeleteList(listID, ctx, user.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("list has been deleted"))
}

func (lh *ListHandlers) HandleUpdateList(ctx *gin.Context) {
	var requestBody models.UpdateListRequest
	user := ctx.MustGet("user").(models.UserInfo)

	listIDStr := ctx.Param("list_id")
	listID, err := strconv.ParseInt(listIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := lh.listService.UpdateList(listID, user.ID, requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("list has been updated"))
}
