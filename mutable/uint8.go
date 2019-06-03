package mutable

import (
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalUint8 struct {
	isPresent bool
	value     uint8
}

func OptionalUint8Create() *OptionalUint8 {
	return &OptionalUint8{
		isPresent: false,
	}
}

func (o *OptionalUint8) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint8) ValueGet() (uint8, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint8(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint8) ValueSet(value uint8) {
	o.value = value
	o.isPresent = true
}

func (o *OptionalUint8) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalUint8) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var value uint8
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
