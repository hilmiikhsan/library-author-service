package interfaces

import (
	"context"

	"github.com/hilmiikhsan/library-author-service/internal/models"
)

type IExternal interface {
	ValidateToken(ctx context.Context, token string) (models.TokenData, error)
}
