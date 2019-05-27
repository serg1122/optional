package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint16_Create(t *testing.T) {
	opUint16 := OptionalUint16Create()
	assert.IsType(t, opUint16, &OptionalUint16{})
}

func TestOptionalUint16_IsPresent(t *testing.T) {
	opUint16 := OptionalUint16Create()
	assert.False(t, opUint16.IsPresent())
	opUint16.ValueSet(uint16(1))
	assert.True(t, opUint16.IsPresent())
}

func TestOptionalUint16_ValueGet(t *testing.T) {
	opUint16 := OptionalUint16Create()

	_, err1 := opUint16.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint16(2)
	opUint16.ValueSet(valueExpected)
	valueGot, err2 := opUint16.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint16_ValueSet(t *testing.T) {
	opUint16 := OptionalUint16Create()

	valueExpected1 := uint16(3)
	opUint16.ValueSet(valueExpected1)
	valueGot1, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot1, valueExpected1)

	valueExpected2 := uint16(4)
	opUint16.ValueSet(valueExpected2)
	valueGot2, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
