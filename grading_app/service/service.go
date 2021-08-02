package service

import (
	"context"
	"fmt"
	"grading_app/registry"
	"log"
	"net/http"
)

func Start(ctx context.Context, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg, port)

	err := registry.RegisterService(reg)

	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, reg registry.Registration, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", reg.ServiceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(reg.ServiceURL)
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
