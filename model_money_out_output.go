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

// MoneyOutOutput struct for MoneyOutOutput
type MoneyOutOutput struct {
	Transaction *TransactionOut `json:"transaction,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewMoneyOutOutput instantiates a new MoneyOutOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyOutOutput() *MoneyOutOutput {
	this := MoneyOutOutput{}
	return &this
}

// NewMoneyOutOutputWithDefaults instantiates a new MoneyOutOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyOutOutputWithDefaults() *MoneyOutOutput {
	this := MoneyOutOutput{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *MoneyOutOutput) GetTransaction() TransactionOut {
	if o == nil || o.Transaction == nil {
		var ret TransactionOut
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyOutOutput) GetTransactionOk() (*TransactionOut, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *MoneyOutOutput) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given TransactionOut and assigns it to the Transaction field.
func (o *MoneyOutOutput) SetTransaction(v TransactionOut) {
	o.Transaction = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MoneyOutOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyOutOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MoneyOutOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *MoneyOutOutput) SetError(v Error) {
	o.Error = &v
}

func (o MoneyOutOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyOutOutput struct {
	value *MoneyOutOutput
	isSet bool
}

func (v NullableMoneyOutOutput) Get() *MoneyOutOutput {
	return v.value
}

func (v *NullableMoneyOutOutput) Set(val *MoneyOutOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyOutOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyOutOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyOutOutput(val *MoneyOutOutput) *NullableMoneyOutOutput {
	return &NullableMoneyOutOutput{value: val, isSet: true}
}

func (v NullableMoneyOutOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyOutOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


