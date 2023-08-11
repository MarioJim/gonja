package parser

import (
	"fmt"

	"github.com/MarioJim/gonja/tokens"
	"github.com/pkg/errors"
)

// Error produces a nice error message and returns an error-object.
// The 'token'-argument is optional. If provided, it will take
// the token's position information. If not provided, it will
// automatically use the CURRENT token's position information.
func (p *Parser) Error(msg string, token *tokens.Token) error {
	if token == nil {
		return errors.New(msg)
	} else {
		return fmt.Errorf("%s, line: %d, col: %d, near: %q, token: %q", msg,
			token.Line, token.Col, token.Val, token)
	}
}
