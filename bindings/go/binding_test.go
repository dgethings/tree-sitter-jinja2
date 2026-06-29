package tree_sitter_jinja2_test

import (
	"testing"

	"github.com/dgethings/tree-sitter-jinja2"
	tree_sitter "github.com/smacker/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_jinja2.Language())
	if language == nil {
		t.Errorf("Error loading Jinja2 grammar")
	}
}
