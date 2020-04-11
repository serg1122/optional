package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalString_(t *testing.T) {
	opStr := CreateOptionalString()
	assert.IsType(t, opStr, &OptionalString{})
}

func TestOptionalString_IsPresent(t *testing.T) {
	opStr := CreateOptionalString()
	assert.False(t, opStr.IsPresent())
	opStr.SetValue("asd")
	assert.True(t, opStr.IsPresent())
}

func TestOptionalString_GetValue(t *testing.T) {
	valueExpected := "GetValue expected value"
	opStr := CreateOptionalString()
	_, err1 := opStr.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opStr.SetValue(valueExpected)
	valueGot, err2 := opStr.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalString_SetValue(t *testing.T) {
	valueExpected := "SetValue expected value"
	opStr := CreateOptionalString()
	err1 := opStr.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opStr.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opStr.SetValue("asd")
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opStr.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptionalString_MarshalJSON(t *testing.T) {
	opString := CreateOptionalString()

	valueGot1, err1 := opString.MarshalJSON()
	assert.Equal(t, string(valueGot1), "null")
	assert.Nil(t, err1)

	valueExpected := "qwe"
	opString.SetValue(valueExpected)
	valueGot2, err2 := opString.MarshalJSON()
	assert.Equal(t, valueGot2, []byte(`"`+valueExpected+`"`))
	assert.Nil(t, err2)
}

func TestOptionalString_UnmarshalJSON(t *testing.T) {
	opString := CreateOptionalString()

	err1 := opString.UnmarshalJSON([]byte("false"))
	assert.IsType(t, err1, &json.UnmarshalTypeError{})
	assert.False(t, opString.IsPresent())

	err2 := opString.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opString.IsPresent())

	valueExpected := "Наш президент - вор"
	err3 := opString.UnmarshalJSON([]byte(`"` + valueExpected + `"`))
	assert.Nil(t, err3)
	valueGot1, valueGotErr1 := opString.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, valueGotErr1)

	err4 := opString.UnmarshalJSON([]byte("asd"))
	assert.IsType(t, err4, optional.CreateErrorValueIsPresent())
	valueGot2, valueGotErr2 := opString.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
	assert.Nil(t, valueGotErr2)
}
