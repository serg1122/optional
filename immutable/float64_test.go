package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalFloat64_(t *testing.T) {
	opFloat64 := CreateOptionalFloat64()
	assert.IsType(t, opFloat64, &OptionalFloat64{})
}

func TestOptionalFloat64_IsPresent(t *testing.T) {
	opFloat64 := CreateOptionalFloat64()
	assert.False(t, opFloat64.IsPresent())
	opFloat64.SetValue(0.12)
	assert.True(t, opFloat64.IsPresent())
}

func TestOptionalFloat64_GetValue(t *testing.T) {
	valueExpected := 1.1
	opFloat64 := CreateOptionalFloat64()
	_, err1 := opFloat64.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opFloat64.SetValue(valueExpected)
	valueGot, err2 := opFloat64.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalFloat64_SetValue(t *testing.T) {
	valueExpected := 2.3
	opFloat64 := CreateOptionalFloat64()
	err1 := opFloat64.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot, err2 := opFloat64.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
	err3 := opFloat64.SetValue(3.1)
	assert.IsType(t, err3, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opFloat64.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptionalFloat64_MarshalJSON(t *testing.T) {
	opFloat64 := CreateOptionalFloat64()

	value1, err1 := opFloat64.MarshalJSON()
	assert.Equal(t, []byte("null"), value1)
	assert.Nil(t, err1)

	opFloat64.SetValue(float64(1.23))
	value2, err2 := opFloat64.MarshalJSON()
	assert.Equal(t, []byte("1.23"), value2)
	assert.Nil(t, err2)
}

func TestOptionalFloat64_UnmarshalJSON(t *testing.T) {
	opFloat64 := CreateOptionalFloat64()

	err1 := opFloat64.UnmarshalJSON([]byte("asd"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opFloat64.IsPresent())

	err2 := opFloat64.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opFloat64.IsPresent())

	err3 := opFloat64.UnmarshalJSON([]byte("3.45"))
	assert.Nil(t, err3)
	assert.True(t, opFloat64.IsPresent())
	valueGot1, _ := opFloat64.GetValue()
	assert.Equal(t, float64(3.45), valueGot1)

	err4 := opFloat64.UnmarshalJSON([]byte("4.56"))
	assert.IsType(t, err4, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opFloat64.GetValue()
	assert.Equal(t, valueGot2, float64(3.45))
}
