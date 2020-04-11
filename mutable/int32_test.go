package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalInt32_(t *testing.T) {
	opInt32 := CreateOptionalInt32()
	assert.IsType(t, opInt32, &OptionalInt32{})
}

func TestOptionalInt32_IsPresent(t *testing.T) {
	opInt32 := CreateOptionalInt32()
	assert.False(t, opInt32.IsPresent())
	opInt32.SetValue(int32(1))
	assert.True(t, opInt32.IsPresent())
}

func TestOptionalInt32_GetValue(t *testing.T) {
	valueExpected := int32(2)
	opInt32 := CreateOptionalInt32()
	_, err1 := opInt32.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opInt32.SetValue(valueExpected)
	valueGot, err2 := opInt32.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt32_SetValue(t *testing.T) {
	valueExpected := int32(3)
	opInt32 := CreateOptionalInt32()
	opInt32.SetValue(valueExpected)
	valueGot1, _ := opInt32.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := int32(4)
	opInt32.SetValue(valueExpected2)
	valueGot2, _ := opInt32.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalInt32_MarshalJSON(t *testing.T) {
	opInt32 := CreateOptionalInt32()

	bytesGot1, errGot1 := opInt32.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opInt32.SetValue(int32(6))
	bytesGot2, errGot2 := opInt32.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalInt32_UnmarshalJSON(t *testing.T) {
	opInt32 := CreateOptionalInt32()

	err1 := opInt32.UnmarshalJSON([]byte("asd"))
	assert.False(t, opInt32.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opInt32.UnmarshalJSON([]byte("null"))
	assert.False(t, opInt32.IsPresent())
	assert.Nil(t, err2)

	err3 := opInt32.UnmarshalJSON([]byte("7"))
	assert.True(t, opInt32.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opInt32.GetValue()
	assert.Equal(t, valueGot1, int32(7))

	err4 := opInt32.UnmarshalJSON([]byte("8"))
	assert.True(t, opInt32.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opInt32.GetValue()
	assert.Equal(t, valueGot2, int32(8))
}
