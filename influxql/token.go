package influxql

import (
	"strings"
)

// Token is a lexical token of the InfluxQL language.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	literal_beg
	// Literals
	IDENT     // main
	NUMBER    // 12345.67
	DURATION  // 13h
	STRING    // "abc"
	BADSTRING // "abc
	BADESCAPE // \q
	TRUE      // true
	FALSE     // false
	literal_end

	operator_beg
	// Operators
	ADD // +
	SUB // -
	MUL // *
	DIV // /

	AND // AND
	OR  // OR

	EQ  // ==
	NEQ // !=
	LT  // <
	LTE // <=
	GT  // >
	GTE // >=
	operator_end

	LPAREN // (
	RPAREN // )
	COMMA  // ,

	keyword_beg
	// Keywords
	AS
	ASC
	BY
	CONTINUOUS
	DELETE
	DESC
	DROP
	EXPLAIN
	FROM
	INNER
	JOIN
	LIMIT
	LIST
	MERGE
	ORDER
	QUERIES
	SELECT
	SERIES
	WHERE
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",

	IDENT:  "IDENT",
	NUMBER: "NUMBER",
	STRING: "STRING",
	TRUE:   "TRUE",
	FALSE:  "FALSE",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	AND: "AND",
	OR:  "OR",

	EQ:  "==",
	NEQ: "!=",
	LT:  "<",
	LTE: "<=",
	GT:  ">",
	GTE: ">=",

	LPAREN: "(",
	RPAREN: ")",
	COMMA:  ",",

	AS:         "AS",
	ASC:        "ASC",
	BY:         "BY",
	CONTINUOUS: "CONTINUOUS",
	DELETE:     "DELETE",
	DESC:       "DESC",
	DROP:       "DROP",
	EXPLAIN:    "EXPLAIN",
	FROM:       "FROM",
	INNER:      "INNER",
	JOIN:       "JOIN",
	LIMIT:      "LIMIT",
	LIST:       "LIST",
	MERGE:      "MERGE",
	ORDER:      "ORDER",
	QUERIES:    "QUERIES",
	SELECT:     "SELECT",
	SERIES:     "SERIES",
	WHERE:      "WHERE",
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for tok := keyword_beg + 1; tok < keyword_end; tok++ {
		keywords[strings.ToUpper(tokens[tok])] = tok
		keywords[strings.ToLower(tokens[tok])] = tok
	}
	for _, tok := range []Token{AND, OR} {
		keywords[strings.ToUpper(tokens[tok])] = tok
		keywords[strings.ToLower(tokens[tok])] = tok
	}
	keywords["true"] = TRUE
	keywords["false"] = FALSE
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return ""
}

// Precedence returns the operator precedence of the binary operator token.
func (tok Token) Precedence() int {
	switch tok {
	case OR:
		return 1
	case AND:
		return 2
	case EQ, NEQ, LT, LTE, GT, GTE:
		return 3
	case ADD, SUB:
		return 4
	case MUL, DIV:
		return 5
	}
	return 0
}

// IsLiteral returns true for literal tokens.
func (tok Token) IsLiteral() bool { return tok > literal_beg && tok < literal_end }

// IsOperator returns true for operator tokens.
func (tok Token) IsOperator() bool { return tok > operator_beg && tok < operator_end }

// IsKeyword returns true for keyword tokens.
func (tok Token) IsKeyword() bool { return tok > keyword_beg && tok < keyword_end }

// Lookup returns the token associated with a given string.
func Lookup(ident string) Token {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Pos specifies the line and character position of a token.
// The Char and Line are both zero-based indexes.
type Pos struct {
	Line int
	Char int
}
