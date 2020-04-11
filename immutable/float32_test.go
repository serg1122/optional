package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalFloat32_Create(t *testing.T) {
	opFloat32 := OptionalFloat32Create()
	assert.IsType(t, opFloat32, &OptionalFloat32{})
}

func TestOptionalFloat32_IsPresent(t *testing.T) {
	opFloat32 := OptionalFloat32Create()
	assert.False(t, opFloat32.IsPresent())
	opFloat32.ValueSet(float32(1.1))
	assert.True(t, opFloat32.IsPresent())
}

func TestOptionalFloat32_GetValue(t *testing.T) {
	valueExpected := float32(1.2)
	opFloat32 := OptionalFloat32Create()
	_, err1 := opFloat32.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opFloat32.ValueSet(valueExpected)
	valueGot, err2 := opFloat32.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalFloat32_ValueSet(t *testing.T) {
	valueExpected := float32(2.3)
	opFloat32 := OptionalFloat32Create()
	err1 := opFloat32.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opFloat32.ValueSet(float32(3.1))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptionalFloat32_MarshalJSON(t *testing.T) {

	opFloat32 := OptionalFloat32Create()

	bytes, err := opFloat32.MarshalJSON()
	assert.Equal(t, []byte("null"), bytes)
	assert.Nil(t, err)

	opFloat32.ValueSet(float32(5.2))
	bytes2, err2 := opFloat32.MarshalJSON()
	assert.Equal(t, []byte("5.2"), bytes2)
	assert.Nil(t, err2)
}

func TestOptionalFloat32_UnmarshalJSON(t *testing.T) {

	opFloat32 := OptionalFloat32Create()

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
	assert.IsType(t, err4, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opFloat32.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}
