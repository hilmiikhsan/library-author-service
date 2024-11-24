package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-author-service/external"
	"github.com/hilmiikhsan/library-author-service/helpers"
	authorAPI "github.com/hilmiikhsan/library-author-service/internal/api/author"
	healthCheckAPI "github.com/hilmiikhsan/library-author-service/internal/api/health_check"
	"github.com/hilmiikhsan/library-author-service/internal/interfaces"
	authorRepository "github.com/hilmiikhsan/library-author-service/internal/repository/author"
	authorServices "github.com/hilmiikhsan/library-author-service/internal/services/author"
	healthCheckServices "github.com/hilmiikhsan/library-author-service/internal/services/health_check"
	"github.com/hilmiikhsan/library-author-service/internal/validator"
	"github.com/sirupsen/logrus"
)

func ServeHTTP() {
	dependency := dependencyInject()

	router := gin.Default()

	router.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	authorV1 := router.Group("/author/v1")
	authorV1.POST("/create", dependency.MiddlewareValidateToken, dependency.AuthorAPI.CreateAuthor)
	authorV1.GET("/:id", dependency.MiddlewareValidateToken, dependency.AuthorAPI.GetDetailAuthor)
	authorV1.GET("/", dependency.MiddlewareValidateToken, dependency.AuthorAPI.GetListAuthor)
	authorV1.PUT("/update", dependency.MiddlewareValidateToken, dependency.AuthorAPI.UpdateAuthor)
	authorV1.DELETE("/:id", dependency.MiddlewareValidateToken, dependency.AuthorAPI.DeleteAuthor)

	err := router.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		helpers.Logger.Fatal("failed to run http server: ", err)
	}
}

type Dependency struct {
	Logger           *logrus.Logger
	AuthorRepository interfaces.IAuthorRepository

	HealthcheckAPI interfaces.IHealthcheckHandler
	AuthorAPI      interfaces.IAuthorHandler
	External       interfaces.IExternal
}

func dependencyInject() Dependency {
	helpers.SetupLogger()

	healthcheckSvc := &healthCheckServices.Healthcheck{}
	healthcheckAPI := &healthCheckAPI.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	authorRepo := &authorRepository.AuthorRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	validator := validator.NewValidator()

	authorSvc := &authorServices.AuthorService{
		AuthorRepo: authorRepo,
		Logger:     helpers.Logger,
	}
	authorAPI := &authorAPI.AuthorHandler{
		AuthorService: authorSvc,
		Validator:     validator,
	}

	external := &external.External{
		Logger: helpers.Logger,
	}

	return Dependency{
		Logger:           helpers.Logger,
		AuthorRepository: authorRepo,
		HealthcheckAPI:   healthcheckAPI,
		AuthorAPI:        authorAPI,
		External:         external,
	}
}
