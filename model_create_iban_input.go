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

// CreateIBANInput struct for CreateIBANInput
type CreateIBANInput struct {
	// Payment Account ID.
	Wallet string `json:"wallet"`
	// ISO2 code of the country from which the IBAN must be generated  We currently serve 5 countries:  <ul><li>BE: Belgium (BNP Fortis)</li><li>DE: Germany (BNP Paribas Niederlassung)</li><li>ES: BNP Paribas Sucursal en España</li><li>FR: France (BNP Paribas)</li><li>IT: Italy (BNL)</li></ul>
	Country string `json:"country"`
	// If activated, this function will return a PDF document describing the virtual IBAN, along with a <a href=\"https://www.europeanpaymentscouncil.eu/document-library/guidance-documents/quick-response-code-guidelines-enable-data-capture-initiation\" target=\"_blank\">standard EPC QR-Code</a> image. Please store these documents in your datawarehouse, as they are not accessible through Lemonway's API after the virtual IBAN has been generated.
	GeneratePDFAndQrCode *bool `json:"generatePDFAndQrCode,omitempty"`
	// ISO 639-1 language code for the PDF document. The supported languages are:  <ul><li>en: English</li><li>es: Spanish</li><li>fr: French</li><li>de: German</li><li>it: Italian</li></ul>  This field is ignored if generatePDFAndQrCode is false.
	PdfLanguage *string `json:"pdfLanguage,omitempty"`
}

// NewCreateIBANInput instantiates a new CreateIBANInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIBANInput(wallet string, country string) *CreateIBANInput {
	this := CreateIBANInput{}
	this.Wallet = wallet
	this.Country = country
	return &this
}

// NewCreateIBANInputWithDefaults instantiates a new CreateIBANInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIBANInputWithDefaults() *CreateIBANInput {
	this := CreateIBANInput{}
	return &this
}

// GetWallet returns the Wallet field value
func (o *CreateIBANInput) GetWallet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *CreateIBANInput) GetWalletOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *CreateIBANInput) SetWallet(v string) {
	o.Wallet = v
}

// GetCountry returns the Country field value
func (o *CreateIBANInput) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *CreateIBANInput) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *CreateIBANInput) SetCountry(v string) {
	o.Country = v
}

// GetGeneratePDFAndQrCode returns the GeneratePDFAndQrCode field value if set, zero value otherwise.
func (o *CreateIBANInput) GetGeneratePDFAndQrCode() bool {
	if o == nil || o.GeneratePDFAndQrCode == nil {
		var ret bool
		return ret
	}
	return *o.GeneratePDFAndQrCode
}

// GetGeneratePDFAndQrCodeOk returns a tuple with the GeneratePDFAndQrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIBANInput) GetGeneratePDFAndQrCodeOk() (*bool, bool) {
	if o == nil || o.GeneratePDFAndQrCode == nil {
		return nil, false
	}
	return o.GeneratePDFAndQrCode, true
}

// HasGeneratePDFAndQrCode returns a boolean if a field has been set.
func (o *CreateIBANInput) HasGeneratePDFAndQrCode() bool {
	if o != nil && o.GeneratePDFAndQrCode != nil {
		return true
	}

	return false
}

// SetGeneratePDFAndQrCode gets a reference to the given bool and assigns it to the GeneratePDFAndQrCode field.
func (o *CreateIBANInput) SetGeneratePDFAndQrCode(v bool) {
	o.GeneratePDFAndQrCode = &v
}

// GetPdfLanguage returns the PdfLanguage field value if set, zero value otherwise.
func (o *CreateIBANInput) GetPdfLanguage() string {
	if o == nil || o.PdfLanguage == nil {
		var ret string
		return ret
	}
	return *o.PdfLanguage
}

// GetPdfLanguageOk returns a tuple with the PdfLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIBANInput) GetPdfLanguageOk() (*string, bool) {
	if o == nil || o.PdfLanguage == nil {
		return nil, false
	}
	return o.PdfLanguage, true
}

// HasPdfLanguage returns a boolean if a field has been set.
func (o *CreateIBANInput) HasPdfLanguage() bool {
	if o != nil && o.PdfLanguage != nil {
		return true
	}

	return false
}

// SetPdfLanguage gets a reference to the given string and assigns it to the PdfLanguage field.
func (o *CreateIBANInput) SetPdfLanguage(v string) {
	o.PdfLanguage = &v
}

func (o CreateIBANInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallet"] = o.Wallet
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if o.GeneratePDFAndQrCode != nil {
		toSerialize["generatePDFAndQrCode"] = o.GeneratePDFAndQrCode
	}
	if o.PdfLanguage != nil {
		toSerialize["pdfLanguage"] = o.PdfLanguage
	}
	return json.Marshal(toSerialize)
}

type NullableCreateIBANInput struct {
	value *CreateIBANInput
	isSet bool
}

func (v NullableCreateIBANInput) Get() *CreateIBANInput {
	return v.value
}

func (v *NullableCreateIBANInput) Set(val *CreateIBANInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIBANInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIBANInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIBANInput(val *CreateIBANInput) *NullableCreateIBANInput {
	return &NullableCreateIBANInput{value: val, isSet: true}
}

func (v NullableCreateIBANInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIBANInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


