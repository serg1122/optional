package mutable

import (
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

func TestOptionalFloat32_ValueGet(t *testing.T) {
	valueExpected := float32(1.2)
	opFloat32 := OptionalFloat32Create()
	_, err1 := opFloat32.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opFloat32.ValueSet(valueExpected)
	valueGot, err2 := opFloat32.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalFloat32_ValueSet(t *testing.T) {
	valueExpected := float32(2.3)
	opFloat32 := OptionalFloat32Create()
	opFloat32.ValueSet(valueExpected)
	valueGot1, _ := opFloat32.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := float32(3.1)
	opFloat32.ValueSet(valueExpected2)
	valueGot2, _ := opFloat32.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
