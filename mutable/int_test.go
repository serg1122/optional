package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt_Create(t *testing.T) {
	opInt := OptionalIntCreate()
	assert.IsType(t, opInt, &OptionalInt{})
}

func TestOptionalInt_IsPresent(t *testing.T) {
	opInt := OptionalIntCreate()
	assert.False(t, opInt.IsPresent())
	opInt.ValueSet(123)
	assert.True(t, opInt.IsPresent())
}

func TestOptionalInt_ValueGet(t *testing.T) {
	valueExpected := 345
	opInt := OptionalIntCreate()
	_, err1 := opInt.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt.ValueSet(valueExpected)
	valueGot, err2 := opInt.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptinoalInt_ValueSet(t *testing.T) {
	valueExpected := int(567)
	opInt := OptionalIntCreate()
	opInt.ValueSet(valueExpected)
	valueGot, _ := opInt.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	valueExpected2 := int(789)
	opInt.ValueSet(valueExpected2)
	valueGot2, _ := opInt.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
