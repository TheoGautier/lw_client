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

// CardInfo struct for CardInfo
type CardInfo struct {
	// Card Type<br/>0 = CB.<br/>1 = Visa.<br/>2 = Mastercard.<br/>
	CardType int32 `json:"cardType"`
	// Card Number
	CardNumber string `json:"cardNumber"`
	// CVV Code
	CardCode string `json:"cardCode"`
	// Card Expiration Date
	CardDate string `json:"cardDate"`
}

// NewCardInfo instantiates a new CardInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfo(cardType int32, cardNumber string, cardCode string, cardDate string) *CardInfo {
	this := CardInfo{}
	this.CardType = cardType
	this.CardNumber = cardNumber
	this.CardCode = cardCode
	this.CardDate = cardDate
	return &this
}

// NewCardInfoWithDefaults instantiates a new CardInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoWithDefaults() *CardInfo {
	this := CardInfo{}
	return &this
}

// GetCardType returns the CardType field value
func (o *CardInfo) GetCardType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetCardTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardType, true
}

// SetCardType sets field value
func (o *CardInfo) SetCardType(v int32) {
	o.CardType = v
}

// GetCardNumber returns the CardNumber field value
func (o *CardInfo) GetCardNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardNumber, true
}

// SetCardNumber sets field value
func (o *CardInfo) SetCardNumber(v string) {
	o.CardNumber = v
}

// GetCardCode returns the CardCode field value
func (o *CardInfo) GetCardCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardCode
}

// GetCardCodeOk returns a tuple with the CardCode field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetCardCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardCode, true
}

// SetCardCode sets field value
func (o *CardInfo) SetCardCode(v string) {
	o.CardCode = v
}

// GetCardDate returns the CardDate field value
func (o *CardInfo) GetCardDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardDate
}

// GetCardDateOk returns a tuple with the CardDate field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetCardDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardDate, true
}

// SetCardDate sets field value
func (o *CardInfo) SetCardDate(v string) {
	o.CardDate = v
}

func (o CardInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cardType"] = o.CardType
	}
	if true {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if true {
		toSerialize["cardCode"] = o.CardCode
	}
	if true {
		toSerialize["cardDate"] = o.CardDate
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfo struct {
	value *CardInfo
	isSet bool
}

func (v NullableCardInfo) Get() *CardInfo {
	return v.value
}

func (v *NullableCardInfo) Set(val *CardInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfo(val *CardInfo) *NullableCardInfo {
	return &NullableCardInfo{value: val, isSet: true}
}

func (v NullableCardInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


