package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalInt_(t *testing.T) {
	opInt := CreateOptionalInt()
	assert.IsType(t, opInt, &OptionalInt{})
}

func TestOptionalInt_IsPresent(t *testing.T) {
	opInt := CreateOptionalInt()
	assert.False(t, opInt.IsPresent())
	opInt.SetValue(123)
	assert.True(t, opInt.IsPresent())
}

func TestOptionalInt_GetValue(t *testing.T) {
	valueExpected := 345
	opInt := CreateOptionalInt()
	_, err1 := opInt.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opInt.SetValue(valueExpected)
	valueGot, err2 := opInt.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptinoalInt_SetValue(t *testing.T) {
	valueExpected := int(567)
	opInt := CreateOptionalInt()
	opInt.SetValue(valueExpected)
	valueGot, _ := opInt.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	valueExpected2 := int(789)
	opInt.SetValue(valueExpected2)
	valueGot2, _ := opInt.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalInt_MarshalJSON(t *testing.T) {
	opInt := CreateOptionalInt()

	bytesGot1, errGot1 := opInt.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opInt.SetValue(int(6))
	bytesGot2, errGot2 := opInt.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalInt_UnmarshalJSON(t *testing.T) {
	opInt := CreateOptionalInt()

	err1 := opInt.UnmarshalJSON([]byte("asd"))
	assert.False(t, opInt.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opInt.UnmarshalJSON([]byte("null"))
	assert.False(t, opInt.IsPresent())
	assert.Nil(t, err2)

	err3 := opInt.UnmarshalJSON([]byte("7"))
	assert.True(t, opInt.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opInt.GetValue()
	assert.Equal(t, valueGot1, int(7))

	err4 := opInt.UnmarshalJSON([]byte("8"))
	assert.True(t, opInt.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opInt.GetValue()
	assert.Equal(t, valueGot2, int(8))
}
