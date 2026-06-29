package tree_sitter_jinja2_test

import (
	"testing"

	ts_j2 "github.com/dgethings/tree-sitter-jinja2/bindings/go"
	tree_sitter "github.com/smacker/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(ts_j2.Language())
	if language == nil {
		t.Errorf("Error loading Jinja2 grammar")
	}
}
