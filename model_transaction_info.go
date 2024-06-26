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

// TransactionInfo struct for TransactionInfo
type TransactionInfo struct {
	// Transaction Reference Sent to the PSP
	Reference string `json:"reference"`
	// Merchant Order ID
	Order string `json:"order"`
	// Transaction DateTime (UTC Unix timestamp)
	DateTime string `json:"dateTime"`
	// Merchant ID Sent to PSP
	MerchantId string `json:"merchantId"`
	// Transaction Authorisation ID Received from PSP
	AuthorisationId string `json:"authorisationId"`
}

// NewTransactionInfo instantiates a new TransactionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInfo(reference string, order string, dateTime string, merchantId string, authorisationId string) *TransactionInfo {
	this := TransactionInfo{}
	this.Reference = reference
	this.Order = order
	this.DateTime = dateTime
	this.MerchantId = merchantId
	this.AuthorisationId = authorisationId
	return &this
}

// NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInfoWithDefaults() *TransactionInfo {
	this := TransactionInfo{}
	return &this
}

// GetReference returns the Reference field value
func (o *TransactionInfo) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *TransactionInfo) SetReference(v string) {
	o.Reference = v
}

// GetOrder returns the Order field value
func (o *TransactionInfo) GetOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *TransactionInfo) SetOrder(v string) {
	o.Order = v
}

// GetDateTime returns the DateTime field value
func (o *TransactionInfo) GetDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateTime, true
}

// SetDateTime sets field value
func (o *TransactionInfo) SetDateTime(v string) {
	o.DateTime = v
}

// GetMerchantId returns the MerchantId field value
func (o *TransactionInfo) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *TransactionInfo) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetAuthorisationId returns the AuthorisationId field value
func (o *TransactionInfo) GetAuthorisationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorisationId
}

// GetAuthorisationIdOk returns a tuple with the AuthorisationId field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetAuthorisationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorisationId, true
}

// SetAuthorisationId sets field value
func (o *TransactionInfo) SetAuthorisationId(v string) {
	o.AuthorisationId = v
}

func (o TransactionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["order"] = o.Order
	}
	if true {
		toSerialize["dateTime"] = o.DateTime
	}
	if true {
		toSerialize["merchantId"] = o.MerchantId
	}
	if true {
		toSerialize["authorisationId"] = o.AuthorisationId
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionInfo struct {
	value *TransactionInfo
	isSet bool
}

func (v NullableTransactionInfo) Get() *TransactionInfo {
	return v.value
}

func (v *NullableTransactionInfo) Set(val *TransactionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInfo(val *TransactionInfo) *NullableTransactionInfo {
	return &NullableTransactionInfo{value: val, isSet: true}
}

func (v NullableTransactionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


