package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	authService "github.com/Constantine27K/crnt-auth-service/internal/app/crnt-auth-service/auth"
	userService "github.com/Constantine27K/crnt-auth-service/internal/app/crnt-auth-service/user"
	"github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	"github.com/Constantine27K/crnt-auth-service/pkg/api/user"
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
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error(fmt.Sprintf("failed to listen localhost:%v", port),
			zap.Error(err),
		)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserRegistryServer(grpcServer, userService.NewService())
	auth.RegisterAuthServer(grpcServer, authService.NewService())
	log.Infof("grpc service started on port %s", port)

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

	grpcPort := os.Getenv("GRPC_PORT")
	// dial the gRPC server above to make a client connection
	conn, err := grpc.Dial(fmt.Sprintf(":%s", grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error(fmt.Sprintf("failed to dial localhost:%s", grpcPort),
			zap.Error(err),
		)
	}
	defer conn.Close()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	clientIssue := user.NewUserRegistryClient(conn)
	err = user.RegisterUserRegistryHandlerClient(ctx, rmux, clientIssue)
	if err != nil {
		log.Error("failed to register user handler client",
			zap.Error(err),
		)
	}
	clientEpic := auth.NewAuthClient(conn)
	err = auth.RegisterAuthHandlerClient(ctx, rmux, clientEpic)
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

	httpPort := os.Getenv("HTTP_PORT")
	log.Infof("http service started on port %s", httpPort)

	// start a standard HTTP server with the router
	err = http.ListenAndServe(fmt.Sprintf("localhost:%s", httpPort), mux)
	if err != nil {
		log.Error("error during listening and serving HTTP",
			zap.Error(err),
		)
	}
}