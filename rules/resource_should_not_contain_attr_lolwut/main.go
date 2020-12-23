// Package main defines a simple example rule.
package main

import (
	"log"

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

	for _, attribute := range hclContent.Attributes {
		if attribute.Name != "lolwut" {
			continue
		}

		location := tfvet.Range{
			Start: tfvet.Position{
				Line:   uint32(attribute.NameRange.Start.Line),
				Column: uint32(attribute.NameRange.Start.Column),
			},
			End: tfvet.Position{
				Line:   uint32(attribute.NameRange.End.Line),
				Column: uint32(attribute.NameRange.End.Column),
			},
		}

		lintErrors = append(lintErrors, tfvet.RuleError{
			Location: location,
		})
	}

	return lintErrors
}

func main() {
	newCheck := Check{}

	newRule := &tfvet.Rule{
		Name:  "Resource should not contain attribute lolwut",
		Short: "Lolwut is inherently unsafe; see link for more details",
		Long: `
This is simply a test description of a resource that effectively alerts on nothingness. In turn
this is essentially a really long description so we can test that our descriptions work properly
and are displayed properly in the terminal.
`,
		Default:  false,
		Severity: tfvet.Error,
		Link:     "http://lolwut.com/",
		Check:    &newCheck,
	}

	tfvet.NewRule(newRule)
}
