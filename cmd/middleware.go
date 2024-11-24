package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-author-service/constants"
	"github.com/hilmiikhsan/library-author-service/helpers"
)

func (d *Dependency) MiddlewareValidateToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get(constants.HeaderAuthorization)
	if authHeader == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateToken - authorization empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthorizationIsEmpty))
		ctx.Abort()
		return
	}

	token := helpers.ExtractBearerToken(authHeader)
	if token == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateToken - invalid bearer token format")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrInvalidAuthorizationFormat))
		ctx.Abort()
		return
	}

	tokenData, err := d.External.ValidateToken(ctx.Request.Context(), token)
	if err != nil {
		helpers.Logger.Error("middleware::MiddlewareValidateToken - failed to validate token", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrInvalidAuthorization))
		ctx.Abort()
		return
	}

	if tokenData.Role != constants.AuthRoleAdmin {
		helpers.Logger.Error("middleware::MiddlewareValidateToken - invalid role you do not permission to access this endpoint")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthRolePermission))
		ctx.Abort()
		return
	}

	ctx.Set(constants.TokenTypeAccess, tokenData)

	ctx.Next()
}
