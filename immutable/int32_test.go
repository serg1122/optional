package immutable

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
	err1 := opInt32.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opInt32.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opInt32.SetValue(int32(4))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opInt32.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalInt32_MarshalJSON(t *testing.T) {
	opInt32 := CreateOptionalInt32()

	valueGot1, err1 := opInt32.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt32.SetValue(int32(5))
	valueGot2, err2 := opInt32.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt32_UnmarshalJSON(t *testing.T) {
	opInt32 := CreateOptionalInt32()

	err1 := opInt32.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opInt32.IsPresent())

	err2 := opInt32.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opInt32.IsPresent())

	valueExpected := int32(6)

	err3 := opInt32.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opInt32.GetValue()
	assert.True(t, opInt32.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opInt32.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opInt32.GetValue()
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
