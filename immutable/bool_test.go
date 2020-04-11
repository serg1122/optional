package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalBool_(t *testing.T) {
	opBool := CreateOptionalBool()
	assert.IsType(t, opBool, &OptionalBool{})
}

func TestOptionalBool_IsPresent(t *testing.T) {
	opBool := CreateOptionalBool()
	assert.False(t, opBool.IsPresent())
	opBool.SetValue(true)
	assert.True(t, opBool.IsPresent())
}

func TestOptionalBool_GetValue(t *testing.T) {
	valueExpexted := true
	opBool := CreateOptionalBool()
	_, err1 := opBool.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opBool.SetValue(valueExpexted)
	valueGot, err2 := opBool.GetValue()
	assert.Equal(t, valueGot, valueExpexted)
	assert.Nil(t, err2)
}

func TestOptinalBool_SetValue(t *testing.T) {
	opBool := CreateOptionalBool()
	err1 := opBool.SetValue(true)
	assert.Nil(t, err1)
	valueGot, err2 := opBool.GetValue()
	assert.True(t, valueGot)
	assert.Nil(t, err2)
	err3 := opBool.SetValue(false)
	assert.IsType(t, err3, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opBool.GetValue()
	assert.True(t, valueGot2)
}

func TestOptionalBool_MarshalJSON(t *testing.T) {
	opBool := CreateOptionalBool()

	valueGot1, err1 := opBool.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	valueExpexted := true
	opBool.SetValue(valueExpexted)
	valueGot2, err2 := opBool.MarshalJSON()
	assert.Equal(t, []byte("true"), valueGot2)
	assert.Nil(t, err2)
}

func TestOptionalBool_UnmarshalJSON(t *testing.T) {
	opBool := CreateOptionalBool()

	err0 := opBool.UnmarshalJSON([]byte("blah"))
	assert.IsType(t, err0, &json.SyntaxError{})
	assert.False(t, opBool.IsPresent())

	err1 := opBool.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err1)
	assert.False(t, opBool.IsPresent())

	err2 := opBool.UnmarshalJSON([]byte("false"))
	assert.Nil(t, err2)
	valueGot, errorGot := opBool.GetValue()
	assert.False(t, valueGot)
	assert.Nil(t, errorGot)

	err3 := opBool.UnmarshalJSON([]byte("true"))
	assert.IsType(t, err3, optional.CreateErrorValueIsPresent())
}
