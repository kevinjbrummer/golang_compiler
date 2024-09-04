package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers and Literals

	IDENT = "IDENT"
	INT = "INT"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
	STRING = "STRING"

	// Operators
	
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	SLASH = "/"
	ASTERISK = "*"
	BANG = "!"
	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="
	GT_EQ = ">="
	LT_EQ = "<="
	EXPONENT = "**"


	// Delimiters

	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"
	COLON = ":"


	// Keywords

	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

func LookUpIdent(ident string) TokenType {
	if tok,ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}