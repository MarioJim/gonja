package parser

import (
	"fmt"

	"github.com/MarioJim/gonja/nodes"
	"github.com/MarioJim/gonja/tokens"
	log "github.com/sirupsen/logrus"
)

func (p *Parser) ParseMath() (nodes.Expression, error) {
	log.WithFields(log.Fields{
		"current": p.Current(),
	}).Trace("ParseMath")

	expr, err := p.parseConcat()
	if err != nil {
		return nil, err
	}

	for p.Current(tokens.Add, tokens.Sub) != nil {
		op := BinOp(p.Pop())
		right, err := p.parseConcat()
		if err != nil {
			return nil, err
		}
		expr = &nodes.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: op,
		}
	}

	log.WithFields(log.Fields{
		"expr": expr,
	}).Trace("ParseMath return")
	return expr, nil
}

func (p *Parser) parseConcat() (nodes.Expression, error) {
	log.WithFields(log.Fields{
		"current": p.Current(),
	}).Trace("parseConcat")

	expr, err := p.ParseMathPrioritary()
	if err != nil {
		return nil, err
	}

	for p.Current(tokens.Tilde) != nil {
		op := BinOp(p.Pop())
		right, err := p.ParseMathPrioritary()
		if err != nil {
			return nil, err
		}
		expr = &nodes.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: op,
		}
	}

	log.WithFields(log.Fields{
		"expr": expr,
	}).Trace("parseConcat return")
	return expr, nil
}

func (p *Parser) ParseMathPrioritary() (nodes.Expression, error) {
	log.WithFields(log.Fields{
		"current": p.Current(),
	}).Trace("ParseMathPrioritary")

	expr, err := p.parseUnary()
	if err != nil {
		return nil, err
	}

	for p.Current(tokens.Mul, tokens.Div, tokens.Floordiv, tokens.Mod) != nil {
		op := BinOp(p.Pop())
		right, err := p.parseUnary()
		if err != nil {
			return nil, err
		}
		expr = &nodes.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: op,
		}
	}

	log.WithFields(log.Fields{
		"expr": expr,
	}).Trace("ParseMathPrioritary return")
	return expr, nil
}

func (p *Parser) parseUnary() (nodes.Expression, error) {
	log.WithFields(log.Fields{
		"current": p.Current(),
	}).Trace("parseUnary")

	sign := p.Match(tokens.Add, tokens.Sub)

	expr, err := p.ParsePower()
	if err != nil {
		return nil, err
	}

	if sign != nil {
		expr = &nodes.UnaryExpression{
			Operator: sign,
			Negative: sign.Val == "-",
			Term:     expr,
		}
	}

	log.WithFields(log.Fields{
		"expr": expr,
	}).Trace("parseUnary return")
	return expr, nil
}

func (p *Parser) ParsePower() (nodes.Expression, error) {
	log.WithFields(log.Fields{
		"current": p.Current(),
	}).Trace("ParsePower")

	if p.Current(tokens.In) != nil {
		return nil, nil
	}

	expr, err := p.ParseVariableOrLiteral()
	if err != nil {
		return nil, err
	}

	for p.Current(tokens.Pow) != nil {
		op := BinOp(p.Pop())
		right, err := p.ParseVariableOrLiteral()
		if err != nil {
			return nil, err
		}
		expr = &nodes.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: op,
		}
	}

	log.WithFields(log.Fields{
		"type": fmt.Sprintf("%T", expr),
		"expr": expr,
	}).Trace("ParsePower return")
	return expr, nil
}
