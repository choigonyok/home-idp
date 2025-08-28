package log

import (
	client "github.com/choigonyok/idp/pkg/client/k8s"
	"github.com/choigonyok/idp/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Log struct {
	k8s      *client.K8s
	upgrader *websocket.Upgrader
	db       *store.SQLiteStore
}

func NewLog(k8s *client.K8s, upgrader *websocket.Upgrader, db *store.SQLiteStore) *Log {
	return &Log{
		k8s:      k8s,
		upgrader: upgrader,
		db:       db,
	}
}

func (l *Log) GetK8sPodsLogFromServiceID(serviceID string, c *gin.Context) error {
	conn, err := l.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	ns, err := l.db.GetServiceNamespace(serviceID)

	_, err = l.k8s.LogService(conn, serviceID, ns)
	if err != nil {
		return err
	}

	return nil
}
