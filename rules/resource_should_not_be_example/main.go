// Package main defines a simple example rule.
package main

import (
	"log"
	"strings"

	tfvet "github.com/clintjedwards/tfvet-sdk"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Check struct{}

func (c *Check) Check(content []byte) []tfvet.RuleError {
	//TODO(clintjedwards): Having to reparse the file for every plugin is very slow, figure
	// out if there is a better way to transfer this information to plugins

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(content, "tmp")
	if diags.HasErrors() {
		log.Fatal(diags)
	}

	hclContent := file.Body.(*hclsyntax.Body)

	lintErrors := []tfvet.RuleError{}

	for _, block := range hclContent.Blocks {
		for _, label := range block.Labels {
			if strings.ToLower(label) != "example" {
				continue
			}

			location := tfvet.Range{
				Start: tfvet.Position{
					Line:   uint32(block.DefRange().Start.Line),
					Column: uint32(block.DefRange().Start.Column),
				},
				End: tfvet.Position{
					Line:   uint32(block.DefRange().End.Line),
					Column: uint32(block.DefRange().End.Column),
				},
			}

			lintErrors = append(lintErrors, tfvet.RuleError{
				Location: location,
			})
		}
	}

	return lintErrors
}

func main() {
	newCheck := Check{}

	newRule := &tfvet.Rule{
		Name:  "No resource with the name 'example'",
		Short: "Example is a poor name for a resources and might lead to naming collisions.",
		Long: `
This is simply a test description of a resource that effectively alerts on nothingness. In turn
this is essentially a really long description so we can test that our descriptions work properly
and are displayed properly in the terminal.
`,
		Default:  true,
		Severity: tfvet.Error,
		Link:     "https://google.com",
		Check:    &newCheck,
	}

	tfvet.NewRule(newRule)
}
