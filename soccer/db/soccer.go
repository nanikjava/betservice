package db

import (
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"time"

	"database/sql"
	"git.neds.sh/matty/entain/soccer/proto/soccer"
	"sync"
)

// SoccerRepo provides repository access to soccer matches.
type SoccerRepo interface {
	// Init will initialise our soccer repository.
	Init() error

	// List will return a list of matches.
	List(filter *soccer.ListMatchesRequest) ([]*soccer.Match, error)
}

type soccerRepo struct {
	db   *sql.DB
	init sync.Once
}

// List will return all soccer matches
func (r *soccerRepo) List(filter *soccer.ListMatchesRequest) ([]*soccer.Match, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSoccerQueries()[soccerList]

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanMatches(rows)
}

// NewSoccerRepo creates a new soccer repository.
func NewSoccerRepo(db *sql.DB) SoccerRepo {
	return &soccerRepo{db: db}
}

// scanMatches to scan the rows from database into soccer.Match
func (m *soccerRepo) scanMatches(
	rows *sql.Rows,
) ([]*soccer.Match, error) {
	var races []*soccer.Match

	for rows.Next() {
		var match soccer.Match
		var advertisedStart time.Time

		if err := rows.Scan(&match.Id, &match.League, &match.TeamHome, &match.TeamHomeManager, &match.TeamAway, &match.TeamAwayManager, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		match.AdvertisedStartTime = ts

		races = append(races, &match)
	}

	return races, nil
}

func (r *soccerRepo) Init() error {
	var err error

	r.init.Do(func() {
		err = r.seed()
	})

	return err
}
