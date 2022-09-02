package main

import (
	"database/sql"
	"flag"
	"git.neds.sh/matty/entain/soccer/db"
	"git.neds.sh/matty/entain/soccer/proto/soccer"
	"git.neds.sh/matty/entain/soccer/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	grpcEndpoint = flag.String("grpc-endpoint", "localhost:9001", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("failed running grpc server: %s", err)
	}
}

func run() error {
	conn, err := net.Listen("tcp", ":9001")
	if err != nil {
		return err
	}

	soccerDB, err := sql.Open("sqlite3", "./db/soccer.db")
	if err != nil {
		return err
	}

	soccerRepo := db.NewSoccerRepo(soccerDB)
	if err := soccerRepo.Init(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	soccer.RegisterSoccerServer(
		grpcServer,
		service.NewSoccerService(
			soccerRepo,
		),
	)

	log.Infof("gRPC server listening on: %s", *grpcEndpoint)

	if err := grpcServer.Serve(conn); err != nil {
		return err
	}

	return nil
}
