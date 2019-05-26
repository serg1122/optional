package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalString_Create(t *testing.T) {
	opStr := OptionalStringCreate()
	assert.IsType(t, opStr, &OptionalString{})
}

func TestOptionalString_IsPresent(t *testing.T) {
	opStr := OptionalStringCreate()
	assert.False(t, opStr.IsPresnt())
	opStr.ValueSet("asd")
	assert.True(t, opStr.IsPresnt())
}

func TestOptionalString_ValueGet(t *testing.T) {
	valueExpected := "ValueGet expected value"
	opStr := OptionalStringCreate()
	_, err1 := opStr.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opStr.ValueSet(valueExpected)
	valueGot, err2 := opStr.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalString_ValueSet(t *testing.T) {
	valueExpected := "ValueSet expected value"
	opStr := OptionalStringCreate()
	opStr.ValueSet(valueExpected)
	valueGot1, _ := opStr.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := "asd"
	opStr.ValueSet(valueExpected2)
	valueGot2, _ := opStr.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
