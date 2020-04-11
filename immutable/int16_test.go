package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt16_Create(t *testing.T) {
	opInt16 := OptionalInt16Create()
	assert.IsType(t, opInt16, &OptionalInt16{})
}

func TestOptionalInt16_IsPresent(t *testing.T) {
	opInt16 := OptionalInt16Create()
	assert.False(t, opInt16.IsPresent())
	opInt16.ValueSet(int16(1))
	assert.True(t, opInt16.IsPresent())
}

func TestOptionalint16_GetValue(t *testing.T) {
	valueExpected := int16(2)
	opInt16 := OptionalInt16Create()
	_, err1 := opInt16.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt16.ValueSet(valueExpected)
	valueGot, err2 := opInt16.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt16_ValueSet(t *testing.T) {
	valueExpected := int16(3)
	opInt16 := OptionalInt16Create()
	opInt16.ValueSet(valueExpected)
	valueGot, _ := opInt16.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	err2 := opInt16.ValueSet(int16(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
}

func TestOptinalInt16_MarshalJSON(t *testing.T) {
	opInt16 := OptionalInt16Create()

	valueGot1, err1 := opInt16.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt16.ValueSet(int16(5))
	valueGot2, err2 := opInt16.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt16_UnmarshalJSON(t *testing.T) {
	opInt16 := OptionalInt16Create()

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
	assert.Equal(t, err4, optional.ErrorValueIsPresentCreate())
	assert.Equal(t, valueGot2, valueExpected)
}
