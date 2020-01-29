package Matches

type MatchRepository interface {
	DoMatch(id int)error
}