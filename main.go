package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lightstep/otel-launcher-go/launcher"
	"github.com/robfig/cron/v3"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"log"
	"os"
)

func main() {
	otel := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("zoom2slack"),
		launcher.WithServiceVersion("0.1"),
	)
	defer otel.Shutdown()

	tracer := global.Tracer("ls")
	ctx := context.Background()
	c, span := tracer.Start(ctx, "start")
	defer span.End()

	_, ok := os.LookupEnv("LS_ACCESS_TOKEN")
	if !ok {
		_, sp := tracer.Start(c, "preflight")
		err := errors.New(fmt.Sprintf("Missing required env var: LS_ACCESS_TOKEN"))
		sp.SetAttribute("error", err)
		log.Fatal(err)
		sp.End()
	}
	sig := "reliability"
	crons(c, tracer, sig)
}

// extend to pull from a list of sigs and start crons for each
func crons(ctx context.Context, tracer trace.Tracer, sig string) {
	c := cron.New()
	counter := 0
	c.AddFunc("@every 30s", func() {
		// if parent is cancelled then cancel inflight crons
		con, cancel := context.WithCancel(ctx)
		defer cancel()
		_, sp := tracer.Start(con, "cron")
		sp.SetName("start cron")
		sp.SetAttribute("sig", sig)

		counter++
		sp.SetAttribute("invocations", counter)
		fmt.Println(counter)
		sp.End()
	})

	c.Run()
}