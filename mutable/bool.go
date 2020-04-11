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

func CreateOptionalBool() *OptionalBool {
	return &OptionalBool{
		isPresent: false,
	}
}

func (o OptionalBool) IsPresent() bool {
	return o.isPresent
}

func (o OptionalBool) GetValue() (bool, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return false, optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalBool) SetValue(value bool) {
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
	o.SetValue(value)
	return nil
}
