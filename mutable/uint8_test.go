package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalUint8_(t *testing.T) {
	opUint8 := CreateOptionalUint8()
	assert.IsType(t, opUint8, &OptionalUint8{})
}

func TestOptionalUint8_IsPresent(t *testing.T) {
	opUint8 := CreateOptionalUint8()
	assert.False(t, opUint8.IsPresent())
	opUint8.SetValue(uint8(1))
	assert.True(t, opUint8.IsPresent())
}

func TestOptionalUint8_GetValue(t *testing.T) {

	opUint8 := CreateOptionalUint8()

	_, err1 := opUint8.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())

	valueExpected := uint8(2)
	opUint8.SetValue(valueExpected)
	valueGot, err2 := opUint8.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint8_SetValue(t *testing.T) {
	opUint8 := CreateOptionalUint8()

	valueExpected := uint8(3)
	opUint8.SetValue(valueExpected)
	valueGot1, _ := opUint8.GetValue()
	assert.Equal(t, valueGot1, valueExpected)

	valueExpected2 := uint8(4)
	opUint8.SetValue(valueExpected2)
	valueGot2, _ := opUint8.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalUint8_MarshalJSON(t *testing.T) {
	opUint8 := CreateOptionalUint8()

	bytesGot1, errGot1 := opUint8.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opUint8.SetValue(uint8(6))
	bytesGot2, errGot2 := opUint8.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalUint8_UnmarshalJSON(t *testing.T) {
	opUint8 := CreateOptionalUint8()

	err1 := opUint8.UnmarshalJSON([]byte("asd"))
	assert.False(t, opUint8.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opUint8.UnmarshalJSON([]byte("null"))
	assert.False(t, opUint8.IsPresent())
	assert.Nil(t, err2)

	err3 := opUint8.UnmarshalJSON([]byte("7"))
	assert.True(t, opUint8.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opUint8.GetValue()
	assert.Equal(t, valueGot1, uint8(7))

	err4 := opUint8.UnmarshalJSON([]byte("8"))
	assert.True(t, opUint8.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opUint8.GetValue()
	assert.Equal(t, valueGot2, uint8(8))
}
