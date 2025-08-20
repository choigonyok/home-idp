package server

import (
	"fmt"
	"strconv"

	"github.com/choigonyok/idp/pkg/handlers"
	"github.com/choigonyok/idp/pkg/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	allowCORSCredentials = true
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	env := viper.GetString("environment")
	setGinMode(env)
	engine := gin.Default()
	cfg := getCORSConfig(env)
	engine.Use(cors.New(*cfg))

	db := store.NewSQLiteStore("idp.db")
	h := handlers.NewHandler(db)
	setupRouter(engine, h)

	return &Server{engine: engine}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", strconv.Itoa(viper.GetViper().GetInt("server.port")))
	return s.engine.Run(addr)
}

func getCORSConfig(env string) *cors.Config {
	cfg := cors.Config{AllowCredentials: allowCORSCredentials}

	switch env {
	case "prod":
		cfg.AllowOrigins = []string{""}
		cfg.AllowMethods = []string{""}
		cfg.AllowHeaders = []string{""}
		return &cfg
	case "dev":
		cfg.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8800", "http://localhost:4040", "http://localhost:8080", "http://localhost:8888"}
		cfg.AllowMethods = []string{"OPTIONS", "POST", "GET", "PUT", "DELETE"}
		cfg.AllowHeaders = []string{"application/json"}
		return &cfg
	case "test":
		cfg.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8800", "http://localhost:4040", "http://localhost:8080", "http://localhost:8888"}
		cfg.AllowMethods = []string{"OPTIONS", "POST", "GET", "PUT", "DELETE"}
		cfg.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
		return &cfg
	default:
		return nil
	}
}

func setGinMode(env string) {
	switch env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.TestMode)
	}
}
