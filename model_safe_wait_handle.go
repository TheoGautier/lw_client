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

// SafeWaitHandle struct for SafeWaitHandle
type SafeWaitHandle struct {
	IsInvalid *bool `json:"IsInvalid,omitempty"`
	IsClosed *bool `json:"IsClosed,omitempty"`
}

// NewSafeWaitHandle instantiates a new SafeWaitHandle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafeWaitHandle() *SafeWaitHandle {
	this := SafeWaitHandle{}
	return &this
}

// NewSafeWaitHandleWithDefaults instantiates a new SafeWaitHandle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafeWaitHandleWithDefaults() *SafeWaitHandle {
	this := SafeWaitHandle{}
	return &this
}

// GetIsInvalid returns the IsInvalid field value if set, zero value otherwise.
func (o *SafeWaitHandle) GetIsInvalid() bool {
	if o == nil || o.IsInvalid == nil {
		var ret bool
		return ret
	}
	return *o.IsInvalid
}

// GetIsInvalidOk returns a tuple with the IsInvalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeWaitHandle) GetIsInvalidOk() (*bool, bool) {
	if o == nil || o.IsInvalid == nil {
		return nil, false
	}
	return o.IsInvalid, true
}

// HasIsInvalid returns a boolean if a field has been set.
func (o *SafeWaitHandle) HasIsInvalid() bool {
	if o != nil && o.IsInvalid != nil {
		return true
	}

	return false
}

// SetIsInvalid gets a reference to the given bool and assigns it to the IsInvalid field.
func (o *SafeWaitHandle) SetIsInvalid(v bool) {
	o.IsInvalid = &v
}

// GetIsClosed returns the IsClosed field value if set, zero value otherwise.
func (o *SafeWaitHandle) GetIsClosed() bool {
	if o == nil || o.IsClosed == nil {
		var ret bool
		return ret
	}
	return *o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeWaitHandle) GetIsClosedOk() (*bool, bool) {
	if o == nil || o.IsClosed == nil {
		return nil, false
	}
	return o.IsClosed, true
}

// HasIsClosed returns a boolean if a field has been set.
func (o *SafeWaitHandle) HasIsClosed() bool {
	if o != nil && o.IsClosed != nil {
		return true
	}

	return false
}

// SetIsClosed gets a reference to the given bool and assigns it to the IsClosed field.
func (o *SafeWaitHandle) SetIsClosed(v bool) {
	o.IsClosed = &v
}

func (o SafeWaitHandle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsInvalid != nil {
		toSerialize["IsInvalid"] = o.IsInvalid
	}
	if o.IsClosed != nil {
		toSerialize["IsClosed"] = o.IsClosed
	}
	return json.Marshal(toSerialize)
}

type NullableSafeWaitHandle struct {
	value *SafeWaitHandle
	isSet bool
}

func (v NullableSafeWaitHandle) Get() *SafeWaitHandle {
	return v.value
}

func (v *NullableSafeWaitHandle) Set(val *SafeWaitHandle) {
	v.value = val
	v.isSet = true
}

func (v NullableSafeWaitHandle) IsSet() bool {
	return v.isSet
}

func (v *NullableSafeWaitHandle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafeWaitHandle(val *SafeWaitHandle) *NullableSafeWaitHandle {
	return &NullableSafeWaitHandle{value: val, isSet: true}
}

func (v NullableSafeWaitHandle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafeWaitHandle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


