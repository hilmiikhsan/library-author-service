package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-author-service/internal/dto"
	"github.com/hilmiikhsan/library-author-service/internal/models"
)

type IAuthorRepository interface {
	InsertNewAuthor(ctx context.Context, author *models.Author) error
	FindAuthorByID(ctx context.Context, id string) (*models.Author, error)
	FindAllAuthor(ctx context.Context, limit, offset int) ([]models.Author, error)
	UpdateNewAuthor(ctx context.Context, author *models.Author) error
	DeleteAuthorByID(ctx context.Context, id string) error
}

type IAuthorService interface {
	CreateAuthor(ctx context.Context, req *dto.CreateAuthorRequest) error
	GetDetailAuthor(ctx context.Context, id string) (*dto.GetDetailAuthorResponse, error)
	GetListAuthor(ctx context.Context, limit, offset int) (*dto.GetListAuthorResponse, error)
	UpdateAuthor(ctx context.Context, req *dto.UpdateAuthorRequest) error
	DeleteAuthor(ctx context.Context, id string) error
}

type IAuthorHandler interface {
	CreateAuthor(*gin.Context)
	GetDetailAuthor(*gin.Context)
	GetListAuthor(*gin.Context)
	UpdateAuthor(*gin.Context)
	DeleteAuthor(*gin.Context)
}
