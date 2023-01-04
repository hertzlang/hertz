package lexer

type Token int

const (
	Iota = iota
	ILLEGAL
	IDENT
	INT
	SEMI
	ADD

	BUMP
	SUB

	MUL
	DIV
	ASSIGN
	OpenParent   
	CloseParent  
	OpenBracket  
	CloseBracket 
	OpenBrace    
	CloseBrace 
	Comma   
	Dot     
	DotDot  
	SemiCol

	RSHIFT 
	LSHIFT 
	OR     
	AND    
	XOR    
	NOT   
	OROR   
	ANDAND 
	GT     
	GTE    
	LT     
	Assign 
	LTE    
	EQ     
	NEQ    

	VARIABLE
	FUNCTION
	NEW 
	CLASS 
	RETURN 
	TRUE 
	FALSE
	IF 
	ELSE 
	ELIF
	FOR 
	IN
	INHERIT
)

var tokens = []string{
	Iota:     "IOTA",
	ILLEGAL:  "ILLEGAL",
	IDENT:    "IDENT",
	INT:      "INT",
	SEMI:     ";",
	// Infix ops
	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	BUMP: "!",	

	//brackets
	OpenParent   : "(",
	CloseParent  : ")",
	OpenBracket  : "[",
	CloseBracket : "]",
	OpenBrace    : "{",
	CloseBrace   : "}",

	Comma   : ",",
	Dot     : ".",
	DotDot  : "..",
	SemiCol : ";",


	//operators
	RSHIFT : ">>",
	LSHIFT : "<<",
	OR     : "|",
	AND    : "&",
	XOR    : "^",
	NOT    : "~",
	OROR   : "||",
	ANDAND : "&&",
	ASSIGN: "=",
	GT    : ">",
	GTE   : ">=",
	LT    : "<",
	Assign: "=",
	LTE   : "<=",
	EQ    : "==",
	NEQ   : "!=",

	//logic
	FUNCTION: "FUNCTION",
	VARIABLE: "VARIABLE",
	CLASS : "CLASS",
	RETURN : "RETURN",
    ELIF: "ELIF",
	INHERIT: "INHERITS",
	TRUE   : "TRUE",
	FALSE  : "FALSE",
	IF  : "IF",
	ELSE   : "ELSE",
	FOR    : "FOR",
	IN   : "IN",
	NEW: "NEW",
}

func (t Token) String() string {
	return tokens[t]
}