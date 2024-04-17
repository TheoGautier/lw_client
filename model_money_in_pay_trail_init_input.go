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

// MoneyInPayTrailInitInput struct for MoneyInPayTrailInitInput
type MoneyInPayTrailInitInput struct {
	// Your site returns the URL, called to terminate the operation and on which the callback will be sent, with data in POST parameters.  This URL must contain a unique ID so you know which operation is related to the return.
	ReturnUrl string `json:"returnUrl"`
	// Unique ID of the call, generated by your server. This ID can be used as a search field when looking for operation details
	Reference *string `json:"reference,omitempty"`
	// Payment Account ID to Credit
	AccountId string `json:"accountId"`
	// Amount to Debit  Amounts are given as integer numbers in cents
	TotalAmount *int32 `json:"totalAmount,omitempty"`
	// Your Fee  Amounts are given as integer numbers in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Comment Regarding the Transaction
	Comment *string `json:"comment,omitempty"`
	// If true:  1. [amountCom] will be ignored and will be replaced with Lemonway's fee  2. You will not receive any fee
	AutoCommission *bool `json:"autoCommission,omitempty"`
}

// NewMoneyInPayTrailInitInput instantiates a new MoneyInPayTrailInitInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInPayTrailInitInput(returnUrl string, accountId string) *MoneyInPayTrailInitInput {
	this := MoneyInPayTrailInitInput{}
	this.ReturnUrl = returnUrl
	this.AccountId = accountId
	return &this
}

// NewMoneyInPayTrailInitInputWithDefaults instantiates a new MoneyInPayTrailInitInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInPayTrailInitInputWithDefaults() *MoneyInPayTrailInitInput {
	this := MoneyInPayTrailInitInput{}
	return &this
}

// GetReturnUrl returns the ReturnUrl field value
func (o *MoneyInPayTrailInitInput) GetReturnUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitInput) GetReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// SetReturnUrl sets field value
func (o *MoneyInPayTrailInitInput) SetReturnUrl(v string) {
	o.ReturnUrl = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MoneyInPayTrailInitInput) SetReference(v string) {
	o.Reference = &v
}

// GetAccountId returns the AccountId field value
func (o *MoneyInPayTrailInitInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitInput) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MoneyInPayTrailInitInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitInput) GetTotalAmount() int32 {
	if o == nil || o.TotalAmount == nil {
		var ret int32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitInput) GetTotalAmountOk() (*int32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitInput) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int32 and assigns it to the TotalAmount field.
func (o *MoneyInPayTrailInitInput) SetTotalAmount(v int32) {
	o.TotalAmount = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *MoneyInPayTrailInitInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MoneyInPayTrailInitInput) SetComment(v string) {
	o.Comment = &v
}

// GetAutoCommission returns the AutoCommission field value if set, zero value otherwise.
func (o *MoneyInPayTrailInitInput) GetAutoCommission() bool {
	if o == nil || o.AutoCommission == nil {
		var ret bool
		return ret
	}
	return *o.AutoCommission
}

// GetAutoCommissionOk returns a tuple with the AutoCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInPayTrailInitInput) GetAutoCommissionOk() (*bool, bool) {
	if o == nil || o.AutoCommission == nil {
		return nil, false
	}
	return o.AutoCommission, true
}

// HasAutoCommission returns a boolean if a field has been set.
func (o *MoneyInPayTrailInitInput) HasAutoCommission() bool {
	if o != nil && o.AutoCommission != nil {
		return true
	}

	return false
}

// SetAutoCommission gets a reference to the given bool and assigns it to the AutoCommission field.
func (o *MoneyInPayTrailInitInput) SetAutoCommission(v bool) {
	o.AutoCommission = &v
}

func (o MoneyInPayTrailInitInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if o.TotalAmount != nil {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if o.CommissionAmount != nil {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.AutoCommission != nil {
		toSerialize["autoCommission"] = o.AutoCommission
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyInPayTrailInitInput struct {
	value *MoneyInPayTrailInitInput
	isSet bool
}

func (v NullableMoneyInPayTrailInitInput) Get() *MoneyInPayTrailInitInput {
	return v.value
}

func (v *NullableMoneyInPayTrailInitInput) Set(val *MoneyInPayTrailInitInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInPayTrailInitInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInPayTrailInitInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInPayTrailInitInput(val *MoneyInPayTrailInitInput) *NullableMoneyInPayTrailInitInput {
	return &NullableMoneyInPayTrailInitInput{value: val, isSet: true}
}

func (v NullableMoneyInPayTrailInitInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInPayTrailInitInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


