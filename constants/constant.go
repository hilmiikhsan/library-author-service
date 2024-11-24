package constants

const (
	SuccessMessage                = "success"
	ErrAuthorAlreadyExist         = "author already exist"
	ErrFailedBadRequest           = "failed to parse request"
	ErrAuthorizationIsEmpty       = "authorization is empty"
	ErrInvalidAuthorizationFormat = "invalid authorization format"
	ErrInvalidAuthorization       = "invalid authorization"
	ErrAuthorNotFound             = "author not found"
	ErrParamIdIsRequired          = "param id is required"
	ErrIdIsNotValidUUID           = "id is not valid uuid"
	ErrInvalidFormatDate          = "invalid format date"
	ErrAuthRolePermission         = "you do not have permission to access this endpoint"
)

const (
	HeaderAuthorization = "Authorization"
	TokenTypeAccess     = "token"
	DateTimeFormat      = "2006-01-02"
	AuthRoleUser        = "User"
	AuthRoleAdmin       = "Admin"
)
