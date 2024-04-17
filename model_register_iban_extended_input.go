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

// RegisterIBANExtendedInput struct for RegisterIBANExtendedInput
type RegisterIBANExtendedInput struct {
	// Payment Account ID
	Wallet string `json:"wallet"`
	// Account Type, by default Other<br/>0 = Other.<br/>1 = IBAN.<br/>2 = BBAN/RIB.<br/>
	AccountType *int32 `json:"accountType,omitempty"`
	// The registered bank account owner: First and Last name, or Company Name
	HolderName string `json:"holderName"`
	// Account Number. The format depends on the account type.
	AccountNumber string `json:"accountNumber"`
	// Country of the beneficiary. Two-letter country code (ISO alpha-2) for example, France=<b>FR</b>
	HolderCountry *string `json:"holderCountry,omitempty"`
	// BIC/SWIFT codes are arranged like this: AAAABBCCDDD  AAAA: 4 character for bank code  BB: 2 char for country code  CC: 2 char for location code  DDD: 3 char for branch code
	BicCode *string `json:"bicCode,omitempty"`
	// Bank Name. This field is mandatory in the following circumstances:  <ul><li>If the currency is USD</li><li>If the selected bank country code requires this field</li></ul>
	BankName *string `json:"bankName,omitempty"`
	// Country of the Bank. Two-letter country code (ISO alpha-2) for example, France=<b>FR</b>
	BankCountry string `json:"bankCountry"`
	// Bank Branch Code (Sort Code in the United Kingdom). It is mandatory for the examples below:  <ul><li>If the selected bank country code requires this field.</li></ul>
	BankBranchCode *string `json:"bankBranchCode,omitempty"`
	BankBranchAddress *BankBranchAddress `json:"bankBranchAddress,omitempty"`
	// BIC/SWIFT Code of the Intermediary Bank.
	IntermediaryBicCode *string `json:"intermediaryBicCode,omitempty"`
	// Intermediary Bank Name
	IntermediaryBankName *string `json:"intermediaryBankName,omitempty"`
	// Bank Country Code of the Intermediary Bank. Two-letter country code (ISO alpha-2) for example, France=<b>FR</b>
	IntermediaryBankCountry *string `json:"intermediaryBankCountry,omitempty"`
	// Reason for New Bank Account Details if another one is already linked to the Payment Account.
	Comment *string `json:"comment,omitempty"`
}

// NewRegisterIBANExtendedInput instantiates a new RegisterIBANExtendedInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterIBANExtendedInput(wallet string, holderName string, accountNumber string, bankCountry string) *RegisterIBANExtendedInput {
	this := RegisterIBANExtendedInput{}
	this.Wallet = wallet
	this.HolderName = holderName
	this.AccountNumber = accountNumber
	this.BankCountry = bankCountry
	return &this
}

// NewRegisterIBANExtendedInputWithDefaults instantiates a new RegisterIBANExtendedInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterIBANExtendedInputWithDefaults() *RegisterIBANExtendedInput {
	this := RegisterIBANExtendedInput{}
	return &this
}

// GetWallet returns the Wallet field value
func (o *RegisterIBANExtendedInput) GetWallet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetWalletOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *RegisterIBANExtendedInput) SetWallet(v string) {
	o.Wallet = v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetAccountType() int32 {
	if o == nil || o.AccountType == nil {
		var ret int32
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetAccountTypeOk() (*int32, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given int32 and assigns it to the AccountType field.
func (o *RegisterIBANExtendedInput) SetAccountType(v int32) {
	o.AccountType = &v
}

// GetHolderName returns the HolderName field value
func (o *RegisterIBANExtendedInput) GetHolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolderName, true
}

// SetHolderName sets field value
func (o *RegisterIBANExtendedInput) SetHolderName(v string) {
	o.HolderName = v
}

// GetAccountNumber returns the AccountNumber field value
func (o *RegisterIBANExtendedInput) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *RegisterIBANExtendedInput) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetHolderCountry returns the HolderCountry field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetHolderCountry() string {
	if o == nil || o.HolderCountry == nil {
		var ret string
		return ret
	}
	return *o.HolderCountry
}

// GetHolderCountryOk returns a tuple with the HolderCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetHolderCountryOk() (*string, bool) {
	if o == nil || o.HolderCountry == nil {
		return nil, false
	}
	return o.HolderCountry, true
}

// HasHolderCountry returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasHolderCountry() bool {
	if o != nil && o.HolderCountry != nil {
		return true
	}

	return false
}

