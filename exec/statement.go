package exec

/* Incomplete:
   -----------

   verbatim (only the "name" argument is missing for verbatim)

   Reconsideration:
   ----------------

   debug (reason: not sure what to output yet)
   regroup / Grouping on other properties (reason: maybe too python-specific; not sure how useful this would be in Go)

   Following built-in tags wont be added:
   --------------------------------------

   csrf_token (reason: web-framework specific)
   load (reason: python-specific)
   url (reason: web-framework specific)
*/

import (
	"github.com/pkg/errors"

	"github.com/MarioJim/gonja/nodes"
	"github.com/MarioJim/gonja/parser"
)

// This is the function signature of the tag's parser you will have
// to implement in order to create a new tag.
//
// 'doc' is providing access to the whole document while 'arguments'
// is providing access to the user's arguments to the tag:
//
//     {% your_tag_name some "arguments" 123 %}
//
// start_token will be the *Token with the tag's name in it (here: your_tag_name).
//
// Please see the Parser documentation on how to use the parser.
// See RegisterTag()'s documentation for more information about
// writing a tag as well.

type Statement interface {
	nodes.Statement
	Execute(*Renderer, *nodes.StatementBlock) error
}

type StatementSet map[string]parser.StatementParser

// Exists returns true if the given test is already registered
func (ss StatementSet) Exists(name string) bool {
	_, existing := ss[name]
	return existing
}

// Registers a new tag. You usually want to call this
// function in the tag's init() function:
// http://golang.org/doc/effective_go.html#init
func (ss *StatementSet) Register(name string, parser parser.StatementParser) error {
	if ss.Exists(name) {
		return errors.Errorf("Statement '%s' is already registered", name)
	}
	(*ss)[name] = parser
	return nil
}

// Replaces an already registered tag with a new implementation. Use this
// function with caution since it allows you to change existing tag behaviour.
func (ss *StatementSet) Replace(name string, parser parser.StatementParser) error {
	if !ss.Exists(name) {
		return errors.Errorf("Statement '%s' does not exist (therefore cannot be overridden)", name)
	}
	(*ss)[name] = parser
	return nil
}

func (ss *StatementSet) Update(other StatementSet) StatementSet {
	for name, parser := range other {
		(*ss)[name] = parser
	}
	return *ss
}
