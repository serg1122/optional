package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint_Create(t *testing.T) {
	opUint := OptionalUintCreate()
	assert.IsType(t, opUint, &OptionalUint{})
}

func TestOptionalUint_IsPresent(t *testing.T) {
	opUint := OptionalUintCreate()
	assert.False(t, opUint.IsPresent())
	opUint.ValueSet(uint(1))
	assert.True(t, opUint.IsPresent())
}

func (o *OptionalUint) TestOptionalUint_ValueGet(t *testing.T) {
	valueExpected := uint(2)
	opUint := OptionalUintCreate()
	_, err1 := opUint.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opUint.ValueSet(valueExpected)
	valueGot, err2 := opUint.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint_ValueSet(t *testing.T) {
	opUint := OptionalUintCreate()
	valueExpected1 := uint(3)
	opUint.ValueSet(valueExpected1)
	valueGot1, _ := opUint.ValueGet()
	assert.Equal(t, valueGot1, valueExpected1)
	valueExpected2 := uint(4)
	opUint.ValueSet(valueExpected2)
	valueGot2, _ := opUint.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalUint_MarshalJSON(t *testing.T) {
	opUint := OptionalUintCreate()

	bytesGot1, errGot1 := opUint.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opUint.ValueSet(uint(6))
	bytesGot2, errGot2 := opUint.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalUint_UnmarshalJSON(t *testing.T) {
	opUint := OptionalUintCreate()

	err1 := opUint.UnmarshalJSON([]byte("asd"))
	assert.False(t, opUint.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opUint.UnmarshalJSON([]byte("null"))
	assert.False(t, opUint.IsPresent())
	assert.Nil(t, err2)

	err3 := opUint.UnmarshalJSON([]byte("7"))
	assert.True(t, opUint.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opUint.ValueGet()
	assert.Equal(t, valueGot1, uint(7))

	err4 := opUint.UnmarshalJSON([]byte("8"))
	assert.True(t, opUint.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opUint.ValueGet()
	assert.Equal(t, valueGot2, uint(8))
}
