package lexer

import (
	"lua-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	local five = 5
	local ten = 10
	
	local function add(x, y)
		return x + y
	end
	
	local result = add(five, ten)
	
	-- This is a comment
	
	if result == 15 then
		print("Success")
	else
		print("Failure")
	end
	
	while x < 10 do
		x = x + 1
	end
	
	for i = 1, 10 do
		print(i)
	end
	
	local str = "hello world"
	local concat = "hello" .. " " .. "world"
	local long = [[
		This is a
		long string
	]]
	
	local arr = {1, 2, 3}
	local len = #arr
	
	~= <= >= < > == and or not
	+ - * / % ^
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LOCAL, "local"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},

		{token.LOCAL, "local"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},

		{token.LOCAL, "local"},
		{token.FUNCTION, "function"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},

		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},

		{token.END, "end"},

		{token.LOCAL, "local"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},

		{token.IF, "if"},
		{token.IDENT, "result"},
		{token.EQ, "=="},
		{token.NUMBER, "15"},
		{token.THEN, "then"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.STRING, "Success"},
		{token.RPAREN, ")"},
		{token.ELSE, "else"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.STRING, "Failure"},
		{token.RPAREN, ")"},
		{token.END, "end"},

		{token.WHILE, "while"},
		{token.IDENT, "x"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.DO, "do"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.NUMBER, "1"},
		{token.END, "end"},

		{token.FOR, "for"},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.NUMBER, "1"},
		{token.COMMA, ","},
		{token.NUMBER, "10"},
		{token.DO, "do"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.RPAREN, ")"},
		{token.END, "end"},

		{token.LOCAL, "local"},
		{token.IDENT, "str"},
		{token.ASSIGN, "="},
		{token.STRING, "hello world"},

		{token.LOCAL, "local"},
		{token.IDENT, "concat"},
		{token.ASSIGN, "="},
		{token.STRING, "hello"},
		{token.CONCAT, ".."},
		{token.STRING, " "},
		{token.CONCAT, ".."},
		{token.STRING, "world"},

		{token.LOCAL, "local"},
		{token.IDENT, "long"},
		{token.ASSIGN, "="},
		{token.STRING, "\tThis is a\n\tlong string\n"},

		{token.LOCAL, "local"},
		{token.IDENT, "arr"},
		{token.ASSIGN, "="},
		{token.LBRACE, "{"},
		{token.NUMBER, "1"},
		{token.COMMA, ","},
		{token.NUMBER, "2"},
		{token.COMMA, ","},
		{token.NUMBER, "3"},
		{token.RBRACE, "}"},

		{token.LOCAL, "local"},
		{token.IDENT, "len"},
		{token.ASSIGN, "="},
		{token.HASH, "#"},
		{token.IDENT, "arr"},

		{token.NOT_EQ, "~="},
		{token.LTE, "<="},
		{token.GTE, ">="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EQ, "=="},
		{token.AND, "and"},
		{token.OR, "or"},
		{token.NOT, "not"},

		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.MODULO, "%"},
		{token.POWER, "^"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q (literal: %q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNumbers(t *testing.T) {
	input := `3 3.0 3.1416 314.16e-2 0.31416E1 0xff 0x0.1E`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NUMBER, "3"},
		{token.NUMBER, "3.0"},
		{token.NUMBER, "3.1416"},
		{token.NUMBER, "314.16e-2"},
		{token.NUMBER, "0.31416E1"},
		{token.NUMBER, "0xff"},
		{token.NUMBER, "0x0"},
		{token.DOT, "."},
		{token.NUMBER, "1E"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestComments(t *testing.T) {
	input := `
	local x = 5 -- single line comment
	local y = 10
	--[[
		This is a multi-line comment
		that spans multiple lines
	]]
	local z = x + y
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LOCAL, "local"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.LOCAL, "local"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.LOCAL, "local"},
		{token.IDENT, "z"},
		{token.ASSIGN, "="},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
	}
}