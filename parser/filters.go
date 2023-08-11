package parser

import (
	"github.com/MarioJim/gonja/nodes"
	"github.com/MarioJim/gonja/tokens"
)

// Filter = IDENT | IDENT ":" FilterArg | IDENT "|" Filter
func (p *Parser) ParseFilter() (*nodes.FilterCall, error) {
	identToken := p.Match(tokens.Name)

	// Check filter ident
	if identToken == nil {
		return nil, p.Error("Filter name must be an identifier.", p.Current())
	}

	filter := &nodes.FilterCall{
		Token:  identToken,
		Name:   identToken.Val,
		Args:   []nodes.Expression{},
		Kwargs: map[string]nodes.Expression{},
	}

	// Check for filter-argument (2 tokens needed: ':' ARG)
	if p.Match(tokens.Lparen) != nil {
		if p.Current(tokens.VariableEnd) != nil {
			return nil, p.Error("Filter parameter required after '('.", nil)
		}

		for p.Match(tokens.Comma) != nil || p.Match(tokens.Rparen) == nil {
			// TODO: Handle multiple args and kwargs
			v, err := p.ParseExpression()
			if err != nil {
				return nil, err
			}

			if p.Match(tokens.Assign) != nil {
				key := v.Position().Val
				value, errValue := p.ParseExpression()
				if errValue != nil {
					return nil, errValue
				}
				filter.Kwargs[key] = value
			} else {
				filter.Args = append(filter.Args, v)
			}
		}
	}

	return filter, nil
}
