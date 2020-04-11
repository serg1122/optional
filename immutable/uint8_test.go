package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint8_Create(t *testing.T) {
	opUint8 := OptionalUint8Create()
	assert.IsType(t, opUint8, &OptionalUint8{})
}

func TestOptionalUint8_IsPresent(t *testing.T) {
	opUint8 := OptionalUint8Create()
	assert.False(t, opUint8.IsPresent())
	opUint8.ValueSet(uint8(1))
	assert.True(t, opUint8.IsPresent())
}

func TestOptionalUint8_GetValue(t *testing.T) {

	opUint8 := OptionalUint8Create()

	_, err1 := opUint8.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint8(2)
	opUint8.ValueSet(valueExpected)
	valueGot, err2 := opUint8.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint8_ValueSet(t *testing.T) {
	opUint8 := OptionalUint8Create()

	valueExpected := uint8(3)
	err1 := opUint8.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint8.GetValue()
	assert.Equal(t, valueGot1, valueExpected)

	err2 := opUint8.ValueSet(uint8(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opUint8.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalUint8_MarshalJSON(t *testing.T) {
	opUint8 := OptionalUint8Create()

	valueGot1, err1 := opUint8.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opUint8.ValueSet(uint8(5))
	valueGot2, err2 := opUint8.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalUint8_UnmarshalJSON(t *testing.T) {
	opUint8 := OptionalUint8Create()

	err1 := opUint8.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opUint8.IsPresent())

	err2 := opUint8.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opUint8.IsPresent())

	valueExpected := uint8(6)

	err3 := opUint8.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opUint8.GetValue()
	assert.True(t, opUint8.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opUint8.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opUint8.GetValue()
	assert.Equal(t, err4, optional.ErrorValueIsPresentCreate())
	assert.Equal(t, valueGot2, valueExpected)
}
