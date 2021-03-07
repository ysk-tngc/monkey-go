package ast

import (
	"testing"

	"github.com/ysk-tngc/monkey-go/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myValue"},
					Value: "myValue",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherValue"},
					Value: "anotherValue",
				},
			},
		},
	}

	if program.String() != "let myValue = anotherValue;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
