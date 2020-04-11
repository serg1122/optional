package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalUint_(t *testing.T) {
	opUint := CreateOptionalUint()
	assert.IsType(t, opUint, &OptionalUint{})
}

func TestOptionalUint_IsPresent(t *testing.T) {
	opUint := CreateOptionalUint()
	assert.False(t, opUint.IsPresent())
	opUint.SetValue(uint(1))
	assert.True(t, opUint.IsPresent())
}

func (o *OptionalUint) TestOptionalUint_GetValue(t *testing.T) {
	valueExpected := uint(2)
	opUint := CreateOptionalUint()
	_, err1 := opUint.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opUint.SetValue(valueExpected)
	valueGot, err2 := opUint.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint_SetValue(t *testing.T) {
	opUint := CreateOptionalUint()
	valueExpected1 := uint(3)
	opUint.SetValue(valueExpected1)
	valueGot1, _ := opUint.GetValue()
	assert.Equal(t, valueGot1, valueExpected1)
	valueExpected2 := uint(4)
	opUint.SetValue(valueExpected2)
	valueGot2, _ := opUint.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalUint_MarshalJSON(t *testing.T) {
	opUint := CreateOptionalUint()

	bytesGot1, errGot1 := opUint.MarshalJSON()
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	opUint.SetValue(uint(6))
	bytesGot2, errGot2 := opUint.MarshalJSON()
	assert.Equal(t, bytesGot2, []byte("6"))
	assert.Nil(t, errGot2)
}

func TestOptinalUint_UnmarshalJSON(t *testing.T) {
	opUint := CreateOptionalUint()

	err1 := opUint.UnmarshalJSON([]byte("asd"))
	assert.False(t, opUint.IsPresent())
	assert.IsType(t, err1, &json.SyntaxError{})

	err2 := opUint.UnmarshalJSON([]byte("null"))
	assert.False(t, opUint.IsPresent())
	assert.Nil(t, err2)

	err3 := opUint.UnmarshalJSON([]byte("7"))
	assert.True(t, opUint.IsPresent())
	assert.Nil(t, err3)
	valueGot1, _ := opUint.GetValue()
	assert.Equal(t, valueGot1, uint(7))

	err4 := opUint.UnmarshalJSON([]byte("8"))
	assert.True(t, opUint.IsPresent())
	assert.Nil(t, err4)
	valueGot2, _ := opUint.GetValue()
	assert.Equal(t, valueGot2, uint(8))
}
