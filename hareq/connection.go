package hareq

import (
	"fmt"

	"github.com/johnhaha/hakit/hadata"
)

type HaConnection struct {
	Host     string
	Port     int
	Protocol string
}

func (conn HaConnection) GetPort() string {
	if conn.Port == 0 {
		return ""
	}
	return hadata.NewStringBinder().BindString(":").BindInt(conn.Port).Value()
}

func (conn HaConnection) GetConnection() string {
	return fmt.Sprintf("%v://%v%v", conn.Protocol, conn.Host, conn.GetPort())
}
