package service

import (
	matches "github.com/AAiTweb/Dating_Application/Matches"
)

type MatchServ struct {
	matchRepo matches.MatchRepository
}

func NewMatchService(_matchRepo matches.MatchRepository)matches.MatchService{
	return MatchServ{matchRepo:_matchRepo}
}

func (m MatchServ) DoMatch(id int) error {
	return m.matchRepo.DoMatch(id)
}


