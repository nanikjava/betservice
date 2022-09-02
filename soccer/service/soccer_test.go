package service

import (
	"context"
	"database/sql"
	"git.neds.sh/matty/entain/soccer/db"
	pb "git.neds.sh/matty/entain/soccer/proto/soccer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"path/filepath"
	"runtime"
	"testing"
)

// setup for setting up GRPC
func setup() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()
	//we want to use soccer.db in the local working directory to isolate from the
	//other database that is available in the soccer directory
	soccerDB, _ := sql.Open("sqlite3", getCurrentDirectory()+"/soccer.db")

	soccerRepo := db.NewSoccerRepo(soccerDB)
	if err := soccerRepo.Init(); err != nil {
		log.Fatal(err)
	}

	pb.RegisterSoccerServer(server, NewSoccerService(soccerRepo))

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestSoccerServer_ListMatches(t *testing.T) {
	tests := []struct {
		testname string
		norecs   int
	}{
		{
			"Total number of records",
			100,
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(setup()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSoccerClient(conn)

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			request := &pb.ListMatchesRequest{}
			response, err := client.ListMatches(ctx, request)

			if err != nil {
				log.Fatal(err)
			}

			if response != nil {
				if len(response.Matches) != tt.norecs {
					t.Error("response: expected", tt.norecs, "received", len(response.Matches))
				}
			}

		})
	}
}

// getCurrentDirectory use this to get the current working directory
func getCurrentDirectory() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Dir(f)
}
