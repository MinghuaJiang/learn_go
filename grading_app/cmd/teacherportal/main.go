package main

import (
	"context"
	"fmt"
	"grading_app/log"
	"grading_app/registry"
	"grading_app/service"
	"grading_app/teacherportal"
	stlog "log"
)

func main() {
	err := teacherportal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.TeacherPortal
	r.ServiceURL = serviceAddress

	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService}

	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		port,
		r,
		teacherportal.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at :%v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()

	fmt.Println("Shutting down teacher portal")
}
