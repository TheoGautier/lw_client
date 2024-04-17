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

// GetMoneyInIBANDetailsOutput struct for GetMoneyInIBANDetailsOutput
type GetMoneyInIBANDetailsOutput struct {
	Transactions *TransactionsTransactionIn `json:"transactions,omitempty"`
	Links *Links `json:"_links,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewGetMoneyInIBANDetailsOutput instantiates a new GetMoneyInIBANDetailsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMoneyInIBANDetailsOutput() *GetMoneyInIBANDetailsOutput {
	this := GetMoneyInIBANDetailsOutput{}
	return &this
}

// NewGetMoneyInIBANDetailsOutputWithDefaults instantiates a new GetMoneyInIBANDetailsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMoneyInIBANDetailsOutputWithDefaults() *GetMoneyInIBANDetailsOutput {
	this := GetMoneyInIBANDetailsOutput{}
	return &this
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *GetMoneyInIBANDetailsOutput) GetTransactions() TransactionsTransactionIn {
	if o == nil || o.Transactions == nil {
		var ret TransactionsTransactionIn
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyInIBANDetailsOutput) GetTransactionsOk() (*TransactionsTransactionIn, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *GetMoneyInIBANDetailsOutput) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given TransactionsTransactionIn and assigns it to the Transactions field.
func (o *GetMoneyInIBANDetailsOutput) SetTransactions(v TransactionsTransactionIn) {
	o.Transactions = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetMoneyInIBANDetailsOutput) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyInIBANDetailsOutput) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetMoneyInIBANDetailsOutput) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *GetMoneyInIBANDetailsOutput) SetLinks(v Links) {
	o.Links = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetMoneyInIBANDetailsOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyInIBANDetailsOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetMoneyInIBANDetailsOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *GetMoneyInIBANDetailsOutput) SetError(v Error) {
	o.Error = &v
}

func (o GetMoneyInIBANDetailsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableGetMoneyInIBANDetailsOutput struct {
	value *GetMoneyInIBANDetailsOutput
	isSet bool
}

func (v NullableGetMoneyInIBANDetailsOutput) Get() *GetMoneyInIBANDetailsOutput {
	return v.value
}

func (v *NullableGetMoneyInIBANDetailsOutput) Set(val *GetMoneyInIBANDetailsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMoneyInIBANDetailsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMoneyInIBANDetailsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMoneyInIBANDetailsOutput(val *GetMoneyInIBANDetailsOutput) *NullableGetMoneyInIBANDetailsOutput {
	return &NullableGetMoneyInIBANDetailsOutput{value: val, isSet: true}
}

func (v NullableGetMoneyInIBANDetailsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMoneyInIBANDetailsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


