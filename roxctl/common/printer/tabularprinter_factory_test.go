package printer

import (
	"testing"

	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTabularPrinterFactory_CreatePrinter(t *testing.T) {
	cases := map[string]struct {
		t          *TabularPrinterFactory
		format     string
		shouldFail bool
		error      error
	}{
		"should not fail with valid factory and format": {
			t:      &TabularPrinterFactory{},
			format: "csv",
		},
		"should fail with invalid factory": {
			t:          &TabularPrinterFactory{Headers: []string{"a"}},
			format:     "csv",
			shouldFail: true,
			error:      errorhelpers.ErrInvalidArgs,
		},
		"should fail with invalid format": {
			t:          &TabularPrinterFactory{},
			shouldFail: true,
			error:      errorhelpers.ErrInvalidArgs,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			p, err := c.t.CreatePrinter(c.format)
			if c.shouldFail {
				require.Error(t, err)
				assert.Nil(t, p)
				assert.ErrorIs(t, err, c.error)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, p)
			}
		})
	}
}

func TestTabularPrinterFactory_Validate(t *testing.T) {
	cases := map[string]struct {
		t          *TabularPrinterFactory
		shouldFail bool
		error      error
	}{
		"should not fail with empty headers and json path expressions": {
			t: &TabularPrinterFactory{},
		},
		"should fail with more headers than json path expressions": {
			t: &TabularPrinterFactory{
				Headers:               []string{"a", "b", "c"},
				RowJSONPathExpression: "a,b",
			},
			shouldFail: true,
			error:      errorhelpers.ErrInvalidArgs,
		},
		"should fail with no header and header as comment set": {
			t: &TabularPrinterFactory{
				NoHeader:        true,
				HeaderAsComment: true,
			},
			shouldFail: true,
			error:      errorhelpers.ErrInvalidArgs,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := c.t.validate()
			if c.shouldFail {
				require.Error(t, err)
				assert.ErrorIs(t, err, c.error)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}