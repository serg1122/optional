package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint32_Create(t *testing.T) {
	opUint32 := OptinoalUint32Create()
	assert.IsType(t, opUint32, &OptionalUint32{})
}

func TestOptionalUint32_IsPresent(t *testing.T) {
	opUint32 := OptinoalUint32Create()
	assert.False(t, opUint32.IsPresent())
	opUint32.ValueSet(uint32(1))
	assert.True(t, opUint32.IsPresent())
}

func TestOptionalUint32_ValueGet(t *testing.T) {
	opUint32 := OptinoalUint32Create()

	_, err1 := opUint32.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint32(2)
	opUint32.ValueSet(valueExpected)
	valueGot, err2 := opUint32.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint32_ValueSet(t *testing.T) {
	opUint32 := OptinoalUint32Create()

	valueExpected := uint32(3)
	opUint32.ValueSet(valueExpected)
	valueGot1, _ := opUint32.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)

	valueExpected2 := uint32(4)
	opUint32.ValueSet(valueExpected2)
	valueGot2, _ := opUint32.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
