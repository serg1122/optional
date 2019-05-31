package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt64_Create(t *testing.T) {
	opInt64 := OptionalInt64Create()
	assert.IsType(t, opInt64, &OptionalInt64{})
}

func TestOptionalInt64_IsPresent(t *testing.T) {
	opInt64 := OptionalInt64Create()
	assert.False(t, opInt64.IsPresent())
	opInt64.ValueSet(int64(1))
	assert.True(t, opInt64.IsPresent())
}

func TestOptionalInt64_ValueGet(t *testing.T) {
	valueExpected := int64(2)
	opInt64 := OptionalInt64Create()
	_, err1 := opInt64.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt64.ValueSet(valueExpected)
	valueGot, err2 := opInt64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt64_ValueSet(t *testing.T) {
	valueExpected := int64(4)
	opInt64 := OptionalInt64Create()
	opInt64.ValueSet(valueExpected)
	valueGot1, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := int64(4)
	opInt64.ValueSet(valueExpected2)
	valueGot2, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalInt64_MarshalJSON(t *testing.T) {
	opInt64 := OptionalInt64Create()

	bytesGot1, errGot1 := opInt64.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opInt64.ValueSet(int64(6))
	bytesGot2, errGot2 := opInt64.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalInt64_UnmarshalJSON(t *testing.T) {
	opInt64 := OptionalInt64Create()

	err1 := opInt64.UnmarshalJSON([]byte("asd"))
	assert.False(t, opInt64.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opInt64.UnmarshalJSON([]byte("null"))
	assert.False(t, opInt64.IsPresent())
	assert.Nil(t, err2)

	err3 := opInt64.UnmarshalJSON([]byte("7"))
	assert.True(t, opInt64.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot1, int64(7))

	err4 := opInt64.UnmarshalJSON([]byte("8"))
	assert.True(t, opInt64.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot2, int64(8))
}
