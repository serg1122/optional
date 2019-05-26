package immutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalFloat64_Create(t *testing.T) {
	opFloat64 := OptionalFloat64Create()
	assert.IsType(t, opFloat64, &OptionalFloat64{})
}

func TestOptionalFloat64_IsPresent(t *testing.T) {
	opFloat64 := OptionalFloat64Create()
	assert.False(t, opFloat64.IsPresent())
	opFloat64.ValueSet(0.12)
	assert.True(t, opFloat64.IsPresent())
}

func TestOptionalFloat64_ValueGet(t *testing.T) {
	valueExpected := 1.1
	opFloat64 := OptionalFloat64Create()
	_, err1 := opFloat64.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opFloat64.ValueSet(valueExpected)
	valueGot, err2 := opFloat64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalFloat64_ValueSet(t *testing.T) {
	valueExpected := 2.3
	opFloat64 := OptionalFloat64Create()
	err1 := opFloat64.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot, err2 := opFloat64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
	err3 := opFloat64.ValueSet(3.1)
	assert.IsType(t, err3, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opFloat64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}
