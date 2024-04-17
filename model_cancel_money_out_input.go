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

// CancelMoneyOutInput struct for CancelMoneyOutInput
type CancelMoneyOutInput struct {
	// Payment Account ID
	AccountId *string `json:"accountId,omitempty"`
}

// NewCancelMoneyOutInput instantiates a new CancelMoneyOutInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMoneyOutInput() *CancelMoneyOutInput {
	this := CancelMoneyOutInput{}
	return &this
}

// NewCancelMoneyOutInputWithDefaults instantiates a new CancelMoneyOutInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMoneyOutInputWithDefaults() *CancelMoneyOutInput {
	this := CancelMoneyOutInput{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CancelMoneyOutInput) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMoneyOutInput) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CancelMoneyOutInput) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CancelMoneyOutInput) SetAccountId(v string) {
	o.AccountId = &v
}

func (o CancelMoneyOutInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	return json.Marshal(toSerialize)
}

type NullableCancelMoneyOutInput struct {
	value *CancelMoneyOutInput
	isSet bool
}

func (v NullableCancelMoneyOutInput) Get() *CancelMoneyOutInput {
	return v.value
}

func (v *NullableCancelMoneyOutInput) Set(val *CancelMoneyOutInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMoneyOutInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMoneyOutInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelMoneyOutInput(val *CancelMoneyOutInput) *NullableCancelMoneyOutInput {
	return &NullableCancelMoneyOutInput{value: val, isSet: true}
}

func (v NullableCancelMoneyOutInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMoneyOutInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


