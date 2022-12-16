package config

import (
	"github.com/infranyx/go-grpc-template/pkg/constant"
	"github.com/infranyx/go-grpc-template/pkg/env"
)

type Config struct {
	App                  AppConfig
	Grpc                 GrpcConfig
	Http                 HttpConfig
	Postgres             PostgresConfig
	GoTemplateGrpcClient GrpcConfig
	Kafka                KafkaConfig
}

var Conf *Config

type AppConfig struct {
	AppEnv string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	DBName          string
	MaxConn         int
	MaxIdleConn     int
	MaxLifeTimeConn int
	SslMode         string
}
type GrpcConfig struct {
	Port int
	Host string
}

type HttpConfig struct {
	Port int
}

type KafkaConfig struct {
	Enabled       bool
	LogEvents     bool
	ClientId      string
	ClientGroupId string
	ClientBrokers []string
	Topic         string
}

func init() {
	Conf = newConfig()
}

func newConfig() *Config {
	return &Config{
		App: AppConfig{
			AppEnv: env.New("APP_ENV", constant.AppEnvDev).AsString(),
		},
		Grpc: GrpcConfig{
			Port: env.New("GRPC_PORT", constant.GrpcPort).AsInt(),
			Host: env.New("GRPC_HOST", constant.GrpcHost).AsString(),
		},
		Http: HttpConfig{
			Port: env.New("HTTP_PORT", constant.HttpPort).AsInt(),
		},
		Postgres: PostgresConfig{
			Host:            env.New("PG_HOST", nil).AsString(),
			Port:            env.New("PG_PORT", nil).AsString(),
			User:            env.New("PG_USER", nil).AsString(),
			Pass:            env.New("PG_PASS", nil).AsString(),
			DBName:          env.New("PG_DB", nil).AsString(),
			MaxConn:         env.New("PG_MAX_CONNECTIONS", constant.PgMaxConn).AsInt(),
			MaxIdleConn:     env.New("PG_MAX_IDLE_CONNECTIONS", constant.PgMaxIdleConn).AsInt(),
			MaxLifeTimeConn: env.New("PG_MAX_LIFETIME_CONNECTIONS", constant.PgMaxLifeTimeConn).AsInt(),
			SslMode:         env.New("PG_SSL_MODE", constant.PgSslMode).AsString(),
		},
		GoTemplateGrpcClient: GrpcConfig{
			Port: env.New("EXTERNAL_GO_TEMPLATE_GRPC_PORT", constant.GrpcPort).AsInt(),
			Host: env.New("EXTERNAL_GO_TEMPLATE_GRPC_HOST", constant.GrpcHost).AsString(),
		},
		Kafka: KafkaConfig{
			Enabled:       env.New("KAFKA_ENABLED", nil).AsBool(),
			LogEvents:     env.New("KAFKA_LOG_EVENTS", nil).AsBool(),
			ClientId:      env.New("KAFKA_CLIENT_ID", nil).AsString(),
			ClientGroupId: env.New("KAFKA_CLIENT_GROUP_ID", nil).AsString(),
			ClientBrokers: env.New("KAFKA_CLIENT_BROKERS", nil).AsStringSlice(","),
			Topic:         env.New("KAFKA_TOPIC", nil).AsString(),
		},
	}
}

func IsDevEnv() bool {
	return Conf.App.AppEnv == constant.AppEnvDev
}

func IsProdEnv() bool {
	return Conf.App.AppEnv == constant.AppEnvProd
}

func IsTestEnv() bool {
	return Conf.App.AppEnv == constant.AppEnvTest
}
