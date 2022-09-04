package service

import (
	"context"
	db "git.neds.sh/matty/entain/soccer/db"
	"git.neds.sh/matty/entain/soccer/proto/soccer"
)

type Soccer interface {
	// ListMatches returns a list of all match.
	ListMatches(ctx context.Context, in *soccer.ListMatchesRequest) (*soccer.ListMatchesResponse, error)
}

// soccerService implements the Soccer interface.
type soccerService struct {
	soccerRepo db.SoccerRepo
}

// NewSoccerService instantiates and returns a new soccerService.
func NewSoccerService(soccerRepo db.SoccerRepo) Soccer {
	return &soccerService{soccerRepo}
}

// ListMatches returns all soccer matches
func (s *soccerService) ListMatches(ctx context.Context, in *soccer.ListMatchesRequest) (*soccer.ListMatchesResponse, error) {
	matches, err := s.soccerRepo.List(in)
	if err != nil {
		return nil, err
	}

	return &soccer.ListMatchesResponse{Matches: matches}, nil
}