// SetHolderCountry gets a reference to the given string and assigns it to the HolderCountry field.
func (o *RegisterIBANExtendedInput) SetHolderCountry(v string) {
	o.HolderCountry = &v
}

// GetBicCode returns the BicCode field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetBicCode() string {
	if o == nil || o.BicCode == nil {
		var ret string
		return ret
	}
	return *o.BicCode
}

// GetBicCodeOk returns a tuple with the BicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetBicCodeOk() (*string, bool) {
	if o == nil || o.BicCode == nil {
		return nil, false
	}
	return o.BicCode, true
}

// HasBicCode returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasBicCode() bool {
	if o != nil && o.BicCode != nil {
		return true
	}

	return false
}

// SetBicCode gets a reference to the given string and assigns it to the BicCode field.
func (o *RegisterIBANExtendedInput) SetBicCode(v string) {
	o.BicCode = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetBankName() string {
	if o == nil || o.BankName == nil {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetBankNameOk() (*string, bool) {
	if o == nil || o.BankName == nil {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasBankName() bool {
	if o != nil && o.BankName != nil {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *RegisterIBANExtendedInput) SetBankName(v string) {
	o.BankName = &v
}

// GetBankCountry returns the BankCountry field value
func (o *RegisterIBANExtendedInput) GetBankCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCountry
}

// GetBankCountryOk returns a tuple with the BankCountry field value
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetBankCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCountry, true
}

// SetBankCountry sets field value
func (o *RegisterIBANExtendedInput) SetBankCountry(v string) {
	o.BankCountry = v
}

// GetBankBranchCode returns the BankBranchCode field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetBankBranchCode() string {
	if o == nil || o.BankBranchCode == nil {
		var ret string
		return ret
	}
	return *o.BankBranchCode
}

// GetBankBranchCodeOk returns a tuple with the BankBranchCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetBankBranchCodeOk() (*string, bool) {
	if o == nil || o.BankBranchCode == nil {
		return nil, false
	}
	return o.BankBranchCode, true
}

// HasBankBranchCode returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasBankBranchCode() bool {
	if o != nil && o.BankBranchCode != nil {
		return true
	}

	return false
}

// SetBankBranchCode gets a reference to the given string and assigns it to the BankBranchCode field.
func (o *RegisterIBANExtendedInput) SetBankBranchCode(v string) {
	o.BankBranchCode = &v
}

// GetBankBranchAddress returns the BankBranchAddress field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetBankBranchAddress() BankBranchAddress {
	if o == nil || o.BankBranchAddress == nil {
		var ret BankBranchAddress
		return ret
	}
	return *o.BankBranchAddress
}

// GetBankBranchAddressOk returns a tuple with the BankBranchAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetBankBranchAddressOk() (*BankBranchAddress, bool) {
	if o == nil || o.BankBranchAddress == nil {
		return nil, false
	}
	return o.BankBranchAddress, true
}

// HasBankBranchAddress returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasBankBranchAddress() bool {
	if o != nil && o.BankBranchAddress != nil {
		return true
	}

	return false
}

// SetBankBranchAddress gets a reference to the given BankBranchAddress and assigns it to the BankBranchAddress field.
func (o *RegisterIBANExtendedInput) SetBankBranchAddress(v BankBranchAddress) {
	o.BankBranchAddress = &v
}

// GetIntermediaryBicCode returns the IntermediaryBicCode field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetIntermediaryBicCode() string {
	if o == nil || o.IntermediaryBicCode == nil {
		var ret string
		return ret
	}
	return *o.IntermediaryBicCode
}

// GetIntermediaryBicCodeOk returns a tuple with the IntermediaryBicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetIntermediaryBicCodeOk() (*string, bool) {
	if o == nil || o.IntermediaryBicCode == nil {
		return nil, false
	}
	return o.IntermediaryBicCode, true
}

// HasIntermediaryBicCode returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasIntermediaryBicCode() bool {
	if o != nil && o.IntermediaryBicCode != nil {
		return true
	}

	return false
}

