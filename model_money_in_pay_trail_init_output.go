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

// MoneyInPayTrailInitOutput struct for MoneyInPayTrailInitOutput
type MoneyInPayTrailInitOutput struct {
	// Transaction ID. You will this this value to confirm the transaction
	Id *int64 `json:"id,omitempty"`
	// Redirected URL for the client on the iDeal page payment
	ActionUrl *string `json:"actionUrl,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewMoneyInPayTrailInitOutput instantiates a new MoneyInPayTrailInitOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInPayTrailInitOutput() *MoneyInPayTrailInitOutput {
	this := MoneyInPayTrailInitOutput{}
	return &this
}

// NewMoneyInPayTrailInitOutputWithDefaults instantiates a new MoneyInPayTrailInitOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInPayTrailInitOutputWithDefaults() *MoneyInPayTrailInitOutput {
	this := MoneyInPayTrailInitOutput{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitOutput) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitOutput) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitOutput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MoneyInPayTrailInitOutput) SetId(v int64) {
	o.Id = &v
}

// GetActionUrl returns the ActionUrl field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitOutput) GetActionUrl() string {
	if o == nil || o.ActionUrl == nil {
		var ret string
		return ret
	}
	return *o.ActionUrl
}

// GetActionUrlOk returns a tuple with the ActionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitOutput) GetActionUrlOk() (*string, bool) {
	if o == nil || o.ActionUrl == nil {
		return nil, false
	}
	return o.ActionUrl, true
}

// HasActionUrl returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitOutput) HasActionUrl() bool {
	if o != nil && o.ActionUrl != nil {
		return true
	}

	return false
}

// SetActionUrl gets a reference to the given string and assigns it to the ActionUrl field.
func (o *MoneyInPayTrailInitOutput) SetActionUrl(v string) {
	o.ActionUrl = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *MoneyInPayTrailInitOutput) SetError(v Error) {
	o.Error = &v
}

func (o MoneyInPayTrailInitOutput) MarshalJSON() ([]byte, error) {
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

type NullableMoneyInPayTrailInitOutput struct {
	value *MoneyInPayTrailInitOutput
	isSet bool
}

func (v NullableMoneyInPayTrailInitOutput) Get() *MoneyInPayTrailInitOutput {
	return v.value
}

func (v *NullableMoneyInPayTrailInitOutput) Set(val *MoneyInPayTrailInitOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInPayTrailInitOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInPayTrailInitOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInPayTrailInitOutput(val *MoneyInPayTrailInitOutput) *NullableMoneyInPayTrailInitOutput {
	return &NullableMoneyInPayTrailInitOutput{value: val, isSet: true}
}

func (v NullableMoneyInPayTrailInitOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInPayTrailInitOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

