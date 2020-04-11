package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint64_Create(t *testing.T) {
	opUint64 := OptionalUint64Create()
	assert.IsType(t, opUint64, &OptionalUint64{})
}

func TestOptionalUint64_IsPresent(t *testing.T) {
	opUint64 := OptionalUint64Create()
	assert.False(t, opUint64.IsPresent())
	opUint64.ValueSet(uint64(1))
	assert.True(t, opUint64.IsPresent())
}

func TestOptionalUint64_GetValue(t *testing.T) {

	opUint64 := OptionalUint64Create()

	_, err1 := opUint64.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint64(2)
	opUint64.ValueSet(valueExpected)
	valueGot, _ := opUint64.GetValue()
	assert.Equal(t, valueGot, valueExpected)
}

func TestOptionalUint64_ValueSet(t *testing.T) {

	opUint64 := OptionalUint64Create()

	valueExpected1 := uint64(3)
	opUint64.ValueSet(valueExpected1)
	valueGot1, _ := opUint64.GetValue()
	assert.Equal(t, valueGot1, valueExpected1)

	valueExpected2 := uint64(4)
	opUint64.ValueSet(valueExpected2)
	valueGot2, _ := opUint64.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalUint64_MarshalJSON(t *testing.T) {
	opUint64 := OptionalUint64Create()

	bytesGot1, errGot1 := opUint64.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opUint64.ValueSet(uint64(6))
	bytesGot2, errGot2 := opUint64.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalUint64_UnmarshalJSON(t *testing.T) {
	opUint64 := OptionalUint64Create()

	err1 := opUint64.UnmarshalJSON([]byte("asd"))
	assert.False(t, opUint64.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opUint64.UnmarshalJSON([]byte("null"))
	assert.False(t, opUint64.IsPresent())
	assert.Nil(t, err2)

	err3 := opUint64.UnmarshalJSON([]byte("7"))
	assert.True(t, opUint64.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opUint64.GetValue()
	assert.Equal(t, valueGot1, uint64(7))

	err4 := opUint64.UnmarshalJSON([]byte("8"))
	assert.True(t, opUint64.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opUint64.GetValue()
	assert.Equal(t, valueGot2, uint64(8))
}
