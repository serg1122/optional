package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalBool struct {
	isPresent bool
	value     bool
}

func OptionalBoolCreate() *OptionalBool {
	return &OptionalBool{
		isPresent: false,
	}
}

func (o OptionalBool) IsPresent() bool {
	return o.isPresent
}

func (o OptionalBool) ValueGet() (bool, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return false, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalBool) ValueSet(value bool) {
	o.value = value
	o.isPresent = true
}

func (o OptionalBool) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value bool
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
