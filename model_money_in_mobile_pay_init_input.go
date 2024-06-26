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

// MoneyInMobilePayInitInput struct for MoneyInMobilePayInitInput
type MoneyInMobilePayInitInput struct {
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

// NewMoneyInMobilePayInitInput instantiates a new MoneyInMobilePayInitInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInMobilePayInitInput(returnUrl string, accountId string) *MoneyInMobilePayInitInput {
	this := MoneyInMobilePayInitInput{}
	this.ReturnUrl = returnUrl
	this.AccountId = accountId
	return &this
}

// NewMoneyInMobilePayInitInputWithDefaults instantiates a new MoneyInMobilePayInitInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInMobilePayInitInputWithDefaults() *MoneyInMobilePayInitInput {
	this := MoneyInMobilePayInitInput{}
	return &this
}

// GetReturnUrl returns the ReturnUrl field value
func (o *MoneyInMobilePayInitInput) GetReturnUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value
// and a boolean to check if the value has been set.
func (o *MoneyInMobilePayInitInput) GetReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// SetReturnUrl sets field value
func (o *MoneyInMobilePayInitInput) SetReturnUrl(v string) {
	o.ReturnUrl = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MoneyInMobilePayInitInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInMobilePayInitInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MoneyInMobilePayInitInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MoneyInMobilePayInitInput) SetReference(v string) {
	o.Reference = &v
}

// GetAccountId returns the AccountId field value
func (o *MoneyInMobilePayInitInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MoneyInMobilePayInitInput) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MoneyInMobilePayInitInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MoneyInMobilePayInitInput) GetTotalAmount() int32 {
	if o == nil || o.TotalAmount == nil {
		var ret int32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInMobilePayInitInput) GetTotalAmountOk() (*int32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MoneyInMobilePayInitInput) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int32 and assigns it to the TotalAmount field.
func (o *MoneyInMobilePayInitInput) SetTotalAmount(v int32) {
	o.TotalAmount = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *MoneyInMobilePayInitInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInMobilePayInitInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *MoneyInMobilePayInitInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *MoneyInMobilePayInitInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MoneyInMobilePayInitInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInMobilePayInitInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MoneyInMobilePayInitInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MoneyInMobilePayInitInput) SetComment(v string) {
	o.Comment = &v
}

// GetAutoCommission returns the AutoCommission field value if set, zero value otherwise.
func (o *MoneyInMobilePayInitInput) GetAutoCommission() bool {
	if o == nil || o.AutoCommission == nil {
		var ret bool
		return ret
	}
	return *o.AutoCommission
}

// GetAutoCommissionOk returns a tuple with the AutoCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInMobilePayInitInput) GetAutoCommissionOk() (*bool, bool) {
	if o == nil || o.AutoCommission == nil {
		return nil, false
	}
	return o.AutoCommission, true
}

// HasAutoCommission returns a boolean if a field has been set.
func (o *MoneyInMobilePayInitInput) HasAutoCommission() bool {
	if o != nil && o.AutoCommission != nil {
		return true
	}

	return false
}

// SetAutoCommission gets a reference to the given bool and assigns it to the AutoCommission field.
func (o *MoneyInMobilePayInitInput) SetAutoCommission(v bool) {
	o.AutoCommission = &v
}

func (o MoneyInMobilePayInitInput) MarshalJSON() ([]byte, error) {
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

type NullableMoneyInMobilePayInitInput struct {
	value *MoneyInMobilePayInitInput
	isSet bool
}

func (v NullableMoneyInMobilePayInitInput) Get() *MoneyInMobilePayInitInput {
	return v.value
}

func (v *NullableMoneyInMobilePayInitInput) Set(val *MoneyInMobilePayInitInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInMobilePayInitInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInMobilePayInitInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInMobilePayInitInput(val *MoneyInMobilePayInitInput) *NullableMoneyInMobilePayInitInput {
	return &NullableMoneyInMobilePayInitInput{value: val, isSet: true}
}

func (v NullableMoneyInMobilePayInitInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInMobilePayInitInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


