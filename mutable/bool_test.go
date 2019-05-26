package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalBool_Create(t *testing.T) {
	opBool := OptionalBoolCreate()
	assert.IsType(t, opBool, &OptionalBool{})
}

func TestOptionalBool_IsPresent(t *testing.T) {
	opBool := OptionalBoolCreate()
	assert.False(t, opBool.IsPresent())
	opBool.ValueSet(true)
	assert.True(t, opBool.IsPresent())
}

func TestOptionalBool_ValueGet(t *testing.T) {
	valueExpexted := true
	opBool := OptionalBoolCreate()
	_, err1 := opBool.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opBool.ValueSet(valueExpexted)
	valueGot, err2 := opBool.ValueGet()
	assert.Equal(t, valueGot, valueExpexted)
	assert.Nil(t, err2)
}

func TestOptinalBool_ValueSet(t *testing.T) {
	opBool := OptionalBoolCreate()
	opBool.ValueSet(true)
	valueGot, err2 := opBool.ValueGet()
	assert.True(t, valueGot)
	assert.Nil(t, err2)
	opBool.ValueSet(false)
	valueGot2, _ := opBool.ValueGet()
	assert.False(t, valueGot2)
}
