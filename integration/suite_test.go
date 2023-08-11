package integration_test

import (
	"flag"
	"fmt"
	"testing"
)

var testdataFlag = flag.String("testdata", "./testdata", "root folder for the test data")

func TestTemplates(t *testing.T) {
	// Add a global to the default set
	root := *testdataFlag
	env := testEnv(root)
	env.Globals.Set("this_is_a_global_variable", "this is a global text")
	GlobTemplateTests(t, root, env)
}

func TestExpressions(t *testing.T) {
	root := fmt.Sprintf("%s/expressions", *testdataFlag)
	env := testEnv(root)
	GlobTemplateTests(t, root, env)
}

func TestFilters(t *testing.T) {
	root := fmt.Sprintf("%s/filters", *testdataFlag)
	env := testEnv(root)
	GlobTemplateTests(t, root, env)
}

func TestFunctions(t *testing.T) {
	root := fmt.Sprintf("%s/functions", *testdataFlag)
	env := testEnv(root)
	GlobTemplateTests(t, root, env)
}

func TestTests(t *testing.T) {
	root := fmt.Sprintf("%s/tests", *testdataFlag)
	env := testEnv(root)
	GlobTemplateTests(t, root, env)
}

func TestStatements(t *testing.T) {
	root := fmt.Sprintf("%s/statements", *testdataFlag)
	env := testEnv(root)
	GlobTemplateTests(t, root, env)
}
