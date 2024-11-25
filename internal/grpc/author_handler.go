package grpc

import (
	"context"
	"strings"

	"github.com/hilmiikhsan/library-author-service/cmd/proto/author"
	"github.com/hilmiikhsan/library-author-service/constants"
	"github.com/hilmiikhsan/library-author-service/helpers"
	"github.com/hilmiikhsan/library-author-service/internal/dto"
	"github.com/hilmiikhsan/library-author-service/internal/interfaces"
	"github.com/hilmiikhsan/library-author-service/internal/validator"
)

type AuthorAPI struct {
	AuthorService interfaces.IAuthorService
	Validator     *validator.Validator
	author.UnimplementedAuthorServiceServer
}

func (api *AuthorAPI) GetDetailAuthor(ctx context.Context, req *author.AuthorRequest) (*author.AuthorResponse, error) {
	internalReq := dto.GetDetailAuthorRequest{
		ID: req.Id,
	}

	if err := api.Validator.Validate(internalReq); err != nil {
		helpers.Logger.Error("api::GetDetailAuthor - Failed to validate request : ", err)
		return &author.AuthorResponse{
			Message: "Failed to validate request",
			Data:    nil,
		}, nil
	}

	res, err := api.AuthorService.GetDetailAuthor(ctx, internalReq.ID)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrAuthorNotFound) {
			helpers.Logger.Error("api::GetDetailAuthor - Author not found")
			return &author.AuthorResponse{
				Message: constants.ErrAuthorNotFound,
				Data:    nil,
			}, nil
		}

		helpers.Logger.Error("api::GetDetailAuthor - Failed to get detail Author : ", err)
		return &author.AuthorResponse{
			Message: "Failed to get detail Author",
			Data:    nil,
		}, nil
	}

	return &author.AuthorResponse{
		Message: constants.SuccessMessage,
		Data: &author.AuthorData{
			Id:   res.ID,
			Name: res.Name,
		},
	}, nil
}
