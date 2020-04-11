package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint32_Create(t *testing.T) {
	opUint32 := OptionalUint32Create()
	assert.IsType(t, opUint32, &OptionalUint32{})
}

func TestOptionalUint32_IsPresent(t *testing.T) {
	opUint32 := OptionalUint32Create()
	assert.False(t, opUint32.IsPresent())
	opUint32.ValueSet(uint32(1))
	assert.True(t, opUint32.IsPresent())
}

func TestOptionalUint32_GetValue(t *testing.T) {
	opUint32 := OptionalUint32Create()

	_, err1 := opUint32.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint32(2)
	opUint32.ValueSet(valueExpected)
	valueGot, err2 := opUint32.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint32_ValueSet(t *testing.T) {
	opUint32 := OptionalUint32Create()

	valueExpected := uint32(3)
	opUint32.ValueSet(valueExpected)
	valueGot1, _ := opUint32.GetValue()
	assert.Equal(t, valueGot1, valueExpected)

	valueExpected2 := uint32(4)
	opUint32.ValueSet(valueExpected2)
	valueGot2, _ := opUint32.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalUint32_MarshalJSON(t *testing.T) {
	opUint32 := OptionalUint32Create()

	bytesGot1, errGot1 := opUint32.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opUint32.ValueSet(uint32(6))
	bytesGot2, errGot2 := opUint32.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalUint32_UnmarshalJSON(t *testing.T) {
	opUint32 := OptionalUint32Create()

	err1 := opUint32.UnmarshalJSON([]byte("asd"))
	assert.False(t, opUint32.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opUint32.UnmarshalJSON([]byte("null"))
	assert.False(t, opUint32.IsPresent())
	assert.Nil(t, err2)

	err3 := opUint32.UnmarshalJSON([]byte("7"))
	assert.True(t, opUint32.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opUint32.GetValue()
	assert.Equal(t, valueGot1, uint32(7))

	err4 := opUint32.UnmarshalJSON([]byte("8"))
	assert.True(t, opUint32.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opUint32.GetValue()
	assert.Equal(t, valueGot2, uint32(8))
}
