package app

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"
	_ "topup_game/docs"
	"topup_game/internal/handler/api"
	response_api "topup_game/internal/mapper/response/api"
	middlewares "topup_game/internal/middleware"
	"topup_game/pkg/auth"
	"topup_game/pkg/dotenv"
	"topup_game/pkg/logger"
	"topup_game/pkg/upload_image"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

// @title Topup Game gRPC
// @version 1.0
// @description gRPC based Topup game service

// @host localhost:5000
// @BasePath /api/

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization
func RunClient() {
	flag.Parse()

	logger, err := logger.NewLogger()

	if err != nil {
		logger.Fatal("Failed to create logger", zap.Error(err))
	}

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.Fatal("Failed to connect to server", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1420"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowCredentials: true,
	}))

	middlewares.WebSecurityConfig(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	token, err := auth.NewManager(viper.GetString("SECRET_KEY"))

	if err != nil {
		logger.Fatal("Failed to create token manager", zap.Error(err))
	}

	mapping := response_api.NewResponseApiMapper()
	imageUpload := upload_image.NewImageUpload()

	depsHandler := api.Deps{
		Conn:        conn,
		Token:       token,
		E:           e,
		Logger:      logger,
		Mapping:     mapping,
		ImageUpload: imageUpload,
	}

	api.NewHandler(depsHandler)

	go func() {
		if err := e.Start(":5000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Server.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
