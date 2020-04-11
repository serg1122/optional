package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalInt16_(t *testing.T) {
	opInt16 := CreateOptionalInt16()
	assert.IsType(t, opInt16, &OptionalInt16{})
}

func TestOptionalInt16_IsPresent(t *testing.T) {
	opInt16 := CreateOptionalInt16()
	assert.False(t, opInt16.IsPresent())
	opInt16.SetValue(int16(1))
	assert.True(t, opInt16.IsPresent())
}

func TestOptionalint16_GetValue(t *testing.T) {
	valueExpected := int16(2)
	opInt16 := CreateOptionalInt16()
	_, err1 := opInt16.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opInt16.SetValue(valueExpected)
	valueGot, err2 := opInt16.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt16_SetValue(t *testing.T) {
	valueExpected := int16(3)
	opInt16 := CreateOptionalInt16()
	opInt16.SetValue(valueExpected)
	valueGot, _ := opInt16.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	err2 := opInt16.SetValue(int16(4))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
}

func TestOptinalInt16_MarshalJSON(t *testing.T) {
	opInt16 := CreateOptionalInt16()

	valueGot1, err1 := opInt16.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt16.SetValue(int16(5))
	valueGot2, err2 := opInt16.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt16_UnmarshalJSON(t *testing.T) {
	opInt16 := CreateOptionalInt16()

	err1 := opInt16.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opInt16.IsPresent())

	err2 := opInt16.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opInt16.IsPresent())

	valueExpected := int16(6)

	err3 := opInt16.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opInt16.GetValue()
	assert.True(t, opInt16.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opInt16.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opInt16.GetValue()
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
