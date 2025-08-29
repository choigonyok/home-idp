package server

import (
	"fmt"
	"net/http"
	"strconv"

	k8sclient "github.com/choigonyok/idp/pkg/client/k8s"
	promclient "github.com/choigonyok/idp/pkg/client/prometheus"
	"github.com/choigonyok/idp/pkg/handlers"
	"github.com/choigonyok/idp/pkg/infra"
	"github.com/choigonyok/idp/pkg/observability"
	"github.com/choigonyok/idp/pkg/service"
	"github.com/choigonyok/idp/pkg/store"
	"github.com/choigonyok/idp/pkg/system"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

const (
	allowCORSCredentials = true
)

type Server struct {
	engine *gin.Engine
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewServer() *Server {
	v := viper.GetViper()
	env := v.GetString("environment")
	setGinMode(env)
	engine := gin.Default()
	cfg := getCORSConfig(env)
	engine.Use(cors.New(*cfg))

	db := store.NewSQLiteStore("idp.db")
	k8s := k8sclient.NewK8sClient(v, db)

	infra := infra.NewInfra(db, k8s)
	svc := service.NewService(db)
	sys := system.NewSystem(db)
	o11y := observability.NewObservability(
		k8s,
		&upgrader,
		db,
		promclient.NewPrometheusClient(v),
		resty.New(),
		v.GetString("traces.host"),
	)

	h := handlers.NewHandler(
		db,
		o11y,
		svc,
		infra,
		sys,
	)
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
		cfg.AllowMethods = []string{"OPTIONS", "POST", "GET", "PUT", "DELETE", "*"}
		cfg.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "*"}
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
