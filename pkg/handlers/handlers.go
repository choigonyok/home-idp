package handlers

import (
	"net/http"

	"github.com/choigonyok/idp/pkg/auth"
	awscli "github.com/choigonyok/idp/pkg/client/aws"
	grafanacli "github.com/choigonyok/idp/pkg/client/grafana"
	jenkinscli "github.com/choigonyok/idp/pkg/client/jenkins"
	k8scli "github.com/choigonyok/idp/pkg/client/k8s"
	"github.com/choigonyok/idp/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Handler struct {
	store   *store.SQLiteStore
	aws     *awscli.AWS
	k8s     *k8scli.K8s
	jenkins *jenkinscli.Jenkins
	grafana *grafanacli.Grafana
	Logger  *zap.SugaredLogger
}

func NewHandler(s *store.SQLiteStore) *Handler {
	logger := &zap.SugaredLogger{}
	v := viper.GetViper()
	env := v.GetString("environment")
	switch env {
	case "test":
		l, _ := zap.NewDevelopment()
		logger = l.Sugar()
	case "dev":
		l, _ := zap.NewDevelopment()
		logger = l.Sugar()
	case "prod":
		l, _ := zap.NewProduction()
		logger = l.Sugar()
	default:
		logger.Fatalf("%s: %s", "Invalid environment configuration:", env)
	}

	return &Handler{
		store:   s,
		aws:     getAWSClient(v),
		k8s:     getK8sClient(v),
		jenkins: getJenkinsClient(v),
		grafana: getGrafanaClient(v),
		Logger:  logger,
	}
}

func getJenkinsClient(v *viper.Viper) *jenkinscli.Jenkins {
	jenkinsHost := v.GetString("jenkins.host")
	auth := v.GetString("jenkins.auth")
	switch auth {
	case "apiToken":
		username := v.GetString("jenkins.credentials.username")
		apiToken := v.GetString("jenkins.credentials.apiToken")
		return jenkinscli.NewJenkinsFromAPIToken(jenkinsHost, username, apiToken)
	default:
		zap.S().Fatalf("%s: %s\n", "Invalid jenkins auth method", auth)
		return nil
	}
}

func getK8sClient(v *viper.Viper) *k8scli.K8s {
	auth := v.GetString("k8s.auth")
	switch auth {
	case "kubeconfig":
		return k8scli.NewFromKubeconfig()
	default:
		zap.S().Fatalf("%s: %s\n", "Invalid k8s auth method", auth)
		return nil
	}
}

func getAWSClient(v *viper.Viper) *awscli.AWS {
	auth := v.GetString("aws.auth")
	switch auth {
	case "env":
		return awscli.NewFromEnv()
	default:
		zap.S().Fatalf("%s: %s\n", "Invalid aws auth method", auth)
		return nil
	}
}

func getGrafanaDataSources(v *viper.Viper) *map[string]string {
	m := make(map[string]string)

	if v.GetBool("grafana.dataSources.prometheus.enabled") {
		m["prometheus"] = v.GetString("grafana.dataSources.prometheus.uid")
	}
	if v.GetBool("grafana.dataSources.tempo.enabled") {
		m["tempo"] = v.GetString("grafana.dataSources.tempo.uid")
	}
	if v.GetBool("grafana.dataSources.pyroscope.enabled") {
		m["pyroscope"] = v.GetString("grafana.dataSources.pyroscope.uid")
	}
	return &m
}

func getGrafanaClient(v *viper.Viper) *grafanacli.Grafana {
	host := v.GetString("grafana.host")
	auth := v.GetString("grafana.auth")
	m := getGrafanaDataSources(v)

	switch auth {
	case "apiToken":
		return grafanacli.NewGrafanaFromAPIToken(host, v.GetString("grafana.credentials.apiToken"), m)
	default:
		zap.S().Fatalf("%s: %s\n", "Invalid grafana auth method", auth)
		return nil
	}
}

// JWT middleware wrapper
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("idp_token")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		userID, err := auth.ValidateJWT(token)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
