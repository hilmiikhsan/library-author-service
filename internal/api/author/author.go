package author

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-author-service/constants"
	"github.com/hilmiikhsan/library-author-service/helpers"
	"github.com/hilmiikhsan/library-author-service/internal/dto"
	"github.com/hilmiikhsan/library-author-service/internal/interfaces"
	"github.com/hilmiikhsan/library-author-service/internal/validator"
)

type AuthorHandler struct {
	AuthorService interfaces.IAuthorService
	Validator     *validator.Validator
}

func (api *AuthorHandler) CreateAuthor(ctx *gin.Context) {
	var (
		req = new(dto.CreateAuthorRequest)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.Logger.Error("handler::CreateAuthor - Failed to bind request : ", err)
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		helpers.Logger.Error("handler::CreateAuthor - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	err := api.AuthorService.CreateAuthor(ctx.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrInvalidFormatDate) {
			helpers.Logger.Error("handler::CreateAuthor - Invalid format date")
			ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrInvalidFormatDate))
			return
		}

		if strings.Contains(err.Error(), constants.ErrAuthorAlreadyExist) {
			helpers.Logger.Error("handler::CreateAuthor - Author already exist")
			ctx.JSON(http.StatusConflict, helpers.Error(constants.ErrAuthorAlreadyExist))
			return
		}

		helpers.Logger.Error("handler::CreateAuthor - Failed to create Author : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Success(nil, ""))
}

func (api *AuthorHandler) GetDetailAuthor(ctx *gin.Context) {
	var (
		id = ctx.Param("id")
	)

	if id == "" {
		helpers.Logger.Error("handler::GetDetailAuthor - Missing required parameter: id")
		ctx.JSON(http.StatusBadRequest, helpers.Error("missing required parameter: id"))
		return
	}

	if !helpers.IsValidUUID(id) {
		helpers.Logger.Error("handler::GetDetailAuthor - Invalid UUID format for parameter: id")
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrParamIdIsRequired))
		return
	}

	res, err := api.AuthorService.GetDetailAuthor(ctx.Request.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrAuthorNotFound) {
			helpers.Logger.Error("handler::GetDetailAuthor - Author not found")
			ctx.JSON(http.StatusNotFound, helpers.Error(constants.ErrAuthorNotFound))
			return
		}

		helpers.Logger.Error("handler::GetDetailAuthor - Failed to get Author detail : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}

func (api *AuthorHandler) GetListAuthor(ctx *gin.Context) {
	pageIndexStr := ctx.Query("page")
	pageSizeStr := ctx.Query("limit")

	pageIndex, _ := strconv.Atoi(pageIndexStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if pageIndex <= 0 {
		pageIndex = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	res, err := api.AuthorService.GetListAuthor(ctx.Request.Context(), pageSize, pageIndex)
	if err != nil {
		helpers.Logger.Error("handler::GetListAuthor - Failed to get list Author : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}

func (api *AuthorHandler) UpdateAuthor(ctx *gin.Context) {
	var (
		req = new(dto.UpdateAuthorRequest)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.Logger.Error("handler::UpdateAuthor - Failed to bind request : ", err)
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		helpers.Logger.Error("handler::UpdateAuthor - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	if !helpers.IsValidUUID(req.ID) {
		helpers.Logger.Error("handler::GetDetailAuthor - Invalid UUID format for parameter: id")
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrIdIsNotValidUUID))
		return
	}

	err := api.AuthorService.UpdateAuthor(ctx.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrInvalidFormatDate) {
			helpers.Logger.Error("handler::CreateAuthor - Invalid format date")
			ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrInvalidFormatDate))
			return
		}

		if strings.Contains(err.Error(), constants.ErrAuthorNotFound) {
			helpers.Logger.Error("handler::UpdateAuthor - Author not found")
			ctx.JSON(http.StatusNotFound, helpers.Error(constants.ErrAuthorNotFound))
			return
		}

		helpers.Logger.Error("handler::UpdateAuthor - Failed to update Author : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(nil, ""))
}

func (api *AuthorHandler) DeleteAuthor(ctx *gin.Context) {
	var (
		id = ctx.Param("id")
	)

	if id == "" {
		helpers.Logger.Error("handler::DeleteAuthor - Missing required parameter: id")
		ctx.JSON(http.StatusBadRequest, helpers.Error("missing required parameter: id"))
		return
	}

	if !helpers.IsValidUUID(id) {
		helpers.Logger.Error("handler::DeleteAuthor - Invalid UUID format for parameter: id")
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrIdIsNotValidUUID))
		return
	}

	err := api.AuthorService.DeleteAuthor(ctx.Request.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrAuthorNotFound) {
			helpers.Logger.Error("handler::DeleteAuthor - Author not found")
			ctx.JSON(http.StatusNotFound, helpers.Error(constants.ErrAuthorNotFound))
			return
		}

		helpers.Logger.Error("handler::DeleteAuthor - Failed to delete Author : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(nil, ""))
}
