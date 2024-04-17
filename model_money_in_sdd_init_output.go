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

// MoneyInSddInitOutput struct for MoneyInSddInitOutput
type MoneyInSddInitOutput struct {
	Transaction *TransactionIn `json:"transaction,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewMoneyInSddInitOutput instantiates a new MoneyInSddInitOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInSddInitOutput() *MoneyInSddInitOutput {
	this := MoneyInSddInitOutput{}
	return &this
}

// NewMoneyInSddInitOutputWithDefaults instantiates a new MoneyInSddInitOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInSddInitOutputWithDefaults() *MoneyInSddInitOutput {
	this := MoneyInSddInitOutput{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *MoneyInSddInitOutput) GetTransaction() TransactionIn {
	if o == nil || o.Transaction == nil {
		var ret TransactionIn
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitOutput) GetTransactionOk() (*TransactionIn, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *MoneyInSddInitOutput) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given TransactionIn and assigns it to the Transaction field.
func (o *MoneyInSddInitOutput) SetTransaction(v TransactionIn) {
	o.Transaction = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MoneyInSddInitOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInSddInitOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MoneyInSddInitOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *MoneyInSddInitOutput) SetError(v Error) {
	o.Error = &v
}

func (o MoneyInSddInitOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyInSddInitOutput struct {
	value *MoneyInSddInitOutput
	isSet bool
}

func (v NullableMoneyInSddInitOutput) Get() *MoneyInSddInitOutput {
	return v.value
}

func (v *NullableMoneyInSddInitOutput) Set(val *MoneyInSddInitOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInSddInitOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInSddInitOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInSddInitOutput(val *MoneyInSddInitOutput) *NullableMoneyInSddInitOutput {
	return &NullableMoneyInSddInitOutput{value: val, isSet: true}
}

func (v NullableMoneyInSddInitOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInSddInitOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

