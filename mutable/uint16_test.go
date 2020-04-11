package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalUint16_(t *testing.T) {
	opUint16 := CreateOptionalUint16()
	assert.IsType(t, opUint16, &OptionalUint16{})
}

func TestOptionalUint16_IsPresent(t *testing.T) {
	opUint16 := CreateOptionalUint16()
	assert.False(t, opUint16.IsPresent())
	opUint16.SetValue(uint16(1))
	assert.True(t, opUint16.IsPresent())
}

func TestOptionalUint16_GetValue(t *testing.T) {
	opUint16 := CreateOptionalUint16()

	_, err1 := opUint16.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())

	valueExpected := uint16(2)
	opUint16.SetValue(valueExpected)
	valueGot, err2 := opUint16.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint16_SetValue(t *testing.T) {
	opUint16 := CreateOptionalUint16()

	valueExpected1 := uint16(3)
	opUint16.SetValue(valueExpected1)
	valueGot1, _ := opUint16.GetValue()
	assert.Equal(t, valueGot1, valueExpected1)

	valueExpected2 := uint16(4)
	opUint16.SetValue(valueExpected2)
	valueGot2, _ := opUint16.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalUint16_MarshalJSON(t *testing.T) {
	opUint16 := CreateOptionalUint16()

	bytesGot1, errGot1 := opUint16.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opUint16.SetValue(uint16(6))
	bytesGot2, errGot2 := opUint16.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalUint16_UnmarshalJSON(t *testing.T) {
	opUint16 := CreateOptionalUint16()

	err1 := opUint16.UnmarshalJSON([]byte("asd"))
	assert.False(t, opUint16.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opUint16.UnmarshalJSON([]byte("null"))
	assert.False(t, opUint16.IsPresent())
	assert.Nil(t, err2)

	err3 := opUint16.UnmarshalJSON([]byte("7"))
	assert.True(t, opUint16.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opUint16.GetValue()
	assert.Equal(t, valueGot1, uint16(7))

	err4 := opUint16.UnmarshalJSON([]byte("8"))
	assert.True(t, opUint16.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opUint16.GetValue()
	assert.Equal(t, valueGot2, uint16(8))
}
