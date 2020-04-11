package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_ErrorValueIsNotPresent(t *testing.T) {
	err := CreateErrorValueIsNotPresent()
	assert.IsType(t, err, &ErrorValueIsNotPresent{})
	assert.EqualError(t, err, ErrorValueIsNotPresentMessage)
}

func TestError_ErrorValueIsPresent(t *testing.T) {
	err := CreateErrorValueIsPresent()
	assert.IsType(t, err, &ErrorValueIsPresent{})
	assert.EqualError(t, err, ErrorValueIsPresentMessage)
}
