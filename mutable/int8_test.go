package mutable

import (
	"encoding/json"
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
	opInt8.ValueSet(valueExpected)
	valueGot1, _ := opInt8.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := int8(4)
	opInt8.ValueSet(valueExpected2)
	valueGot2, _ := opInt8.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalInt8_MarshalJSON(t *testing.T) {
	opInt8 := OptionalInt8Create()

	bytesGot1, errGot1 := opInt8.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opInt8.ValueSet(int8(6))
	bytesGot2, errGot2 := opInt8.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalInt8_UnmarshalJSON(t *testing.T) {
	opInt8 := OptionalInt8Create()

	err1 := opInt8.UnmarshalJSON([]byte("asd"))
	assert.False(t, opInt8.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opInt8.UnmarshalJSON([]byte("null"))
	assert.False(t, opInt8.IsPresent())
	assert.Nil(t, err2)

	err3 := opInt8.UnmarshalJSON([]byte("7"))
	assert.True(t, opInt8.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opInt8.ValueGet()
	assert.Equal(t, valueGot1, int8(7))

	err4 := opInt8.UnmarshalJSON([]byte("8"))
	assert.True(t, opInt8.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opInt8.ValueGet()
	assert.Equal(t, valueGot2, int8(8))
}
