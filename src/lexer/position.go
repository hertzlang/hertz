package lexer

type position struct {
	Line   int
	Column int
}

func NewPos() position {
	return position{
		Line:   1,
		Column: 0,
	}
}