package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt struct {
	isPresent bool
	value     int
}

func OptionalIntCreate() *OptionalInt {
	return &OptionalInt{
		isPresent: false,
	}
}

func (o *OptionalInt) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalInt) ValueGet() (int, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt) ValueSet(value int) {
	o.value = value
	o.isPresent = true
}

func (o *OptionalInt) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value int
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
