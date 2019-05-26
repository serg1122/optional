package immutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt8_Create(t *testing.T) {
	opInt8 := OptionalInt8Create()
	assert.IsType(t, opInt8, &OptionalInt8{})
}

func TestOptionalInt8_IsPresent(t *testing.T) {
	opInt8 := OptionalInt8Create()
	assert.False(t, opInt8.IsPresent())
	opInt8.ValueSet(int8(1))
	assert.True(t, opInt8.IsPresent())
}

func TestOptionalInt8_ValueGet(t *testing.T) {
	valueExpected := int8(2)
	opInt8 := OptionalInt8Create()
	_, err1 := opInt8.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt8.ValueSet(valueExpected)
	valueGot, err2 := opInt8.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt8_ValueSet(t *testing.T) {
	valueExpected := int8(3)
	opInt8 := OptionalInt8Create()
	err1 := opInt8.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opInt8.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opInt8.ValueSet(int8(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opInt8.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}
