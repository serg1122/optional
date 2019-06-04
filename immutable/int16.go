package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalInt16 struct {
	isPresent bool
	value     int16
}

func OptionalInt16Create() *OptionalInt16 {
	return &OptionalInt16{
		isPresent: false,
	}
}

func (o *OptionalInt16) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalInt16) ValueGet() (int16, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return int16(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalInt16) ValueSet(value int16) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o *OptionalInt16) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalInt16) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value int16
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.ValueSet(value); err2 != nil {
		return err2
	}
	return nil
}
