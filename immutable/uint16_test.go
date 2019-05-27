package immutable

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

	valueExpected := uint16(3)
	err1 := opUint16.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)

	err2 := opUint16.ValueSet(uint16(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}
