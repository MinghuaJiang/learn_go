package work

import (
	"log"
	"time"
)

type NamePrinter struct {
	Name string
}

func (m *NamePrinter) Task() {
	log.Println(m.Name)
	time.Sleep(time.Second)
}
