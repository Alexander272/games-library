package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alexander272/games-library/internal/config"
	"github.com/Alexander272/games-library/internal/repository"
	"github.com/Alexander272/games-library/internal/server"
	"github.com/Alexander272/games-library/internal/service"
	"github.com/Alexander272/games-library/internal/transport"
	"github.com/Alexander272/games-library/pkg/auth"
	"github.com/Alexander272/games-library/pkg/database/mongo"
	"github.com/Alexander272/games-library/pkg/database/redis"
	"github.com/Alexander272/games-library/pkg/hasher"
	"github.com/Alexander272/games-library/pkg/logger"
	"github.com/Alexander272/games-library/pkg/storage"
	"github.com/joho/godotenv"
)

// @title Games Library
// @version 0.1
// @description API Server for Games Library App

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logger.Init(os.Stdout)
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	conf, err := config.Init("configs")
	if err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}

	// Dependencies
	mongoClient, err := mongo.NewClient(conf.Mongo.URI, conf.Mongo.User, conf.Mongo.Password)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	db := mongoClient.Database(conf.Mongo.Name)

	client, err := redis.NewRedisClient(redis.Config{
		Host:     conf.Redis.Host,
		Port:     conf.Redis.Port,
		DB:       conf.Redis.DB,
		Password: conf.Redis.Password,
	})
	if err != nil {
		logger.Fatalf("failed to initialize redis %s", err.Error())
	}

	hasher := hasher.NewBcryptHasher(conf.Auth.Bcrypt.MinCost, conf.Auth.Bcrypt.DefaultCost, conf.Auth.Bcrypt.MaxCost)
	tokenManager, err := auth.NewManager(conf.Auth.JWT.Key)
	if err != nil {
		logger.Fatalf("failed to initialize token manager: %s", err.Error())
	}

	storage, err := storage.NewFileStorage(conf.FileStorage.Bucket, conf.FileStorage.Endpoint)
	if err != nil {
		logger.Fatalf("failed to initialize file storage: %s", err.Error())
	}

	// Services, Repos & API Handlers
	repos := repository.NewRepo(db, client)
	services := service.NewServices(service.Deps{
		Repos:           repos,
		StorageProvider: storage,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		AccessTokenTTL:  conf.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: conf.Auth.JWT.RefreshTokenTTL,
		Domain:          conf.Http.Domain,
	})
	handlers := transport.NewHandler(services)

	// HTTP Server
	srv := server.NewServer(conf, handlers.Init(conf))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Infof("Application started on port: %s", conf.Http.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Errorf("error occured on db connection close: %s", err.Error())
	}
}
