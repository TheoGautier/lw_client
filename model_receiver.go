/*
Lemonway DirectKit API 2.0

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Receiver Receiver
type Receiver struct {
	// Full name
	FullName string `json:"fullName"`
}

// NewReceiver instantiates a new Receiver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiver(fullName string) *Receiver {
	this := Receiver{}
	this.FullName = fullName
	return &this
}

// NewReceiverWithDefaults instantiates a new Receiver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiverWithDefaults() *Receiver {
	this := Receiver{}
	return &this
}

// GetFullName returns the FullName field value
func (o *Receiver) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *Receiver) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *Receiver) SetFullName(v string) {
	o.FullName = v
}

func (o Receiver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fullName"] = o.FullName
	}
	return json.Marshal(toSerialize)
}

type NullableReceiver struct {
	value *Receiver
	isSet bool
}

func (v NullableReceiver) Get() *Receiver {
	return v.value
}

func (v *NullableReceiver) Set(val *Receiver) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiver) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiver(val *Receiver) *NullableReceiver {
	return &NullableReceiver{value: val, isSet: true}
}

func (v NullableReceiver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


