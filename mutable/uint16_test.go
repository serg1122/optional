package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint16_Create(t *testing.T) {
	opUint16 := OptionalUint16Create()
	assert.IsType(t, opUint16, &OptionalUint16{})
}

func TestOptionalUint16_IsPresent(t *testing.T) {
	opUint16 := OptionalUint16Create()
	assert.False(t, opUint16.IsPresent())
	opUint16.ValueSet(uint16(1))
	assert.True(t, opUint16.IsPresent())
}

func TestOptionalUint16_ValueGet(t *testing.T) {
	opUint16 := OptionalUint16Create()

	_, err1 := opUint16.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint16(2)
	opUint16.ValueSet(valueExpected)
	valueGot, err2 := opUint16.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint16_ValueSet(t *testing.T) {
	opUint16 := OptionalUint16Create()

	valueExpected1 := uint16(3)
	opUint16.ValueSet(valueExpected1)
	valueGot1, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot1, valueExpected1)

	valueExpected2 := uint16(4)
	opUint16.ValueSet(valueExpected2)
	valueGot2, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalUint16_MarshalJSON(t *testing.T) {
	opUint16 := OptionalUint16Create()

	bytesGot1, errGot1 := opUint16.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opUint16.ValueSet(uint16(6))
	bytesGot2, errGot2 := opUint16.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalUint16_UnmarshalJSON(t *testing.T) {
	opUint16 := OptionalUint16Create()

	err1 := opUint16.UnmarshalJSON([]byte("asd"))
	assert.False(t, opUint16.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opUint16.UnmarshalJSON([]byte("null"))
	assert.False(t, opUint16.IsPresent())
	assert.Nil(t, err2)

	err3 := opUint16.UnmarshalJSON([]byte("7"))
	assert.True(t, opUint16.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot1, uint16(7))

	err4 := opUint16.UnmarshalJSON([]byte("8"))
	assert.True(t, opUint16.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opUint16.ValueGet()
	assert.Equal(t, valueGot2, uint16(8))
}
