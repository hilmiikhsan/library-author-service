package cmd

import (
	"net"

	"github.com/hilmiikhsan/library-author-service/cmd/proto/author"
	"github.com/hilmiikhsan/library-author-service/helpers"
	api "github.com/hilmiikhsan/library-author-service/internal/grpc"
	authorRepository "github.com/hilmiikhsan/library-author-service/internal/repository/author"
	authorServices "github.com/hilmiikhsan/library-author-service/internal/services/author"
	"github.com/hilmiikhsan/library-author-service/internal/validator"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	dependencyGrpc := dependencyGrpcInject()

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "6002"))
	if err != nil {
		helpers.Logger.Fatal("failed to listen grpc port: ", err)
	}

	server := grpc.NewServer()
	author.RegisterAuthorServiceServer(server, dependencyGrpc.AuthorAPI)

	helpers.Logger.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "6002"))
	if err := server.Serve(lis); err != nil {
		helpers.Logger.Fatal("failed to serve grpc port: ", err)
	}
}

type DependencyGrpc struct {
	Logger    *logrus.Logger
	AuthorAPI *api.AuthorAPI
}

func dependencyGrpcInject() *DependencyGrpc {
	authorRepo := &authorRepository.AuthorRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	validator := validator.NewValidator()

	AuthorSvc := &authorServices.AuthorService{
		AuthorRepo: authorRepo,
		Logger:     helpers.Logger,
	}
	AuthorAPI := &api.AuthorAPI{
		AuthorService: AuthorSvc,
		Validator:     validator,
	}

	return &DependencyGrpc{
		Logger:    helpers.Logger,
		AuthorAPI: AuthorAPI,
	}
}
