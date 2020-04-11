package immutable

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

func (o *OptionalUint) TestOptionalUint_GetValue(t *testing.T) {
	valueExpected := uint(2)
	opUint := OptionalUintCreate()
	_, err1 := opUint.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opUint.ValueSet(valueExpected)
	valueGot, err2 := opUint.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint_ValueSet(t *testing.T) {
	opUint := OptionalUintCreate()
	valueExpected := uint(3)
	err1 := opUint.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opUint.ValueSet(uint(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opUint.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalUint_MarshalJSON(t *testing.T) {
	opUint := OptionalUintCreate()

	valueGot1, err1 := opUint.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opUint.ValueSet(uint(5))
	valueGot2, err2 := opUint.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalUint_UnmarshalJSON(t *testing.T) {
	opUint := OptionalUintCreate()

	err1 := opUint.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opUint.IsPresent())

	err2 := opUint.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opUint.IsPresent())

	valueExpected := uint(6)

	err3 := opUint.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opUint.GetValue()
	assert.True(t, opUint.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opUint.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opUint.GetValue()
	assert.Equal(t, err4, optional.ErrorValueIsPresentCreate())
	assert.Equal(t, valueGot2, valueExpected)
}
