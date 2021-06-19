package pool

import (
	"io"
	"log"
	"sync/atomic"
)

type DbConnection struct {
	ID int32
}

// Make DBConnection pointer implments IO.Closer
func (dbConn *DbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func CreateConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &DbConnection{id}, nil
}
