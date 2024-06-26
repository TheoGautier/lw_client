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

// CancellationToken struct for CancellationToken
type CancellationToken struct {
	IsCancellationRequested *bool `json:"IsCancellationRequested,omitempty"`
	CanBeCanceled *bool `json:"CanBeCanceled,omitempty"`
	WaitHandle *WaitHandle `json:"WaitHandle,omitempty"`
}

// NewCancellationToken instantiates a new CancellationToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancellationToken() *CancellationToken {
	this := CancellationToken{}
	return &this
}

// NewCancellationTokenWithDefaults instantiates a new CancellationToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancellationTokenWithDefaults() *CancellationToken {
	this := CancellationToken{}
	return &this
}

// GetIsCancellationRequested returns the IsCancellationRequested field value if set, zero value otherwise.
func (o *CancellationToken) GetIsCancellationRequested() bool {
	if o == nil || o.IsCancellationRequested == nil {
		var ret bool
		return ret
	}
	return *o.IsCancellationRequested
}

// GetIsCancellationRequestedOk returns a tuple with the IsCancellationRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancellationToken) GetIsCancellationRequestedOk() (*bool, bool) {
	if o == nil || o.IsCancellationRequested == nil {
		return nil, false
	}
	return o.IsCancellationRequested, true
}

// HasIsCancellationRequested returns a boolean if a field has been set.
func (o *CancellationToken) HasIsCancellationRequested() bool {
	if o != nil && o.IsCancellationRequested != nil {
		return true
	}

	return false
}

// SetIsCancellationRequested gets a reference to the given bool and assigns it to the IsCancellationRequested field.
func (o *CancellationToken) SetIsCancellationRequested(v bool) {
	o.IsCancellationRequested = &v
}

// GetCanBeCanceled returns the CanBeCanceled field value if set, zero value otherwise.
func (o *CancellationToken) GetCanBeCanceled() bool {
	if o == nil || o.CanBeCanceled == nil {
		var ret bool
		return ret
	}
	return *o.CanBeCanceled
}

// GetCanBeCanceledOk returns a tuple with the CanBeCanceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancellationToken) GetCanBeCanceledOk() (*bool, bool) {
	if o == nil || o.CanBeCanceled == nil {
		return nil, false
	}
	return o.CanBeCanceled, true
}

// HasCanBeCanceled returns a boolean if a field has been set.
func (o *CancellationToken) HasCanBeCanceled() bool {
	if o != nil && o.CanBeCanceled != nil {
		return true
	}

	return false
}

// SetCanBeCanceled gets a reference to the given bool and assigns it to the CanBeCanceled field.
func (o *CancellationToken) SetCanBeCanceled(v bool) {
	o.CanBeCanceled = &v
}

// GetWaitHandle returns the WaitHandle field value if set, zero value otherwise.
func (o *CancellationToken) GetWaitHandle() WaitHandle {
	if o == nil || o.WaitHandle == nil {
		var ret WaitHandle
		return ret
	}
	return *o.WaitHandle
}

// GetWaitHandleOk returns a tuple with the WaitHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancellationToken) GetWaitHandleOk() (*WaitHandle, bool) {
	if o == nil || o.WaitHandle == nil {
		return nil, false
	}
	return o.WaitHandle, true
}

// HasWaitHandle returns a boolean if a field has been set.
func (o *CancellationToken) HasWaitHandle() bool {
	if o != nil && o.WaitHandle != nil {
		return true
	}

	return false
}

// SetWaitHandle gets a reference to the given WaitHandle and assigns it to the WaitHandle field.
func (o *CancellationToken) SetWaitHandle(v WaitHandle) {
	o.WaitHandle = &v
}

func (o CancellationToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsCancellationRequested != nil {
		toSerialize["IsCancellationRequested"] = o.IsCancellationRequested
	}
	if o.CanBeCanceled != nil {
		toSerialize["CanBeCanceled"] = o.CanBeCanceled
	}
	if o.WaitHandle != nil {
		toSerialize["WaitHandle"] = o.WaitHandle
	}
	return json.Marshal(toSerialize)
}

type NullableCancellationToken struct {
	value *CancellationToken
	isSet bool
}

func (v NullableCancellationToken) Get() *CancellationToken {
	return v.value
}

func (v *NullableCancellationToken) Set(val *CancellationToken) {
	v.value = val
	v.isSet = true
}

func (v NullableCancellationToken) IsSet() bool {
	return v.isSet
}

func (v *NullableCancellationToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancellationToken(val *CancellationToken) *NullableCancellationToken {
	return &NullableCancellationToken{value: val, isSet: true}
}

func (v NullableCancellationToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancellationToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


