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

// RefundMoneyInOutput struct for RefundMoneyInOutput
type RefundMoneyInOutput struct {
	Transaction *TransactionRefund `json:"transaction,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewRefundMoneyInOutput instantiates a new RefundMoneyInOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundMoneyInOutput() *RefundMoneyInOutput {
	this := RefundMoneyInOutput{}
	return &this
}

// NewRefundMoneyInOutputWithDefaults instantiates a new RefundMoneyInOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundMoneyInOutputWithDefaults() *RefundMoneyInOutput {
	this := RefundMoneyInOutput{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *RefundMoneyInOutput) GetTransaction() TransactionRefund {
	if o == nil || o.Transaction == nil {
		var ret TransactionRefund
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundMoneyInOutput) GetTransactionOk() (*TransactionRefund, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *RefundMoneyInOutput) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given TransactionRefund and assigns it to the Transaction field.
func (o *RefundMoneyInOutput) SetTransaction(v TransactionRefund) {
	o.Transaction = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RefundMoneyInOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundMoneyInOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RefundMoneyInOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *RefundMoneyInOutput) SetError(v Error) {
	o.Error = &v
}

func (o RefundMoneyInOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableRefundMoneyInOutput struct {
	value *RefundMoneyInOutput
	isSet bool
}

func (v NullableRefundMoneyInOutput) Get() *RefundMoneyInOutput {
	return v.value
}

func (v *NullableRefundMoneyInOutput) Set(val *RefundMoneyInOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundMoneyInOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundMoneyInOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundMoneyInOutput(val *RefundMoneyInOutput) *NullableRefundMoneyInOutput {
	return &NullableRefundMoneyInOutput{value: val, isSet: true}
}

func (v NullableRefundMoneyInOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundMoneyInOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

