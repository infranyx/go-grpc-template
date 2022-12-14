package articleConfigurator

import (
	"context"

	articleV1 "github.com/infranyx/protobuf-template-go/golang_template/article/v1"

	sampleExtServiceUseCase "github.com/infranyx/go-microservice-template/external/sample_ext_service/usecase"
	articleGrpcController "github.com/infranyx/go-microservice-template/internal/article/delivery/grpc"
	articleHttpController "github.com/infranyx/go-microservice-template/internal/article/delivery/http"
	articleKafkaConsumer "github.com/infranyx/go-microservice-template/internal/article/delivery/kafka/consumer"
	articleKafkaProducer "github.com/infranyx/go-microservice-template/internal/article/delivery/kafka/producer"
	articleDomain "github.com/infranyx/go-microservice-template/internal/article/domain"
	articleJob "github.com/infranyx/go-microservice-template/internal/article/job"
	articleRepository "github.com/infranyx/go-microservice-template/internal/article/repository"
	articleUseCase "github.com/infranyx/go-microservice-template/internal/article/usecase"
	externalBridge "github.com/infranyx/go-microservice-template/pkg/external_bridge"
	infraContainer "github.com/infranyx/go-microservice-template/pkg/infra_container"
)

type configurator struct {
	ic        *infraContainer.IContainer
	extBridge *externalBridge.ExternalBridge
}

func NewConfigurator(ic *infraContainer.IContainer, extBridge *externalBridge.ExternalBridge) articleDomain.Configurator {
	return &configurator{ic: ic, extBridge: extBridge}
}

func (c *configurator) Configure(ctx context.Context) error {
	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(c.extBridge.SampleExtGrpcService)
	kafkaProducer := articleKafkaProducer.NewProducer(c.ic.KafkaWriter)
	repository := articleRepository.NewRepository(c.ic.Postgres)
	useCase := articleUseCase.NewUseCase(repository, seServiceUseCase, kafkaProducer)

	// grpc
	grpcController := articleGrpcController.NewController(useCase)
	articleV1.RegisterArticleServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	httpController := articleHttpController.NewController(useCase)
	articleHttpController.NewRouter(httpController).Register(httpRouterGp)

	// consumers
	articleKafkaConsumer.NewConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// jobs
	articleJob.NewJob(c.ic.Logger).StartJobs(ctx)

	return nil
}
