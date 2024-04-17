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

// UpdateAccountStatusInput struct for UpdateAccountStatusInput
type UpdateAccountStatusInput struct {
	// New payment account status<br/>0 = Unknown.<br/>1 = Not Opened.<br/>2 = Opened, need more documents.<br/>3 = Opened, document rejected.<br/>5 = Opened, KYC1.<br/>6 = Opened, KYC2.<br/>7 = Opened, KYC3.<br/>8 = Opened, document expired.<br/>9 = Frozen (by backoffice).<br/>10 = Blocked.<br/>11 = Locked (by Web Service).<br/>12 = Closed.<br/>13 = Pending KYC3.<br/>14 = One-time customer.<br/>15 = CGE.<br/>16 = Technical Payment Account.<br/>
	Status int32 `json:"status"`
}

// NewUpdateAccountStatusInput instantiates a new UpdateAccountStatusInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountStatusInput(status int32) *UpdateAccountStatusInput {
	this := UpdateAccountStatusInput{}
	this.Status = status
	return &this
}

// NewUpdateAccountStatusInputWithDefaults instantiates a new UpdateAccountStatusInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountStatusInputWithDefaults() *UpdateAccountStatusInput {
	this := UpdateAccountStatusInput{}
	return &this
}

// GetStatus returns the Status field value
func (o *UpdateAccountStatusInput) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountStatusInput) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateAccountStatusInput) SetStatus(v int32) {
	o.Status = v
}

func (o UpdateAccountStatusInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccountStatusInput struct {
	value *UpdateAccountStatusInput
	isSet bool
}

func (v NullableUpdateAccountStatusInput) Get() *UpdateAccountStatusInput {
	return v.value
}

func (v *NullableUpdateAccountStatusInput) Set(val *UpdateAccountStatusInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountStatusInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountStatusInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountStatusInput(val *UpdateAccountStatusInput) *NullableUpdateAccountStatusInput {
	return &NullableUpdateAccountStatusInput{value: val, isSet: true}
}

func (v NullableUpdateAccountStatusInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountStatusInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


