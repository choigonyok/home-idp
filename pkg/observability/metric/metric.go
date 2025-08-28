package metric

import (
	"context"
	"fmt"
	"time"

	client "github.com/choigonyok/idp/pkg/client/prometheus"
	idpmodel "github.com/choigonyok/idp/pkg/model"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type Metric struct {
	prom *client.Prometheus
}

func NewMetric(prom *client.Prometheus) *Metric {
	return &Metric{
		prom: prom,
	}
}

func (m *Metric) GetPrometheusMetric(metric *idpmodel.Metric) (model.Value, error) {
	v1api := v1.NewAPI(*m.prom.Client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := metric.Query
	result, warnings, err := v1api.Query(ctx, query, time.Now())
	if err != nil {
		return nil, err
	}
	if len(warnings) > 0 {
		fmt.Println("Warnings:", warnings)
	}

	return result, nil

	// for _, sample := range result.(model.Vector) {
	// 	fmt.Printf("Metric: %v Value: %v\n", sample.Metric, sample.Value)
	// }

}
