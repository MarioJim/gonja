package exec

import (
	"bytes"
	"io"
	"strings"

	"github.com/pkg/errors"

	"github.com/MarioJim/gonja/nodes"
	"github.com/MarioJim/gonja/parser"
	"github.com/MarioJim/gonja/tokens"
)

type TemplateLoader interface {
	GetTemplate(string) (*Template, error)
	Path(string) (string, error)
}

type Template struct {
	Name   string
	Reader io.Reader
	Source string

	Env    *EvalConfig
	Loader TemplateLoader

	Tokens *tokens.Stream
	Parser *parser.Parser

	Root   *nodes.Template
	Macros MacroSet
}

func NewTemplate(name string, source string, cfg *EvalConfig) (*Template, error) {
	// Create the template
	t := &Template{
		Env:    cfg,
		Name:   name,
		Source: source,
		Tokens: tokens.Lex(source),
	}

	// Parse it
	t.Parser = parser.NewParser(name, cfg.Config, t.Tokens)
	t.Parser.Statements = *t.Env.Statements
	t.Parser.TemplateParser = t.Env.GetTemplate
	root, err := t.Parser.Parse()
	if err != nil {
		return nil, err
	}
	t.Root = root

	return t, nil
}

func (tpl *Template) execute(ctx map[string]any, out io.StringWriter) error {
	exCtx := tpl.Env.Globals.Inherit()
	exCtx.Update(ctx)

	var builder strings.Builder
	renderer := NewRenderer(exCtx, &builder, tpl.Env, tpl)

	err := renderer.Execute()
	if err != nil {
		return errors.Wrap(err, `Unable to execute template`)
	}
	if _, err = out.WriteString(renderer.String()); err != nil {
		return errors.Wrap(err, `Unable to execute template`)
	}

	return nil
}

func (tpl *Template) newBufferAndExecute(ctx map[string]any) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	if err := tpl.execute(ctx, &buffer); err != nil {
		return nil, err
	}
	return &buffer, nil
}

// Executes the template and returns the rendered template as a []byte
func (tpl *Template) ExecuteBytes(ctx map[string]any) ([]byte, error) {
	buffer, err := tpl.newBufferAndExecute(ctx)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Executes the template and returns the rendered template as a string
func (tpl *Template) Execute(ctx map[string]any) (string, error) {
	var b strings.Builder
	err := tpl.execute(ctx, &b)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
