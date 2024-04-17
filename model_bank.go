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

// Bank struct for Bank
type Bank struct {
	// Country code associated with the payment bank. ISO Alpha-2 codes accepted: (FR) France, (ES), Spain (DE), (IT) Italy, Germany, and (PT) Portugal
	CountryCode *string `json:"countryCode,omitempty"`
	// The bank's parent company name. Example: Banque Populaire
	ParentName *string `json:"parentName,omitempty"`
	// A weblink to the bank's logo
	LogoURL *string `json:"logoURL,omitempty"`
	// This is the bank name. Example: Banque Populaire Alsace Lorraine
	BankName *string `json:"bankName,omitempty"`
	// This is the unique `bankId` associated with the bank.
	BankId *string `json:"bankId,omitempty"`
	// Personal or Business Acoount
	Segments []string `json:"segments,omitempty"`
}

// NewBank instantiates a new Bank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBank() *Bank {
	this := Bank{}
	return &this
}

// NewBankWithDefaults instantiates a new Bank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankWithDefaults() *Bank {
	this := Bank{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Bank) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Bank) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Bank) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *Bank) GetParentName() string {
	if o == nil || o.ParentName == nil {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetParentNameOk() (*string, bool) {
	if o == nil || o.ParentName == nil {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *Bank) HasParentName() bool {
	if o != nil && o.ParentName != nil {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *Bank) SetParentName(v string) {
	o.ParentName = &v
}

// GetLogoURL returns the LogoURL field value if set, zero value otherwise.
func (o *Bank) GetLogoURL() string {
	if o == nil || o.LogoURL == nil {
		var ret string
		return ret
	}
	return *o.LogoURL
}

// GetLogoURLOk returns a tuple with the LogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetLogoURLOk() (*string, bool) {
	if o == nil || o.LogoURL == nil {
		return nil, false
	}
	return o.LogoURL, true
}

// HasLogoURL returns a boolean if a field has been set.
func (o *Bank) HasLogoURL() bool {
	if o != nil && o.LogoURL != nil {
		return true
	}

	return false
}

// SetLogoURL gets a reference to the given string and assigns it to the LogoURL field.
func (o *Bank) SetLogoURL(v string) {
	o.LogoURL = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *Bank) GetBankName() string {
	if o == nil || o.BankName == nil {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetBankNameOk() (*string, bool) {
	if o == nil || o.BankName == nil {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *Bank) HasBankName() bool {
	if o != nil && o.BankName != nil {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *Bank) SetBankName(v string) {
	o.BankName = &v
}

// GetBankId returns the BankId field value if set, zero value otherwise.
func (o *Bank) GetBankId() string {
	if o == nil || o.BankId == nil {
		var ret string
		return ret
	}
	return *o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetBankIdOk() (*string, bool) {
	if o == nil || o.BankId == nil {
		return nil, false
	}
	return o.BankId, true
}

// HasBankId returns a boolean if a field has been set.
func (o *Bank) HasBankId() bool {
	if o != nil && o.BankId != nil {
		return true
	}

	return false
}

// SetBankId gets a reference to the given string and assigns it to the BankId field.
func (o *Bank) SetBankId(v string) {
	o.BankId = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *Bank) GetSegments() []string {
	if o == nil || o.Segments == nil {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetSegmentsOk() ([]string, bool) {
	if o == nil || o.Segments == nil {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *Bank) HasSegments() bool {
	if o != nil && o.Segments != nil {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *Bank) SetSegments(v []string) {
	o.Segments = v
}

func (o Bank) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.ParentName != nil {
		toSerialize["parentName"] = o.ParentName
	}
	if o.LogoURL != nil {
		toSerialize["logoURL"] = o.LogoURL
	}
	if o.BankName != nil {
		toSerialize["bankName"] = o.BankName
	}
	if o.BankId != nil {
		toSerialize["bankId"] = o.BankId
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	return json.Marshal(toSerialize)
}

type NullableBank struct {
	value *Bank
	isSet bool
}

func (v NullableBank) Get() *Bank {
	return v.value
}

func (v *NullableBank) Set(val *Bank) {
	v.value = val
	v.isSet = true
}

func (v NullableBank) IsSet() bool {
	return v.isSet
}

func (v *NullableBank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBank(val *Bank) *NullableBank {
	return &NullableBank{value: val, isSet: true}
}

func (v NullableBank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