// SetIntermediaryBicCode gets a reference to the given string and assigns it to the IntermediaryBicCode field.
func (o *RegisterIBANExtendedInput) SetIntermediaryBicCode(v string) {
	o.IntermediaryBicCode = &v
}

// GetIntermediaryBankName returns the IntermediaryBankName field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetIntermediaryBankName() string {
	if o == nil || o.IntermediaryBankName == nil {
		var ret string
		return ret
	}
	return *o.IntermediaryBankName
}

// GetIntermediaryBankNameOk returns a tuple with the IntermediaryBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetIntermediaryBankNameOk() (*string, bool) {
	if o == nil || o.IntermediaryBankName == nil {
		return nil, false
	}
	return o.IntermediaryBankName, true
}

// HasIntermediaryBankName returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasIntermediaryBankName() bool {
	if o != nil && o.IntermediaryBankName != nil {
		return true
	}

	return false
}

// SetIntermediaryBankName gets a reference to the given string and assigns it to the IntermediaryBankName field.
func (o *RegisterIBANExtendedInput) SetIntermediaryBankName(v string) {
	o.IntermediaryBankName = &v
}

// GetIntermediaryBankCountry returns the IntermediaryBankCountry field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetIntermediaryBankCountry() string {
	if o == nil || o.IntermediaryBankCountry == nil {
		var ret string
		return ret
	}
	return *o.IntermediaryBankCountry
}

// GetIntermediaryBankCountryOk returns a tuple with the IntermediaryBankCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetIntermediaryBankCountryOk() (*string, bool) {
	if o == nil || o.IntermediaryBankCountry == nil {
		return nil, false
	}
	return o.IntermediaryBankCountry, true
}

// HasIntermediaryBankCountry returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasIntermediaryBankCountry() bool {
	if o != nil && o.IntermediaryBankCountry != nil {
		return true
	}

	return false
}

// SetIntermediaryBankCountry gets a reference to the given string and assigns it to the IntermediaryBankCountry field.
func (o *RegisterIBANExtendedInput) SetIntermediaryBankCountry(v string) {
	o.IntermediaryBankCountry = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RegisterIBANExtendedInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterIBANExtendedInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RegisterIBANExtendedInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RegisterIBANExtendedInput) SetComment(v string) {
	o.Comment = &v
}

func (o RegisterIBANExtendedInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallet"] = o.Wallet
	}
	if o.AccountType != nil {
		toSerialize["accountType"] = o.AccountType
	}
	if true {
		toSerialize["holderName"] = o.HolderName
	}
	if true {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if o.HolderCountry != nil {
		toSerialize["holderCountry"] = o.HolderCountry
	}
	if o.BicCode != nil {
		toSerialize["bicCode"] = o.BicCode
	}
	if o.BankName != nil {
		toSerialize["bankName"] = o.BankName
	}
	if true {
		toSerialize["bankCountry"] = o.BankCountry
	}
	if o.BankBranchCode != nil {
		toSerialize["bankBranchCode"] = o.BankBranchCode
	}
	if o.BankBranchAddress != nil {
		toSerialize["bankBranchAddress"] = o.BankBranchAddress
	}
	if o.IntermediaryBicCode != nil {
		toSerialize["intermediaryBicCode"] = o.IntermediaryBicCode
	}
	if o.IntermediaryBankName != nil {
		toSerialize["intermediaryBankName"] = o.IntermediaryBankName
	}
	if o.IntermediaryBankCountry != nil {
		toSerialize["intermediaryBankCountry"] = o.IntermediaryBankCountry
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterIBANExtendedInput struct {
	value *RegisterIBANExtendedInput
	isSet bool
}

func (v NullableRegisterIBANExtendedInput) Get() *RegisterIBANExtendedInput {
	return v.value
}

func (v *NullableRegisterIBANExtendedInput) Set(val *RegisterIBANExtendedInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterIBANExtendedInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterIBANExtendedInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterIBANExtendedInput(val *RegisterIBANExtendedInput) *NullableRegisterIBANExtendedInput {
	return &NullableRegisterIBANExtendedInput{value: val, isSet: true}
}

func (v NullableRegisterIBANExtendedInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterIBANExtendedInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


