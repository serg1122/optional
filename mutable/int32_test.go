package mutable

import (
	"encoding/json"
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

func TestOptionalInt32_ValueSet(t *testing.T) {
	valueExpected := int32(3)
	opInt32 := OptionalInt32Create()
	opInt32.ValueSet(valueExpected)
	valueGot1, _ := opInt32.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := int32(4)
	opInt32.ValueSet(valueExpected2)
	valueGot2, _ := opInt32.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalInt32_MarshalJSON(t *testing.T) {
	opInt32 := OptionalInt32Create()

	bytesGot1, errGot1 := opInt32.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opInt32.ValueSet(int32(6))
	bytesGot2, errGot2 := opInt32.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalInt32_UnmarshalJSON(t *testing.T) {
	opInt32 := OptionalInt32Create()

	err1 := opInt32.UnmarshalJSON([]byte("asd"))
	assert.False(t, opInt32.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opInt32.UnmarshalJSON([]byte("null"))
	assert.False(t, opInt32.IsPresent())
	assert.Nil(t, err2)

	err3 := opInt32.UnmarshalJSON([]byte("7"))
	assert.True(t, opInt32.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opInt32.ValueGet()
	assert.Equal(t, valueGot1, int32(7))

	err4 := opInt32.UnmarshalJSON([]byte("8"))
	assert.True(t, opInt32.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opInt32.ValueGet()
	assert.Equal(t, valueGot2, int32(8))
}
