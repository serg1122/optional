package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt8 struct {
	isPresent bool
	value     int8
}

func OptionalInt8Create() *OptionalInt8 {
	return &OptionalInt8{
		isPresent: false,
	}
}

func (o *OptionalInt8) IsPresent() bool {
	return o.isPresent
}

func (o OptionalInt8) ValueGet() (int8, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int8(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt8) ValueSet(value int8) {
	o.value = value
	o.isPresent = true
}

func (o *OptionalInt8) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt8) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value int8
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
