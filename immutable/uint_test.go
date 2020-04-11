package immutable

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
	valueExpected := uint(3)
	err1 := opUint.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opUint.SetValue(uint(4))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opUint.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalUint_MarshalJSON(t *testing.T) {
	opUint := CreateOptionalUint()

	valueGot1, err1 := opUint.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opUint.SetValue(uint(5))
	valueGot2, err2 := opUint.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalUint_UnmarshalJSON(t *testing.T) {
	opUint := CreateOptionalUint()

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
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
