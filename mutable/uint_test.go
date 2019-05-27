package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint_Create(t *testing.T) {
	opUint := OptionalUintCreate()
	assert.IsType(t, opUint, &OptionalUint{})
}

func TestOptionalUint_IsPresent(t *testing.T) {
	opUint := OptionalUintCreate()
	assert.False(t, opUint.IsPresent())
	opUint.ValueSet(uint(1))
	assert.True(t, opUint.IsPresent())
}

func (o *OptionalUint) TestOptionalUint_ValueGet(t *testing.T) {
	valueExpected := uint(2)
	opUint := OptionalUintCreate()
	_, err1 := opUint.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opUint.ValueSet(valueExpected)
	valueGot, err2 := opUint.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint_ValueSet(t *testing.T) {
	opUint := OptionalUintCreate()
	valueExpected1 := uint(3)
	opUint.ValueSet(valueExpected1)
	valueGot1, _ := opUint.ValueGet()
	assert.Equal(t, valueGot1, valueExpected1)
	valueExpected2 := uint(4)
	opUint.ValueSet(valueExpected2)
	valueGot2, _ := opUint.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
