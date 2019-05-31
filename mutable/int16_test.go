package mutable

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

func TestOptionalint16_ValueGet(t *testing.T) {
	valueExpected := int16(2)
	opInt16 := OptionalInt16Create()
	_, err1 := opInt16.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt16.ValueSet(valueExpected)
	valueGot, err2 := opInt16.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt16_ValueSet(t *testing.T) {
	valueExpected := int16(3)
	opInt16 := OptionalInt16Create()
	opInt16.ValueSet(valueExpected)
	valueGot, _ := opInt16.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	valueExpected2 := int16(4)
	opInt16.ValueSet(valueExpected2)
	valueGot2, _ := opInt16.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalInt16_MarshalJSON(t *testing.T) {
	opInt16 := OptionalInt16Create()

	bytesGot1, errGot1 := opInt16.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opInt16.ValueSet(int16(6))
	bytesGot2, errGot2 := opInt16.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalInt16_UnmarshalJSON(t *testing.T) {
	opInt16 := OptionalInt16Create()

	err1 := opInt16.UnmarshalJSON([]byte("asd"))
	assert.False(t, opInt16.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opInt16.UnmarshalJSON([]byte("null"))
	assert.False(t, opInt16.IsPresent())
	assert.Nil(t, err2)

	err3 := opInt16.UnmarshalJSON([]byte("7"))
	assert.True(t, opInt16.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opInt16.ValueGet()
	assert.Equal(t, valueGot1, int16(7))

	err4 := opInt16.UnmarshalJSON([]byte("8"))
	assert.True(t, opInt16.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opInt16.ValueGet()
	assert.Equal(t, valueGot2, int16(8))
}
