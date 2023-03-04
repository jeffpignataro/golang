package webserver

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"golang/pkg/helpers"
	"golang/pkg/webserver/endpoints"
	"golang/pkg/webserver/mux"
	"golang/pkg/webserver/routing"
)

func New() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				mux.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			routing.AsRoute(endpoints.NewHealthzHandler),
			routing.AsRoute(endpoints.NewQuitQuitQuitHandler),
			routing.AsRoute(endpoints.NewDebugzHandler),
			zap.NewExample,
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	port, _ := helpers.LookupEnv("PORT", "9001")
	srv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
