package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalUint struct {
	isPresent bool
	value     uint
}

func OptionalUintCreate() *OptionalUint {
	return &OptionalUint{
		isPresent: false,
	}
}

func (o *OptionalUint) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint) ValueGet() (uint, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint) ValueSet(value uint) {
	o.value = value
	o.isPresent = true
}

func (o *OptionalUint) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalUint) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value uint
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
