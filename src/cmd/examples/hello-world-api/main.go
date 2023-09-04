package main

import (
	"context"
	"net/http"

	helloworldapi "github.com/jeffpignataro/golang/pkg/examples/hello-world-api"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(http.NewServeMux),
		fx.Invoke(helloworldapi.New),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lifecycle fx.Lifecycle, mux *http.ServeMux) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go http.ListenAndServe(":9001", mux)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				srv := http.Server{Addr: "9001", Handler: mux}
				go srv.Shutdown(ctx)
				return nil
			},
		},
	)
}
