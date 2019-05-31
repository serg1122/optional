package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt_Create(t *testing.T) {
	opInt := OptionalIntCreate()
	assert.IsType(t, opInt, &OptionalInt{})
}

func TestOptionalInt_IsPresent(t *testing.T) {
	opInt := OptionalIntCreate()
	assert.False(t, opInt.IsPresent())
	opInt.ValueSet(123)
	assert.True(t, opInt.IsPresent())
}

func TestOptionalInt_ValueGet(t *testing.T) {
	valueExpected := 345
	opInt := OptionalIntCreate()
	_, err1 := opInt.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt.ValueSet(valueExpected)
	valueGot, err2 := opInt.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptinoalInt_ValueSet(t *testing.T) {
	valueExpected := int(567)
	opInt := OptionalIntCreate()
	opInt.ValueSet(valueExpected)
	valueGot, _ := opInt.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	valueExpected2 := int(789)
	opInt.ValueSet(valueExpected2)
	valueGot2, _ := opInt.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalInt_MarshalJSON(t *testing.T) {
	opInt := OptionalIntCreate()

	bytesGot1, errGot1 := opInt.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opInt.ValueSet(int(6))
	bytesGot2, errGot2 := opInt.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalInt_UnmarshalJSON(t *testing.T) {
	opInt := OptionalIntCreate()

	err1 := opInt.UnmarshalJSON([]byte("asd"))
	assert.False(t, opInt.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opInt.UnmarshalJSON([]byte("null"))
	assert.False(t, opInt.IsPresent())
	assert.Nil(t, err2)

	err3 := opInt.UnmarshalJSON([]byte("7"))
	assert.True(t, opInt.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opInt.ValueGet()
	assert.Equal(t, valueGot1, int(7))

	err4 := opInt.UnmarshalJSON([]byte("8"))
	assert.True(t, opInt.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opInt.ValueGet()
	assert.Equal(t, valueGot2, int(8))
}
