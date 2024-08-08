package lexer

import "go-interpreter/token"

type Lexer struct {
	input string
	position int //pointer to curr position of ch
	readPosition int //poiner to next position
	ch byte
}
//*Lexer is a pointer of Lexer type

func New(input string) *Lexer {
	l := &Lexer{input: input} // new Lexer object, initialises input as input
	l.readChar()
	return l
}

func(l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
	tok = newToken(token.ASSIGN, l.ch)
	case ';':
	tok = newToken(token.SEMICOLON, l.ch)
	case '(':
	tok = newToken(token.LPAREN, l.ch)
	case ')':
	tok = newToken(token.RPAREN, l.ch)
	case ',':
	tok = newToken(token.COMMA, l.ch)
	case '+':
	tok = newToken(token.PLUS, l.ch)
	case '{':
	tok = newToken(token.LBRACE, l.ch)
	case '}':
	tok = newToken(token.RBRACE, l.ch)
	case 0:
	tok.Literal = ""
	tok.Type = token.EOF
	}
	l.readChar()
	return tok
	}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
	}