package articleGrpcController

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	articleV1 "github.com/infranyx/protobuf-template-go/golang_template/article/v1"

	articleDomain "github.com/infranyx/go-microservice-template/internal/article/domain"
	articleDto "github.com/infranyx/go-microservice-template/internal/article/dto"
	articleException "github.com/infranyx/go-microservice-template/internal/article/exception"
)

type controller struct {
	useCase articleDomain.UseCase
}

func NewController(uc articleDomain.UseCase) articleDomain.GrpcController {
	return &controller{
		useCase: uc,
	}
}

func (c *controller) CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error) {
	aDto := &articleDto.CreateArticleRequestDto{
		Name:        req.Name,
		Description: req.Desc,
	}
	err := aDto.ValidateCreateArticleDto()
	if err != nil {
		return nil, articleException.CreateArticleValidationExc(err)
	}

	article, err := c.useCase.CreateArticle(ctx, aDto)
	if err != nil {
		return nil, err
	}

	return &articleV1.CreateArticleResponse{
		Id:   article.ID.String(),
		Name: article.Name,
		Desc: article.Description,
	}, nil
}

func (c *controller) GetArticleById(ctx context.Context, req *articleV1.GetArticleByIdRequest) (*articleV1.GetArticleByIdResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
