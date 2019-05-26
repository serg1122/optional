package immutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt32_Create(t *testing.T) {
	opInt32 := OptionalInt32Create()
	assert.IsType(t, opInt32, &OptionalInt32{})
}

func TestOptionalInt32_IsPresent(t *testing.T) {
	opInt32 := OptionalInt32Create()
	assert.False(t, opInt32.IsPresent())
	opInt32.ValueSet(int32(1))
	assert.True(t, opInt32.IsPresent())
}

func TestOptionalInt32_ValueGet(t *testing.T) {
	valueExpected := int32(2)
	opInt32 := OptionalInt32Create()
	_, err1 := opInt32.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt32.ValueSet(valueExpected)
	valueGot, err2 := opInt32.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt32_Valueset(t *testing.T) {
	valueExpected := int32(3)
	opInt32 := OptionalInt32Create()
	err1 := opInt32.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opInt32.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opInt32.ValueSet(int32(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opInt32.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}
