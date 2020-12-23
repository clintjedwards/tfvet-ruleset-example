package main

import (
	"testing"

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

lolwut = "weow"
happy = "test1"
shouldnt_matter = "lolwut"
`)
	parser := hclparse.NewParser()
	hclFile, err := parser.ParseHCL(hclFileRaw, "test")
	if err != nil {
		t.Error(err)
	}

	hclFileBody := hclFile.Body.(*hclsyntax.Body)

	c := Check{}
	errs := c.Check(hclFileBody)

	assert.Len(t, errs, 1, "only one attr should cause an error")
	assert.Equal(t, 17, int(errs[0].Location.Start.Line), "error location should be on line 2")
	assert.Equal(t, 17, int(errs[0].Location.End.Line), "error location should be on line 2")
}
