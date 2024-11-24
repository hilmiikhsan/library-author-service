package author

import (
	"context"
	"errors"
	"time"

	"github.com/hilmiikhsan/library-author-service/constants"
	"github.com/hilmiikhsan/library-author-service/helpers"
	"github.com/hilmiikhsan/library-author-service/internal/dto"
	"github.com/hilmiikhsan/library-author-service/internal/interfaces"
	"github.com/hilmiikhsan/library-author-service/internal/models"
	"github.com/sirupsen/logrus"
)

type AuthorService struct {
	AuthorRepo interfaces.IAuthorRepository
	Logger     *logrus.Logger
}

func (s *AuthorService) CreateAuthor(ctx context.Context, req *dto.CreateAuthorRequest) error {
	birthDate, err := helpers.ParseDate(req.BirthDate, constants.DateTimeFormat)
	if err != nil {
		s.Logger.Error("author::CreateAuthor - failed to parse birth date: ", err)
		return errors.New(constants.ErrInvalidFormatDate)
	}

	err = s.AuthorRepo.InsertNewAuthor(ctx, &models.Author{
		Name:      req.Name,
		Bio:       req.Bio,
		BirthDate: birthDate,
	})
	if err != nil {
		s.Logger.Error("author::CreateAuthor - failed to insert new author: ", err)
		return err
	}

	return nil
}

func (s *AuthorService) GetDetailAuthor(ctx context.Context, id string) (*dto.GetDetailAuthorResponse, error) {
	authorData, err := s.AuthorRepo.FindAuthorByID(ctx, id)
	if err != nil {
		s.Logger.Error("author::GetAuthorDetail - failed to find Author by id: ", err)
		return nil, err
	}

	return &dto.GetDetailAuthorResponse{
		ID:        authorData.ID.String(),
		Name:      authorData.Name,
		Bio:       authorData.Bio,
		BirthDate: authorData.BirthDate.Format(constants.DateTimeFormat),
		DeathDate: helpers.FormatNullableDate(authorData.DeathDate, constants.DateTimeFormat),
	}, nil
}

func (s *AuthorService) GetListAuthor(ctx context.Context, limit, offset int) (*dto.GetListAuthorResponse, error) {
	pageSize := limit
	pageIndex := (offset - 1) * limit

	authorData, err := s.AuthorRepo.FindAllAuthor(ctx, pageSize, pageIndex)
	if err != nil {
		s.Logger.Error("author::GetListAuthor - failed to find all Author: ", err)
		return nil, err
	}

	categories := make([]dto.Author, 0)
	for _, author := range authorData {
		categories = append(categories, dto.Author{
			ID:        author.ID.String(),
			Name:      author.Name,
			Bio:       author.Bio,
			BirthDate: author.BirthDate.Format(constants.DateTimeFormat),
			DeathDate: helpers.FormatNullableDate(author.DeathDate, constants.DateTimeFormat),
		})
	}

	pagination := dto.Pagination{
		Page:  offset,
		Limit: limit,
	}

	response := &dto.GetListAuthorResponse{
		AuthorList: categories,
		Pagination: pagination,
	}

	return response, nil
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, req *dto.UpdateAuthorRequest) error {
	authorData, err := s.AuthorRepo.FindAuthorByID(ctx, req.ID)
	if err != nil {
		s.Logger.Error("author::UpdateAuthor - failed to find Author by id: ", err)
		return err
	}

	if len(authorData.Name) == 0 {
		s.Logger.Error("author::UpdateAuthor - Author not found")
		return errors.New(constants.ErrAuthorNotFound)
	}

	birthDate, err := helpers.ParseDate(req.BirthDate, constants.DateTimeFormat)
	if err != nil {
		s.Logger.Error("author::UpdateAuthor - failed to parse birth date: ", err)
		return errors.New(constants.ErrInvalidFormatDate)
	}

	var deathDate time.Time
	if req.DeathDate != "" {
		deathDate, err = helpers.ParseDate(req.DeathDate, constants.DateTimeFormat)
		if err != nil {
			s.Logger.Error("author::UpdateAuthor - failed to parse birth date: ", err)
			return errors.New(constants.ErrInvalidFormatDate)
		}
	}

	err = s.AuthorRepo.UpdateNewAuthor(ctx, &models.Author{
		ID:        authorData.ID,
		Name:      req.Name,
		Bio:       req.Bio,
		BirthDate: birthDate,
		DeathDate: helpers.NullTimeScan(deathDate),
	})
	if err != nil {
		s.Logger.Error("author::UpdateAuthor - failed to update Author: ", err)
		return err
	}

	return nil
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, id string) error {
	authorData, err := s.AuthorRepo.FindAuthorByID(ctx, id)
	if err != nil {
		s.Logger.Error("author::DeleteAuthor - failed to find Author by id: ", err)
		return err
	}

	if len(authorData.Name) == 0 {
		s.Logger.Error("author::DeleteAuthor - Author not found")
		return errors.New(constants.ErrAuthorNotFound)
	}

	err = s.AuthorRepo.DeleteAuthorByID(ctx, id)
	if err != nil {
		s.Logger.Error("author::DeleteAuthor - failed to delete Author: ", err)
		return err
	}

	return nil
}
