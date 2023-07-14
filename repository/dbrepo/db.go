package dbrepo

type Predicate string

var (
	EPd   = Predicate("=")
	NEPd  = Predicate("<>")
	GPd   = Predicate(">")
	GOEPd = Predicate(">=")
	SPd   = Predicate("<")
	SOEPd = Predicate("<=")
	LPd   = Predicate("LIKE")
)
