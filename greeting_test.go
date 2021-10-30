package gogreeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyGreeter(t *testing.T) {
	testConfigs := []struct {
		Description    string
		Name           string
		ExpectedResult string
		ExpectedError  string
	}{
		{
			Description:   "Zero length name not allowed",
			Name:          "",
			ExpectedError: "zero length name not allowed",
		},
		{
			Description:   "Steve not allowed",
			Name:          "Steve",
			ExpectedError: "name not allowed (Steve)",
		},
		{
			Description:    "Amber is okay",
			Name:           "Amber",
			ExpectedResult: "Hello Amber",
		},
	}
	for _, c := range testConfigs {
		g := MyGreeter{}
		t.Run(c.Description, func(t *testing.T) {
			result, err := g.Hello(c.Name)
			if c.ExpectedResult != "" {
				assert.Equal(t, c.ExpectedResult, result)
			}
			if c.ExpectedError != "" {
				assert.EqualError(t, err, c.ExpectedError)
			}

		})
	}
}
