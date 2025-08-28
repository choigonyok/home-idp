package observability

import (
	"net/http"

	k8sclient "github.com/choigonyok/idp/pkg/client/k8s"
	promclient "github.com/choigonyok/idp/pkg/client/prometheus"
	"github.com/choigonyok/idp/pkg/observability/log"
	"github.com/choigonyok/idp/pkg/observability/metric"
	"github.com/choigonyok/idp/pkg/observability/trace"
	tracemodel "github.com/choigonyok/idp/pkg/observability/trace/model"
	"github.com/choigonyok/idp/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/model"
	"github.com/spf13/viper"
)

type Observability struct {
	log    *log.Log
	metric *metric.Metric
	trace  *trace.Trace
	db     *store.SQLiteStore
}

func NewObservability(k8sClient *k8sclient.K8s, upgrader *websocket.Upgrader, db *store.SQLiteStore, promClient *promclient.Prometheus, resty *resty.Client, traceHost string) *Observability {
	return &Observability{
		log: log.NewLog(
			k8sClient,
			upgrader,
			db,
		),
		metric: metric.NewMetric(
			promClient,
		),
		trace: trace.NewTrace(
			resty,
			traceHost,
		),
		db: db,
	}
}

func (o *Observability) GetServiceLog(serviceID string, c *gin.Context) error {
	platform, err := o.db.GetServicePlatform(c.Param("serviceID"))
	if err != nil {
		return err
	}

	switch platform {
	case "k8s":
		o.log.GetK8sPodsLogFromServiceID(serviceID, c)
	}

	return nil
}

func (o *Observability) GetServiceMetrics(serviceID string, c *gin.Context) (model.Value, error) {
	metrics, err := o.db.GetMetrics(c.Param("serviceID"))
	if err != nil {
		return nil, err
	}

	for _, m := range *metrics {
		switch m.Source {
		case "prometheus":
			return o.metric.GetPrometheusMetric(&m)
		}
	}

	return nil, nil
}

func (o *Observability) GetServiceTraces(serviceID string, c *gin.Context) (*tracemodel.Traces, error) {
	v := viper.GetViper()
	if !v.GetBool("traces.enabled") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tracing is not enabled"})
		return nil, nil
	}

	switch v.GetString("traces.source") {
	case "tempo":
		return o.trace.GetTempoTraces(c.Param("serviceID"))
	default:
		return nil, nil
	}
}
