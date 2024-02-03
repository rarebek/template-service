package main

import (
	"net"
	"template-service3/config"
	pb "template-service3/genproto/user_service"
	"template-service3/pkg/db"
	"template-service3/pkg/logger"
	"template-service3/service"
	grpcClient "template-service3/service/grpc_client"

	"google.golang.org/grpc"
)

func main() {
    cfg := config.Load()

    log := logger.New(cfg.LogLevel, "template-service")
    defer logger.Cleanup(log)

    log.Info("main: sqlxConfig",
        logger.String("host", cfg.PostgresHost),
        logger.Int("port", cfg.PostgresPort),
        logger.String("database", cfg.PostgresDatabase))

    connDB, _, err := db.ConnectToDB(cfg)
    if err != nil {
        log.Fatal("sqlx connection to postgres error", logger.Error(err))
    }

    client, err := grpcClient.New(cfg)
    if err != nil {
        log.Fatal("error while creating new client", logger.Error(err))
    }

    userService := service.NewUserService(connDB, log, client)

    lis, err := net.Listen("tcp", cfg.RPCPort)
    if err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, userService)
    log.Info("main: server running",
        logger.String("port", cfg.RPCPort))

    if err := s.Serve(lis); err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }
}
