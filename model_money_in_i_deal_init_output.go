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

// MoneyInIDealInitOutput struct for MoneyInIDealInitOutput
type MoneyInIDealInitOutput struct {
	// Transaction ID. You will this this value to confirm the transaction
	Id *int64 `json:"id,omitempty"`
	// Redirected URL for the client on the iDeal page payment
	ActionUrl *string `json:"actionUrl,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewMoneyInIDealInitOutput instantiates a new MoneyInIDealInitOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInIDealInitOutput() *MoneyInIDealInitOutput {
	this := MoneyInIDealInitOutput{}
	return &this
}

// NewMoneyInIDealInitOutputWithDefaults instantiates a new MoneyInIDealInitOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInIDealInitOutputWithDefaults() *MoneyInIDealInitOutput {
	this := MoneyInIDealInitOutput{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MoneyInIDealInitOutput) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInIDealInitOutput) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MoneyInIDealInitOutput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MoneyInIDealInitOutput) SetId(v int64) {
	o.Id = &v
}

// GetActionUrl returns the ActionUrl field value if set, zero value otherwise.
func (o *MoneyInIDealInitOutput) GetActionUrl() string {
	if o == nil || o.ActionUrl == nil {
		var ret string
		return ret
	}
	return *o.ActionUrl
}

// GetActionUrlOk returns a tuple with the ActionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInIDealInitOutput) GetActionUrlOk() (*string, bool) {
	if o == nil || o.ActionUrl == nil {
		return nil, false
	}
	return o.ActionUrl, true
}

// HasActionUrl returns a boolean if a field has been set.
func (o *MoneyInIDealInitOutput) HasActionUrl() bool {
	if o != nil && o.ActionUrl != nil {
		return true
	}

	return false
}

// SetActionUrl gets a reference to the given string and assigns it to the ActionUrl field.
func (o *MoneyInIDealInitOutput) SetActionUrl(v string) {
	o.ActionUrl = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MoneyInIDealInitOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInIDealInitOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MoneyInIDealInitOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *MoneyInIDealInitOutput) SetError(v Error) {
	o.Error = &v
}

func (o MoneyInIDealInitOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ActionUrl != nil {
		toSerialize["actionUrl"] = o.ActionUrl
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyInIDealInitOutput struct {
	value *MoneyInIDealInitOutput
	isSet bool
}

func (v NullableMoneyInIDealInitOutput) Get() *MoneyInIDealInitOutput {
	return v.value
}

func (v *NullableMoneyInIDealInitOutput) Set(val *MoneyInIDealInitOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInIDealInitOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInIDealInitOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInIDealInitOutput(val *MoneyInIDealInitOutput) *NullableMoneyInIDealInitOutput {
	return &NullableMoneyInIDealInitOutput{value: val, isSet: true}
}

func (v NullableMoneyInIDealInitOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInIDealInitOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


