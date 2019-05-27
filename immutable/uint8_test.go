package immutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint8_Create(t *testing.T) {
	opUint8 := OptionalUint8Create()
	assert.IsType(t, opUint8, &OptionalUint8{})
}

func TestOptionalUint8_IsPresent(t *testing.T) {
	opUint8 := OptionalUint8Create()
	assert.False(t, opUint8.IsPresent())
	opUint8.ValueSet(uint8(1))
	assert.True(t, opUint8.IsPresent())
}

func TestOptionalUint8_Valueget(t *testing.T) {

	opUint8 := OptionalUint8Create()

	_, err1 := opUint8.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint8(2)
	opUint8.ValueSet(valueExpected)
	valueGot, err2 := opUint8.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint8_ValueSet(t *testing.T) {
	opUint8 := OptionalUint8Create()

	valueExpected := uint8(3)
	err1 := opUint8.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint8.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)

	err2 := opUint8.ValueSet(uint8(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opUint8.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}
