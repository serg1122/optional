package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalInt8_(t *testing.T) {
	opInt8 := CreateOptionalInt8()
	assert.IsType(t, opInt8, &OptionalInt8{})
}

func TestOptionalInt8_IsPresent(t *testing.T) {
	opInt8 := CreateOptionalInt8()
	assert.False(t, opInt8.IsPresent())
	opInt8.SetValue(int8(1))
	assert.True(t, opInt8.IsPresent())
}

func TestOptionalInt8_GetValue(t *testing.T) {
	valueExpected := int8(2)
	opInt8 := CreateOptionalInt8()
	_, err1 := opInt8.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opInt8.SetValue(valueExpected)
	valueGot, err2 := opInt8.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt8_SetValue(t *testing.T) {
	valueExpected := int8(3)
	opInt8 := CreateOptionalInt8()
	err1 := opInt8.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opInt8.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opInt8.SetValue(int8(4))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opInt8.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalInt8_MarshalJSON(t *testing.T) {
	opInt8 := CreateOptionalInt8()

	valueGot1, err1 := opInt8.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt8.SetValue(int8(5))
	valueGot2, err2 := opInt8.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt8_UnmarshalJSON(t *testing.T) {
	opInt8 := CreateOptionalInt8()

	err1 := opInt8.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opInt8.IsPresent())

	err2 := opInt8.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opInt8.IsPresent())

	valueExpected := int8(6)

	err3 := opInt8.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opInt8.GetValue()
	assert.True(t, opInt8.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opInt8.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opInt8.GetValue()
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
