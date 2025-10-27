package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	NUMBER = "NUMBER" // 123, 3.14, 0x1A, 1e10
	STRING = "STRING" // "hello", 'world', [[long string]]

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	MODULO   = "%"
	POWER    = "^"
	CONCAT   = ".."

	EQ     = "=="
	NOT_EQ = "~="
	LT     = "<"
	LTE    = "<="
	GT     = ">"
	GTE    = ">="

	// Logical operators
	AND = "and"
	OR  = "or"
	NOT = "not"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	HASH      = "#"

	// Keywords
	FUNCTION = "function"
	LOCAL    = "local"
	IF       = "if"
	THEN     = "then"
	ELSE     = "else"
	ELSEIF   = "elseif"
	END      = "end"
	WHILE    = "while"
	DO       = "do"
	FOR      = "for"
	IN       = "in"
	REPEAT   = "repeat"
	UNTIL    = "until"
	RETURN   = "return"
	BREAK    = "break"
	TRUE     = "true"
	FALSE    = "false"
	NIL      = "nil"
)

var keywords = map[string]TokenType{
	"and":      AND,
	"break":    BREAK,
	"do":       DO,
	"else":     ELSE,
	"elseif":   ELSEIF,
	"end":      END,
	"false":    FALSE,
	"for":      FOR,
	"function": FUNCTION,
	"if":       IF,
	"in":       IN,
	"local":    LOCAL,
	"nil":      NIL,
	"not":      NOT,
	"or":       OR,
	"repeat":   REPEAT,
	"return":   RETURN,
	"then":     THEN,
	"true":     TRUE,
	"until":    UNTIL,
	"while":    WHILE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
