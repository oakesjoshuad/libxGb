// Package xgbgen intends to generate source code for xproto and extensions.
// Influenced by xcb
package xgbgen

import (
	"io"
	"text/scanner"
)

// TODO: xgbgen will lex and parse header files and create a go equivalent, siloed into respective modules at libxgb base

type tokenType int
type tokenLiteral string

type token struct {
	tokenType
	tokenLiteral
}

type stateFn func(*lexer) stateFn

// lexer ...
type lexer struct {
	scanner.Scanner
	syntax
	state  stateFn
	tokens chan token
}

// Syntax Rules for lexing/parsing
type syntax struct {
	description string
	definition  map[tokenType]interface{}
}

// lex returns a new lexer for the given input
func lex(input io.Reader, syn syntax) *lexer {
	l := &lexer{
		syntax: syn,
		state:  lexText,
		tokens: make(chan token),
	}
	// initialize the input
	l.Init(input)
	return l
}

func lexText(l *lexer) stateFn {
	return nil
}
