package mutable

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
	opFloat32.SetValue(valueExpected)
	valueGot1, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := float32(3.1)
	opFloat32.SetValue(valueExpected2)
	valueGot2, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalFloat32_MarshalJSON(t *testing.T) {
	opFloat32 := CreateOptionalFloat32()

	jsonValueGot1, jsonValueErr1 := opFloat32.MarshalJSON()
	assert.Equal(t, jsonValueGot1, []byte("null"))
	assert.Nil(t, jsonValueErr1)

	opFloat32.SetValue(float32(6.12))
	jsonValueGot2, jsonValueErr2 := opFloat32.MarshalJSON()
	assert.Equal(t, jsonValueGot2, []byte("6.12"))
	assert.Nil(t, jsonValueErr2)
}

func TestOptionalFloat32_UnmarshalJSON(t *testing.T) {
	opFloat32 := CreateOptionalFloat32()

	err1 := opFloat32.UnmarshalJSON([]byte("asd"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opFloat32.IsPresent())

	err2 := opFloat32.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opFloat32.IsPresent())

	err3 := opFloat32.UnmarshalJSON([]byte("7.23"))
	assert.Nil(t, err3)
	assert.True(t, opFloat32.IsPresent())
	valueGot1, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot1, float32(7.23))

	err4 := opFloat32.UnmarshalJSON([]byte("8.34"))
	assert.Nil(t, err4)
	assert.True(t, opFloat32.IsPresent())
	valueGot2, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot2, float32(8.34))
}
