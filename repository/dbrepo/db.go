package dbrepo

type Predicate string

var (
	EqualPd          = Predicate("=")
	NotEqualPd       = Predicate("<>")
	GreaterPd        = Predicate(">")
	GreaterOrEqualPd = Predicate(">=")
	SmallerPd        = Predicate("<")
	SmallerOrEqualPd = Predicate("<=")
	LikePd           = Predicate("LIKE")
)
