package immutable

import (
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalString struct {
	isPresent bool
	value     string
}

func OptionalStringCreate() *OptionalString {
	return &OptionalString{
		isPresent: false,
	}
}

func (o *OptionalString) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalString) ValueGet() (string, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return "", optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalString) ValueSet(value string) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o *OptionalString) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalString) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	if string(data) == "null" {
		return nil
	}
	var value string
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.ValueSet(value); err2 != nil {
		return err2
	}
	return nil
}
