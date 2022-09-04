package main

import (
	"context"
	"flag"
	"git.neds.sh/matty/entain/api/proto/soccer"
	"net/http"

	"git.neds.sh/matty/entain/api/proto/racing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	apiEndpoint    = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	grpcEndpoint   = flag.String("grpc-endpoint", "localhost:9000", "gRPC server endpoint")
	soccerEndpoint = flag.String("soccer-endpoint", "localhost:9001", "Soccer gRPC endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("failed running api server: %s", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*grpcEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	//register soccer sport handler endpoint to GRPC server
	if err := soccer.RegisterSoccerHandlerFromEndpoint(
		ctx,
		mux,
		*soccerEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	log.Infof("API server listening on: %s", *apiEndpoint)

	return http.ListenAndServe(*apiEndpoint, mux)
}
