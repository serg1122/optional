package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt64_Create(t *testing.T) {
	opInt64 := OptionalInt64Create()
	assert.IsType(t, opInt64, &OptionalInt64{})
}

func TestOptionalInt64_IsPresent(t *testing.T) {
	opInt64 := OptionalInt64Create()
	assert.False(t, opInt64.IsPresent())
	opInt64.ValueSet(int64(1))
	assert.True(t, opInt64.IsPresent())
}

func TestOptionalInt64_ValueGet(t *testing.T) {
	valueExpected := int64(2)
	opInt64 := OptionalInt64Create()
	_, err1 := opInt64.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt64.ValueSet(valueExpected)
	valueGot, err2 := opInt64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt64_ValueSet(t *testing.T) {
	valueExpected := int64(4)
	opInt64 := OptionalInt64Create()
	opInt64.ValueSet(valueExpected)
	valueGot1, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := int64(4)
	opInt64.ValueSet(valueExpected2)
	valueGot2, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
