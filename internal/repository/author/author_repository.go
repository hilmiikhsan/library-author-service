package author

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hilmiikhsan/library-author-service/constants"
	"github.com/hilmiikhsan/library-author-service/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthorRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
	Redis  *redis.Client
}

func (r *AuthorRepository) InsertNewAuthor(ctx context.Context, author *models.Author) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryInsertNewAuthor),
		author.Name,
		author.Bio,
		author.BirthDate,
	)
	if err != nil {
		r.Logger.Error("author::InsertNewAuthor - failed to insert new Author: ", err)
		return err
	}

	return nil
}

func (r *AuthorRepository) FindAuthorByID(ctx context.Context, id string) (*models.Author, error) {
	var (
		res = new(models.Author)
	)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindAuthorByID), id)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("author::FindAuthorByID - author doesnt exist")
			return res, errors.New(constants.ErrAuthorNotFound)
		}

		r.Logger.Error("author::FindAuthorByID - failed to find author by id: ", err)
		return nil, err
	}

	return res, nil
}

func (r *AuthorRepository) FindAllAuthor(ctx context.Context, limit, offset int) ([]models.Author, error) {
	var (
		res      = make([]models.Author, 0)
		cacheKey = fmt.Sprintf("authors:limit:%d:offset:%d", limit, offset)
	)

	cachedData, err := r.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &res)
		if err == nil {
			r.Logger.Info("category::FindAllAuthor - Data retrieved from cache")
			return res, nil
		}
		r.Logger.Warn("category::FindAllAuthor - Failed to unmarshal cache data: ", err)
	}

	err = r.DB.SelectContext(ctx, &res, r.DB.Rebind(queryFindAllAuthor), limit, offset)
	if err != nil {
		r.Logger.Error("author::FindAllAuthor - failed to find all author: ", err)
		return nil, err
	}

	dataToCache, err := json.Marshal(res)
	if err != nil {
		r.Logger.Warn("category::FindAllAuthor - Failed to marshal data for caching: ", err)
	} else {
		err = r.Redis.Set(ctx, cacheKey, dataToCache, 5*time.Minute).Err()
		if err != nil {
			r.Logger.Warn("category::FindAllAuthor - Failed to cache data: ", err)
		}
	}

	return res, nil
}

func (r *AuthorRepository) UpdateNewAuthor(ctx context.Context, author *models.Author) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryUpdateNewAuthor),
		author.Name,
		author.Bio,
		author.BirthDate,
		author.DeathDate,
		author.ID,
	)
	if err != nil {
		r.Logger.Error("author::UpdateNewAuthor - failed to update new author: ", err)
		return err
	}

	return nil
}

func (r *AuthorRepository) DeleteAuthorByID(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryDeleteAuthorByID), id)
	if err != nil {
		r.Logger.Error("author::DeleteAuthorByID - failed to delete author by id: ", err)
		return err
	}

	return nil
}
