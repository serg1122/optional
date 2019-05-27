package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint64_Create(t *testing.T) {
	opUint64 := OptionalUint64Create()
	assert.IsType(t, opUint64, &OptionalUint64{})
}

func TestOptionalUint64_IsPresent(t *testing.T) {
	opUint64 := OptionalUint64Create()
	assert.False(t, opUint64.IsPresent())
	opUint64.ValueSet(uint64(1))
	assert.True(t, opUint64.IsPresent())
}

func TestOptionalUint64_ValueGet(t *testing.T) {

	opUint64 := OptionalUint64Create()

	_, err1 := opUint64.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint64(2)
	opUint64.ValueSet(valueExpected)
	valueGot, _ := opUint64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
}

func TestOptionalUint64_ValueSet(t *testing.T) {

	opUint64 := OptionalUint64Create()

	valueExpected1 := uint64(3)
	opUint64.ValueSet(valueExpected1)
	valueGot1, _ := opUint64.ValueGet()
	assert.Equal(t, valueGot1, valueExpected1)

	valueExpected2 := uint64(4)
	opUint64.ValueSet(valueExpected2)
	valueGot2, _ := opUint64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
