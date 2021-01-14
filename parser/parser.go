package parser

import (
	"github.com/ysk-tngc/monkey-go/ast"
	"github.com/ysk-tngc/monkey-go/lexer"
	"github.com/ysk-tngc/monkey-go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つトークンを読み込む。curToken と peekToken の両方がセットされる
	p.newToken()
	p.newToken()

	return p
}

func (p *Parser) newToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}
