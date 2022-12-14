package articleDomain

import (
	"context"

	"github.com/google/uuid"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
	"github.com/labstack/echo/v4"
)

type Article struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"desc"`
}

type ArticleConfigurator interface {
	ConfigureArticle(ctx context.Context) error
}

type ArticleGrpcController interface {
	CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
	GetArticleById(ctx context.Context, req *articleV1.GetArticleByIdRequest) (*articleV1.GetArticleByIdResponse, error)
}

type ArticleHttpController interface {
	Create(c echo.Context) error
}

type ArticleUseCase interface {
	Create(ctx context.Context, article *articleDto.CreateArticle) (*Article, error)
}

type ArticleRepository interface {
	Create(ctx context.Context, article *articleDto.CreateArticle) (*Article, error)
}
