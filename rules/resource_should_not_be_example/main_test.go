package main

import (
	"testing"

	tfvet "github.com/clintjedwards/tfvet-sdk"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {

	hclFileRaw := []byte(`
resource "google_compute_instance" "example" {
	provider            = "google"
	name                = "basecoat"
	machine_type        = "f1-micro"
	deletion_protection = "true"
	hostname            = "basecoat.clintjedwards.com"
	metadata            = {}
	labels = {
		"basecoat" = ""
	}
	tags = [
		"basecoat"
	]
}

resource "aws_compute_instance" "not_example" {

}

bogus_attr = "example"
`)
	parser := hclparse.NewParser()
	hclFile, err := parser.ParseHCL(hclFileRaw, "test")
	if err != nil {
		t.Error(err)
	}

	hclFileBody := hclFile.Body.(*hclsyntax.Body)
	content := tfvet.Convert(hclFileBody)

	c := Check{}
	errs := c.Check(content)

	assert.Len(t, errs, 1, "only one resource should cause an error")
	assert.Equal(t, 2, int(errs[0].Location.Start.Line), "error location should be on line 2")
	assert.Equal(t, 2, int(errs[0].Location.End.Line), "error location should be on line 2")
}
