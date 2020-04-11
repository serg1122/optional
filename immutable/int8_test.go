package immutable

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

func TestOptionalInt8_GetValue(t *testing.T) {
	valueExpected := int8(2)
	opInt8 := OptionalInt8Create()
	_, err1 := opInt8.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt8.ValueSet(valueExpected)
	valueGot, err2 := opInt8.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt8_ValueSet(t *testing.T) {
	valueExpected := int8(3)
	opInt8 := OptionalInt8Create()
	err1 := opInt8.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opInt8.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opInt8.ValueSet(int8(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opInt8.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalInt8_MarshalJSON(t *testing.T) {
	opInt8 := OptionalInt8Create()

	valueGot1, err1 := opInt8.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt8.ValueSet(int8(5))
	valueGot2, err2 := opInt8.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt8_UnmarshalJSON(t *testing.T) {
	opInt8 := OptionalInt8Create()

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
	assert.Equal(t, err4, optional.ErrorValueIsPresentCreate())
	assert.Equal(t, valueGot2, valueExpected)
}
