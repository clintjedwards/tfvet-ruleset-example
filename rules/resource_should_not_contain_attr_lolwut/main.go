// Package main defines a simple example rule.
package main

import (
	tfvet "github.com/clintjedwards/tfvet-sdk"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Check struct{}

func (c *Check) Check(file *hclsyntax.Body) []tfvet.RuleError {
	lintErrors := []tfvet.RuleError{}

	for _, attribute := range file.Attributes {
		if attribute.Name != "lolwut" {
			continue
		}

		location := tfvet.Location{
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
