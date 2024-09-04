package main

import (
	"fmt"
	"unicode"
)

// TokenType represents the type of token.
type TokenType int

const (
	EOF TokenType = iota
	IDENTIFIER
	NUMBER
	PLUS
	MINUS
	ASTERISK
	SLASH
	LPAREN
	RPAREN
)

// Token represents a lexical token.
type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

// Lexer represents a lexical analyzer.
type Lexer struct {
	input                string
	start, current, line int
}

// NewLexer initializes a new lexer.
func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.isAtEnd() {
		return Token{Type: EOF, Lexeme: "", Line: l.line}
	}

	l.start = l.current

	char := l.advance()

	switch {
	case unicode.IsLetter(char):
		return l.identifier()
	case unicode.IsDigit(char):
		return l.number()
	case char == '+':
		return l.makeToken(PLUS)
	case char == '-':
		return l.makeToken(MINUS)
	case char == '*':
		return l.makeToken(ASTERISK)
	case char == '/':
		return l.makeToken(SLASH)
	case char == '(':
		return l.makeToken(LPAREN)
	case char == ')':
		return l.makeToken(RPAREN)
	}

	return Token{Type: EOF, Lexeme: "", Line: l.line}
}

// Helper methods
func (l *Lexer) advance() rune {
	l.current++
	return rune(l.input[l.current-1])
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.input)
}

func (l *Lexer) skipWhitespace() {
	for !l.isAtEnd() {
		char := l.peek()
		switch char {
		case ' ', '\r', '\t':
			l.advance()
		case '\n':
			l.line++
			l.advance()
		default:
			return
		}
	}
}

func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return 0
	}
	return rune(l.input[l.current])
}

func (l *Lexer) identifier() Token {
	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) {
		l.advance()
	}
	return Token{Type: IDENTIFIER, Lexeme: l.input[l.start:l.current], Line: l.line}
}

func (l *Lexer) number() Token {
	for unicode.IsDigit(l.peek()) {
		l.advance()
	}
	return Token{Type: NUMBER, Lexeme: l.input[l.start:l.current], Line: l.line}
}

func (l *Lexer) makeToken(tokenType TokenType) Token {
	return Token{Type: tokenType, Lexeme: l.input[l.start:l.current], Line: l.line}
}

func main() {
	input := "var x = 42 + 3 * (y - 5)"
	lexer := NewLexer(input)

	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		fmt.Printf("Token: %v\n", token)
	}
}
