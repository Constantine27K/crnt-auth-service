package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	authService "github.com/Constantine27K/crnt-auth-service/internal/app/crnt-auth-service/auth"
	secretsGateway "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/gateway"
	secretsStorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/infrastructure/postgres"
	userService "github.com/Constantine27K/crnt-auth-service/internal/pkg/services/crnt-user-service"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/validation"
	"github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	"github.com/Constantine27K/crnt-sdk/pkg/authorization"
	"github.com/Constantine27K/crnt-sdk/pkg/token"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	loadEnv()
	setLogger()

	go createGrpcServer()
	createHttpServer()
}

func loadEnv() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Error("No .env file found",
			zap.Error(err),
		)
	}
}

func setLogger() {
	logLevel, _ := strconv.ParseInt(os.Getenv("LOG_LEVEL"), 10, 32)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(log.Level(logLevel))
}

func createGrpcServer() {
	address := os.Getenv("GRPC_ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Error(fmt.Sprintf("failed to listen:%v", address),
			zap.Error(err),
		)
	}

	grpcServer := grpc.NewServer()

	db, err := postgres.NewPostgres(postgres.Options{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
	})
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	secretGw := secretsGateway.NewSecretsGateway(db)
	secretStorage := secretsStorage.NewSecretStorage(secretGw)

	tokenMaker, err := token.NewMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	if err != nil {
		log.Fatalf("failed to create token maker: %v", err)
	}

	validator := validation.NewValidator()
	authorizer := authorization.NewAuthorizer(tokenMaker)

	connUserService, err := grpc.Dial(
		os.Getenv("USER_SERVICE_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cannot dial user-service: %v", err)
	}
	userServ := userService.NewService(connUserService)

	auth.RegisterAuthServer(grpcServer, authService.NewService(secretStorage, userServ, tokenMaker, authorizer, validator))
	log.Infof("grpc service started on %s", address)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("error during serving GRPC",
			zap.Error(err),
		)
	}
}

func createHttpServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcAddress := os.Getenv("GRPC_ADDRESS")
	// dial the gRPC server above to make a client connection
	conn, err := grpc.Dial(grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error(fmt.Sprintf("failed to dial:%s", grpcAddress),
			zap.Error(err),
		)
	}
	defer conn.Close()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	if err != nil {
		log.Error("failed to register user handler client",
			zap.Error(err),
		)
	}
	clientAuth := auth.NewAuthClient(conn)
	err = auth.RegisterAuthHandlerClient(ctx, rmux, clientAuth)
	if err != nil {
		log.Error("failed to register auth handler client",
			zap.Error(err),
		)
	}

	// creating swagger
	mux := http.NewServeMux()
	// mount the gRPC HTTP gateway to the root
	mux.Handle("/", rmux)
	fs := http.FileServer(http.Dir("./swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.Info("swagger created")

	httpAddress := os.Getenv("HTTP_ADDRESS")
	log.Infof("http service started on %s", httpAddress)

	// start a standard HTTP server with the router
	err = http.ListenAndServe(httpAddress, mux)
	if err != nil {
		log.Error("error during listening and serving HTTP",
			zap.Error(err),
		)
	}
}
