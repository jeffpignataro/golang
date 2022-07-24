package main

import (
	helloworldapi "github.com/jeffpignataro/golang/pkg/examples/hello-world-api"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(helloworldapi.Run),
	).Run()
}
