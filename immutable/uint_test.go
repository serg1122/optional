package immutable

import (
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
	valueExpected := uint(3)
	err1 := opUint.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opUint.ValueSet(uint(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opUint.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}
