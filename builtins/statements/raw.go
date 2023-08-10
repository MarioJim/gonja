package statements

import (
	"fmt"

	"github.com/MarioJim/gonja/exec"
	"github.com/MarioJim/gonja/nodes"
	"github.com/MarioJim/gonja/parser"
	"github.com/MarioJim/gonja/tokens"
)

type RawStmt struct {
	Data *nodes.Data
	// Content string
}

func (stmt *RawStmt) Position() *tokens.Token { return stmt.Data.Position() }
func (stmt *RawStmt) String() string {
	t := stmt.Position()
	return fmt.Sprintf("RawStmt(Line=%d Col=%d)", t.Line, t.Col)
}

func (stmt *RawStmt) Execute(r *exec.Renderer, tag *nodes.StatementBlock) error {
	_, err := r.Out.WriteString(stmt.Data.Data.Val)
	return err
}

func rawParser(p *parser.Parser, args *parser.Parser) (nodes.Statement, error) {
	stmt := &RawStmt{}

	wrapper, _, err := p.WrapUntil("endraw")
	if err != nil {
		return nil, err
	}
	node := wrapper.Nodes[0]
	data, ok := node.(*nodes.Data)
	if ok {
		stmt.Data = data
	} else {
		return nil, p.Error("raw statement can only contains a single data node", node.Position())
	}

	if !args.End() {
		return nil, args.Error("raw statement doesn't accept parameters.", args.Current())
	}

	return stmt, nil
}

func init() {
	All.Register("raw", rawParser)
}
