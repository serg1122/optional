package immutable

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

	valueExpected := uint16(3)
	err1 := opUint16.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint16.GetValue()
	assert.Equal(t, valueGot1, valueExpected)

	err2 := opUint16.SetValue(uint16(4))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opUint16.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalUint16_MarshalJSON(t *testing.T) {
	opUint16 := CreateOptionalUint16()

	valueGot1, err1 := opUint16.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opUint16.SetValue(uint16(5))
	valueGot2, err2 := opUint16.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalUint16_UnmarshalJSON(t *testing.T) {
	opUint16 := CreateOptionalUint16()

	err1 := opUint16.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opUint16.IsPresent())

	err2 := opUint16.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opUint16.IsPresent())

	valueExpected := uint16(6)

	err3 := opUint16.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opUint16.GetValue()
	assert.True(t, opUint16.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opUint16.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opUint16.GetValue()
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
