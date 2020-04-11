package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalFloat32_(t *testing.T) {
	opFloat32 := CreateOptionalFloat32()
	assert.IsType(t, opFloat32, &OptionalFloat32{})
}

func TestOptionalFloat32_IsPresent(t *testing.T) {
	opFloat32 := CreateOptionalFloat32()
	assert.False(t, opFloat32.IsPresent())
	opFloat32.SetValue(float32(1.1))
	assert.True(t, opFloat32.IsPresent())
}

func TestOptionalFloat32_GetValue(t *testing.T) {
	valueExpected := float32(1.2)
	opFloat32 := CreateOptionalFloat32()
	_, err1 := opFloat32.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opFloat32.SetValue(valueExpected)
	valueGot, err2 := opFloat32.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalFloat32_SetValue(t *testing.T) {
	valueExpected := float32(2.3)
	opFloat32 := CreateOptionalFloat32()
	err1 := opFloat32.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opFloat32.SetValue(float32(3.1))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptionalFloat32_MarshalJSON(t *testing.T) {

	opFloat32 := CreateOptionalFloat32()

	bytes, err := opFloat32.MarshalJSON()
	assert.Equal(t, []byte("null"), bytes)
	assert.Nil(t, err)

	opFloat32.SetValue(float32(5.2))
	bytes2, err2 := opFloat32.MarshalJSON()
	assert.Equal(t, []byte("5.2"), bytes2)
	assert.Nil(t, err2)
}

func TestOptionalFloat32_UnmarshalJSON(t *testing.T) {

	opFloat32 := CreateOptionalFloat32()

	err1 := opFloat32.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opFloat32.IsPresent())

	err2 := opFloat32.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opFloat32.IsPresent())

	valueExpected := float32(6.3)

	err3 := opFloat32.UnmarshalJSON([]byte("6.3"))
	assert.Nil(t, err3)
	valueGot1, _ := opFloat32.GetValue()
	assert.Equal(t, valueExpected, valueGot1)

	err4 := opFloat32.UnmarshalJSON([]byte("7.1"))
	assert.IsType(t, err4, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}
