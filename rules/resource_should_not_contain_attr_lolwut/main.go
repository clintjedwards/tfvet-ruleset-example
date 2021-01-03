// Package main defines a simple example rule.
package main

import (
	tfvet "github.com/clintjedwards/tfvet/sdk"
)

// Check is constructed here so that we can fulfill the interface which requires a check function.
type Check struct{}

// Check contains the logic of the linting rule.
func (c *Check) Check(content []byte) ([]tfvet.RuleError, error) {
	// We declare lintErrors here so that we can append to it as we find errors within the file.
	var lintErrors []tfvet.RuleError

	// ParseHCL gives us back a simplified datastructure of the file which we can use to parse
	// through and find linting errors.
	hclContent := tfvet.ParseHCL(content)

	// This is where the actual linting logic is applied. Everytime we find an error we add
	// it to the errors list with its location.
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

		// Every error we find we construct a "RuleError" struct and add it to our errors list.
		lintErrors = append(lintErrors, tfvet.RuleError{
			Suggestion:  "Disable lolwut",
			Remediation: `anything_else = true`,
			Location:    location,
			Metadata: map[string]string{
				"severity": "error",
			},
		})
	}

	return lintErrors, nil
}

func main() {
	// We instantiate an instance of our check interface we filled out above so we can register
	// it into the rule below.
	newCheck := Check{}

	// Here we can fill out more information about the rule, it's purpose, and where to find more
	// documentation.
	// The documentation for each of these fields can be found looking at the sdk documentation
	// here: https://pkg.go.dev/github.com/clintjedwards/tfvet/sdk#Rule
	newRule := &tfvet.Rule{
		Name:  "Lolwut unsafe attribute",
		Short: "Lolwut is inherently unsafe; Use a different attribute",
		Long: `
This is simply a test description of a resource that effectively alerts on nothingness. In turn
this is essentially a really long description so we can test that our descriptions work properly
and are displayed properly in the terminal.
`,
		Enabled: true,
		Link:    "http://lolwut.com/",
		Check:   &newCheck,
	}

	// Lastly we add our new rule so that it is properly registered.
	tfvet.NewRule(newRule)
}
