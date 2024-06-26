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

// Document struct for Document
type Document struct {
	// Document ID
	Id *int64 `json:"id,omitempty"`
	// Document Status<br/>0 = Document put on hold, waiting for another document.<br/>1 = Received, need manual validation.<br/>2 = Accepted.<br/>3 = Rejected.<br/>4 = Rejected. Unreadable by human (Cropped, blur, glare…).<br/>5 = Rejected. Expired (Expiration Date is passed).<br/>6 = Rejected. Wrong Type (Document not accepted).<br/>7 = Rejected. Wrong Name (Name not matching user information).<br/>8 = Rejected. Duplicated Document.<br/>
	Status *int32 `json:"status,omitempty"`
	// Document Type<br/>0 = ID card (both sides in one file).<br/>1 = Proof of address.<br/>2 = Scan of a proof of IBAN.<br/>3 = Passport (European Union).<br/>4 = Passport (outside the European Union).<br/>5 = Residence permit (both sides in one file).<br/>6 = Other document type.<br/>7 = Official company registration document (Kbis extract or equivalent).<br/>11 = Driver licence (both sides in one file).<br/>12 = Status.<br/>13 = Selfie.<br/>14 = Other document type.<br/>15 = Other document type.<br/>16 = Other document type.<br/>17 = Other document type.<br/>18 = Other document type.<br/>19 = Other document type.<br/>20 = Other document type.<br/>21 = SDD mandate.<br/>
	Type *int32 `json:"type,omitempty"`
	// Document validity date  Format: dd/mm/yyyy  Empty if unknown
	ValidityDate *string `json:"validityDate,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewDocument instantiates a new Document object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocument() *Document {
	this := Document{}
	return &this
}

// NewDocumentWithDefaults instantiates a new Document object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentWithDefaults() *Document {
	this := Document{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Document) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Document) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Document) SetId(v int64) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Document) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Document) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *Document) SetStatus(v int32) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Document) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Document) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Document) SetType(v int32) {
	o.Type = &v
}

// GetValidityDate returns the ValidityDate field value if set, zero value otherwise.
func (o *Document) GetValidityDate() string {
	if o == nil || o.ValidityDate == nil {
		var ret string
		return ret
	}
	return *o.ValidityDate
}

// GetValidityDateOk returns a tuple with the ValidityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetValidityDateOk() (*string, bool) {
	if o == nil || o.ValidityDate == nil {
		return nil, false
	}
	return o.ValidityDate, true
}

// HasValidityDate returns a boolean if a field has been set.
func (o *Document) HasValidityDate() bool {
	if o != nil && o.ValidityDate != nil {
		return true
	}

	return false
}

// SetValidityDate gets a reference to the given string and assigns it to the ValidityDate field.
func (o *Document) SetValidityDate(v string) {
	o.ValidityDate = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Document) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Document) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Document) SetComment(v string) {
	o.Comment = &v
}

func (o Document) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ValidityDate != nil {
		toSerialize["validityDate"] = o.ValidityDate
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableDocument struct {
	value *Document
	isSet bool
}

func (v NullableDocument) Get() *Document {
	return v.value
}

func (v *NullableDocument) Set(val *Document) {
	v.value = val
	v.isSet = true
}

func (v NullableDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocument(val *Document) *NullableDocument {
	return &NullableDocument{value: val, isSet: true}
}

func (v NullableDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


