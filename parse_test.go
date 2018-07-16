package terraform2d_test

import (
	"testing"

	"github.com/miketmoore/terraform2d/assert"
	terraform2d "github.com/miketmoore/terraform2d/csv"
)

func TestParse(t *testing.T) {
	a := "86,86,86,86,\n86,86,86,86,"
	got := terraform2d.ParseCSV(a)
	expected := [][]string{
		[]string{"86,86,86,86"},
		[]string{"86,86,86,86"},
	}
	assert.Ok(t, got != nil)
	assert.Ok(t, len(got) == len(expected))
}
