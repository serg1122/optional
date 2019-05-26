package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt16_Create(t *testing.T) {
	opInt16 := OptionalInt16Create()
	assert.IsType(t, opInt16, &OptionalInt16{})
}

func TestOptionalInt16_IsPresent(t *testing.T) {
	opInt16 := OptionalInt16Create()
	assert.False(t, opInt16.IsPresent())
	opInt16.ValueSet(int16(1))
	assert.True(t, opInt16.IsPresent())
}

func TestOptionalint16_ValueGet(t *testing.T) {
	valueExpected := int16(2)
	opInt16 := OptionalInt16Create()
	_, err1 := opInt16.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt16.ValueSet(valueExpected)
	valueGot, err2 := opInt16.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt16_ValueSet(t *testing.T) {
	valueExpected := int16(3)
	opInt16 := OptionalInt16Create()
	opInt16.ValueSet(valueExpected)
	valueGot, _ := opInt16.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	valueExpected2 := int16(4)
	opInt16.ValueSet(valueExpected2)
	valueGot2, _ := opInt16.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
