package dto

type CreateAuthorRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=100"`
	Bio       string `json:"bio" validate:"required,min=2,max=100"`
	BirthDate string `json:"birth_date" validate:"required"`
}

type UpdateAuthorRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required,min=2,max=100"`
	Bio       string `json:"bio" validate:"required,min=2,max=100"`
	BirthDate string `json:"birth_date" validate:"required"`
	DeathDate string `json:"death_date"`
}

type GetDetailAuthorRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetDetailAuthorResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	BirthDate string `json:"birth_date"`
	DeathDate string `json:"death_date"`
}

type GetListAuthorResponse struct {
	AuthorList []Author   `json:"author_list"`
	Pagination Pagination `json:"pagination"`
}

type Author struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	BirthDate string `json:"birth_date"`
	DeathDate string `json:"death_date"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
