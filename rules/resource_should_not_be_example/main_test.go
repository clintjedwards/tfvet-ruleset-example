package main

import (
	"testing"

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

	c := Check{}
	errs, err := c.Check(hclFileRaw)
	if err != nil {
		t.Error(err)
	}

	if !assert.Len(t, errs, 1, "only one resource should cause an error") {
		t.FailNow()
	}
	assert.Equal(t, 2, int(errs[0].Location.Start.Line), "error location should be on line 2")
	assert.Equal(t, 2, int(errs[0].Location.End.Line), "error location should be on line 2")
}
