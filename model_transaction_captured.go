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

// TransactionCaptured struct for TransactionCaptured
type TransactionCaptured struct {
	// Credited Account
	ReceiverAccountId *string `json:"receiverAccountId,omitempty"`
	// Credited Amount  Amounts represented in integer (cents)
	CreditAmount *int32 `json:"creditAmount,omitempty"`
	// 
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	// 
	ScheduledNumber *string `json:"scheduledNumber,omitempty"`
	MaskedLabel *string `json:"maskedLabel,omitempty"`
	PSP *PSP `json:"PSP,omitempty"`
	Card *Card `json:"card,omitempty"`
	BankStatus *string `json:"bankStatus,omitempty"`
	RefundAmount *float64 `json:"refundAmount,omitempty"`
	// Bank reference
	BankReference *string `json:"bankReference,omitempty"`
	// A specified postal address for the cheque
	ChequeSendingAddressCorporateName *string `json:"ChequeSendingAddress_CorporateName,omitempty"`
	// Cheque sending address street name
	ChequeSendingAddressStreet *string `json:"ChequeSendingAddress_Street,omitempty"`
	// Cheque sending address city name
	ChequeSendingAddressCity *string `json:"ChequeSendingAddress_City,omitempty"`
	// Cheque sending address post code
	ChequeSendingAddressPostCode *string `json:"ChequeSendingAddress_PostCode,omitempty"`
	ThreeDs *ThreeDs `json:"threeDS,omitempty"`
	BuyNowPayLaterInfo *BnplInfo `json:"buyNowPayLaterInfo,omitempty"`
	// Transaction ID
	Id *int64 `json:"id,omitempty"`
	// Payment Method    0 Card      1 Bank transfer (MoneyIn)       3 Bank transfer (MoneyOut)       4 P2P      13 iDEAL      14 SEPA DirectDebit      15 Cheque      17 Sofort      19 Multibanco      21 MBWAY      23 Pagare        30 BNPL    35 PayPal     ---   **Important:** The following services have been discontinued.      16 Neosurf    18 PFS Physical Card    20 Payshop    22 Polish Instant Transfer     24 MobilePay  25 Paytrail  26 WeChat    27 P24    28 MoneyIn by TPE    29 Trustly
	Method *int32 `json:"method,omitempty"`
	// Payment Method Details  0 Standard  1 Pay By Bank<br/>0 = STANDARD.<br/>1 = PAY_BY_BANK.<br/>
	MethodDetails *int32 `json:"methodDetails,omitempty"`
	// Transaction initialization date, UTC Unix timestamp
	Date *int32 `json:"date,omitempty"`
	// Your fee  Amounts are given as integer numbers in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Comment
	Comment *string `json:"comment,omitempty"`
	// **Money-In and Money-Out**         0: Success    3: Lemonway Error    4: Pending    6: PSP error    7: Cancelled    16: Validation pending      **Important:**  For /v2/moneyins/{transactionid}/validate` use the status codes:      0: Waiting for finalization      3: Success      4: Lemonway Error      6: PSP Error    7: Cancelled    16: Validation Pending                **P2P**    0: Pending    3: Success
	Status *int32 `json:"status,omitempty"`
	// Transaction execution date
	ExecutionDate *int32 `json:"executionDate,omitempty"`
	LemonWayCommission *LemonWayCommission `json:"lemonWayCommission,omitempty"`
	// Unique ID generated by your server
	Reference *string `json:"reference,omitempty"`
	RemainingAmount *int32 `json:"remainingAmount,omitempty"`
	Authorization *TransactionCapturedAuthorization `json:"authorization,omitempty"`
}

// NewTransactionCaptured instantiates a new TransactionCaptured object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCaptured() *TransactionCaptured {
	this := TransactionCaptured{}
	return &this
}

// NewTransactionCapturedWithDefaults instantiates a new TransactionCaptured object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCapturedWithDefaults() *TransactionCaptured {
	this := TransactionCaptured{}
	return &this
}

// GetReceiverAccountId returns the ReceiverAccountId field value if set, zero value otherwise.
func (o *TransactionCaptured) GetReceiverAccountId() string {
	if o == nil || o.ReceiverAccountId == nil {
		var ret string
		return ret
	}
	return *o.ReceiverAccountId
}

// GetReceiverAccountIdOk returns a tuple with the ReceiverAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetReceiverAccountIdOk() (*string, bool) {
	if o == nil || o.ReceiverAccountId == nil {
		return nil, false
	}
	return o.ReceiverAccountId, true
}

// HasReceiverAccountId returns a boolean if a field has been set.
func (o *TransactionCaptured) HasReceiverAccountId() bool {
	if o != nil && o.ReceiverAccountId != nil {
		return true
	}

	return false
}

// SetReceiverAccountId gets a reference to the given string and assigns it to the ReceiverAccountId field.
func (o *TransactionCaptured) SetReceiverAccountId(v string) {
	o.ReceiverAccountId = &v
}

// GetCreditAmount returns the CreditAmount field value if set, zero value otherwise.
func (o *TransactionCaptured) GetCreditAmount() int32 {
	if o == nil || o.CreditAmount == nil {
		var ret int32
		return ret
	}
	return *o.CreditAmount
}

// GetCreditAmountOk returns a tuple with the CreditAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetCreditAmountOk() (*int32, bool) {
	if o == nil || o.CreditAmount == nil {
		return nil, false
	}
	return o.CreditAmount, true
}

// HasCreditAmount returns a boolean if a field has been set.
func (o *TransactionCaptured) HasCreditAmount() bool {
	if o != nil && o.CreditAmount != nil {
		return true
	}

	return false
}

// SetCreditAmount gets a reference to the given int32 and assigns it to the CreditAmount field.
func (o *TransactionCaptured) SetCreditAmount(v int32) {
	o.CreditAmount = &v
}

// GetScheduledDate returns the ScheduledDate field value if set, zero value otherwise.
func (o *TransactionCaptured) GetScheduledDate() string {
	if o == nil || o.ScheduledDate == nil {
		var ret string
		return ret
	}
	return *o.ScheduledDate
}

// GetScheduledDateOk returns a tuple with the ScheduledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetScheduledDateOk() (*string, bool) {
	if o == nil || o.ScheduledDate == nil {
		return nil, false
	}
	return o.ScheduledDate, true
}

// HasScheduledDate returns a boolean if a field has been set.
func (o *TransactionCaptured) HasScheduledDate() bool {
	if o != nil && o.ScheduledDate != nil {
		return true
	}

	return false
}

// SetScheduledDate gets a reference to the given string and assigns it to the ScheduledDate field.
func (o *TransactionCaptured) SetScheduledDate(v string) {
	o.ScheduledDate = &v
}

// GetScheduledNumber returns the ScheduledNumber field value if set, zero value otherwise.
func (o *TransactionCaptured) GetScheduledNumber() string {
	if o == nil || o.ScheduledNumber == nil {
		var ret string
		return ret
	}
	return *o.ScheduledNumber
}

// GetScheduledNumberOk returns a tuple with the ScheduledNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetScheduledNumberOk() (*string, bool) {
	if o == nil || o.ScheduledNumber == nil {
		return nil, false
	}
	return o.ScheduledNumber, true
}

// HasScheduledNumber returns a boolean if a field has been set.
func (o *TransactionCaptured) HasScheduledNumber() bool {
	if o != nil && o.ScheduledNumber != nil {
		return true
	}

	return false
}

// SetScheduledNumber gets a reference to the given string and assigns it to the ScheduledNumber field.
func (o *TransactionCaptured) SetScheduledNumber(v string) {
	o.ScheduledNumber = &v
}

// GetMaskedLabel returns the MaskedLabel field value if set, zero value otherwise.
func (o *TransactionCaptured) GetMaskedLabel() string {
	if o == nil || o.MaskedLabel == nil {
		var ret string
		return ret
	}
	return *o.MaskedLabel
}

// GetMaskedLabelOk returns a tuple with the MaskedLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetMaskedLabelOk() (*string, bool) {
	if o == nil || o.MaskedLabel == nil {
		return nil, false
	}
	return o.MaskedLabel, true
}

// HasMaskedLabel returns a boolean if a field has been set.
func (o *TransactionCaptured) HasMaskedLabel() bool {
	if o != nil && o.MaskedLabel != nil {
		return true
	}

	return false
}

// SetMaskedLabel gets a reference to the given string and assigns it to the MaskedLabel field.
func (o *TransactionCaptured) SetMaskedLabel(v string) {
	o.MaskedLabel = &v
}

// GetPSP returns the PSP field value if set, zero value otherwise.
func (o *TransactionCaptured) GetPSP() PSP {
	if o == nil || o.PSP == nil {
		var ret PSP
		return ret
	}
	return *o.PSP
}

// GetPSPOk returns a tuple with the PSP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetPSPOk() (*PSP, bool) {
	if o == nil || o.PSP == nil {
		return nil, false
	}
	return o.PSP, true
}

// HasPSP returns a boolean if a field has been set.
func (o *TransactionCaptured) HasPSP() bool {
	if o != nil && o.PSP != nil {
		return true
	}

	return false
}

// SetPSP gets a reference to the given PSP and assigns it to the PSP field.
func (o *TransactionCaptured) SetPSP(v PSP) {
	o.PSP = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *TransactionCaptured) GetCard() Card {
	if o == nil || o.Card == nil {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetCardOk() (*Card, bool) {
	if o == nil || o.Card == nil {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *TransactionCaptured) HasCard() bool {
	if o != nil && o.Card != nil {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *TransactionCaptured) SetCard(v Card) {
	o.Card = &v
}

// GetBankStatus returns the BankStatus field value if set, zero value otherwise.
func (o *TransactionCaptured) GetBankStatus() string {
	if o == nil || o.BankStatus == nil {
		var ret string
		return ret
	}
	return *o.BankStatus
}

// GetBankStatusOk returns a tuple with the BankStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetBankStatusOk() (*string, bool) {
	if o == nil || o.BankStatus == nil {
		return nil, false
	}
	return o.BankStatus, true
}

// HasBankStatus returns a boolean if a field has been set.
func (o *TransactionCaptured) HasBankStatus() bool {
	if o != nil && o.BankStatus != nil {
		return true
	}

	return false
}

// SetBankStatus gets a reference to the given string and assigns it to the BankStatus field.
func (o *TransactionCaptured) SetBankStatus(v string) {
	o.BankStatus = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *TransactionCaptured) GetRefundAmount() float64 {
	if o == nil || o.RefundAmount == nil {
		var ret float64
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetRefundAmountOk() (*float64, bool) {
	if o == nil || o.RefundAmount == nil {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *TransactionCaptured) HasRefundAmount() bool {
	if o != nil && o.RefundAmount != nil {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given float64 and assigns it to the RefundAmount field.
func (o *TransactionCaptured) SetRefundAmount(v float64) {
	o.RefundAmount = &v
}

// GetBankReference returns the BankReference field value if set, zero value otherwise.
func (o *TransactionCaptured) GetBankReference() string {
	if o == nil || o.BankReference == nil {
		var ret string
		return ret
	}
	return *o.BankReference
}

// GetBankReferenceOk returns a tuple with the BankReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetBankReferenceOk() (*string, bool) {
	if o == nil || o.BankReference == nil {
		return nil, false
	}
	return o.BankReference, true
}

// HasBankReference returns a boolean if a field has been set.
func (o *TransactionCaptured) HasBankReference() bool {
	if o != nil && o.BankReference != nil {
		return true
	}

	return false
}

// SetBankReference gets a reference to the given string and assigns it to the BankReference field.
func (o *TransactionCaptured) SetBankReference(v string) {
	o.BankReference = &v
}

// GetChequeSendingAddressCorporateName returns the ChequeSendingAddressCorporateName field value if set, zero value otherwise.
func (o *TransactionCaptured) GetChequeSendingAddressCorporateName() string {
	if o == nil || o.ChequeSendingAddressCorporateName == nil {
		var ret string
		return ret
	}
	return *o.ChequeSendingAddressCorporateName
}

// GetChequeSendingAddressCorporateNameOk returns a tuple with the ChequeSendingAddressCorporateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetChequeSendingAddressCorporateNameOk() (*string, bool) {
	if o == nil || o.ChequeSendingAddressCorporateName == nil {
		return nil, false
	}
	return o.ChequeSendingAddressCorporateName, true
}

// HasChequeSendingAddressCorporateName returns a boolean if a field has been set.
func (o *TransactionCaptured) HasChequeSendingAddressCorporateName() bool {
	if o != nil && o.ChequeSendingAddressCorporateName != nil {
		return true
	}

	return false
}

// SetChequeSendingAddressCorporateName gets a reference to the given string and assigns it to the ChequeSendingAddressCorporateName field.
func (o *TransactionCaptured) SetChequeSendingAddressCorporateName(v string) {
	o.ChequeSendingAddressCorporateName = &v
}

// GetChequeSendingAddressStreet returns the ChequeSendingAddressStreet field value if set, zero value otherwise.
func (o *TransactionCaptured) GetChequeSendingAddressStreet() string {
	if o == nil || o.ChequeSendingAddressStreet == nil {
		var ret string
		return ret
	}
	return *o.ChequeSendingAddressStreet
}

// GetChequeSendingAddressStreetOk returns a tuple with the ChequeSendingAddressStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetChequeSendingAddressStreetOk() (*string, bool) {
	if o == nil || o.ChequeSendingAddressStreet == nil {
		return nil, false
	}
	return o.ChequeSendingAddressStreet, true
}

// HasChequeSendingAddressStreet returns a boolean if a field has been set.
func (o *TransactionCaptured) HasChequeSendingAddressStreet() bool {
	if o != nil && o.ChequeSendingAddressStreet != nil {
		return true
	}

	return false
}

// SetChequeSendingAddressStreet gets a reference to the given string and assigns it to the ChequeSendingAddressStreet field.
func (o *TransactionCaptured) SetChequeSendingAddressStreet(v string) {
	o.ChequeSendingAddressStreet = &v
}

// GetChequeSendingAddressCity returns the ChequeSendingAddressCity field value if set, zero value otherwise.
func (o *TransactionCaptured) GetChequeSendingAddressCity() string {
	if o == nil || o.ChequeSendingAddressCity == nil {
		var ret string
		return ret
	}
	return *o.ChequeSendingAddressCity
}

// GetChequeSendingAddressCityOk returns a tuple with the ChequeSendingAddressCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetChequeSendingAddressCityOk() (*string, bool) {
	if o == nil || o.ChequeSendingAddressCity == nil {
		return nil, false
	}
	return o.ChequeSendingAddressCity, true
}

// HasChequeSendingAddressCity returns a boolean if a field has been set.
func (o *TransactionCaptured) HasChequeSendingAddressCity() bool {
	if o != nil && o.ChequeSendingAddressCity != nil {
		return true
	}

	return false
}

// SetChequeSendingAddressCity gets a reference to the given string and assigns it to the ChequeSendingAddressCity field.
func (o *TransactionCaptured) SetChequeSendingAddressCity(v string) {
	o.ChequeSendingAddressCity = &v
}

// GetChequeSendingAddressPostCode returns the ChequeSendingAddressPostCode field value if set, zero value otherwise.
func (o *TransactionCaptured) GetChequeSendingAddressPostCode() string {
	if o == nil || o.ChequeSendingAddressPostCode == nil {
		var ret string
		return ret
	}
	return *o.ChequeSendingAddressPostCode
}

// GetChequeSendingAddressPostCodeOk returns a tuple with the ChequeSendingAddressPostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetChequeSendingAddressPostCodeOk() (*string, bool) {
	if o == nil || o.ChequeSendingAddressPostCode == nil {
		return nil, false
	}
	return o.ChequeSendingAddressPostCode, true
}

// HasChequeSendingAddressPostCode returns a boolean if a field has been set.
func (o *TransactionCaptured) HasChequeSendingAddressPostCode() bool {
	if o != nil && o.ChequeSendingAddressPostCode != nil {
		return true
	}

	return false
}

// SetChequeSendingAddressPostCode gets a reference to the given string and assigns it to the ChequeSendingAddressPostCode field.
func (o *TransactionCaptured) SetChequeSendingAddressPostCode(v string) {
	o.ChequeSendingAddressPostCode = &v
}

// GetThreeDS returns the ThreeDs field value if set, zero value otherwise.
func (o *TransactionCaptured) GetThreeDS() ThreeDs {
	if o == nil || o.ThreeDs == nil {
		var ret ThreeDs
		return ret
	}
	return *o.ThreeDs
}

// GetThreeDSOk returns a tuple with the ThreeDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetThreeDSOk() (*ThreeDs, bool) {
	if o == nil || o.ThreeDs == nil {
		return nil, false
	}
	return o.ThreeDs, true
}

// HasThreeDS returns a boolean if a field has been set.
func (o *TransactionCaptured) HasThreeDS() bool {
	if o != nil && o.ThreeDs != nil {
		return true
	}

	return false
}

// SetThreeDS gets a reference to the given ThreeDs and assigns it to the ThreeDs field.
func (o *TransactionCaptured) SetThreeDS(v ThreeDs) {
	o.ThreeDs = &v
}

// GetBuyNowPayLaterInfo returns the BuyNowPayLaterInfo field value if set, zero value otherwise.
func (o *TransactionCaptured) GetBuyNowPayLaterInfo() BnplInfo {
	if o == nil || o.BuyNowPayLaterInfo == nil {
		var ret BnplInfo
		return ret
	}
	return *o.BuyNowPayLaterInfo
}

// GetBuyNowPayLaterInfoOk returns a tuple with the BuyNowPayLaterInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetBuyNowPayLaterInfoOk() (*BnplInfo, bool) {
	if o == nil || o.BuyNowPayLaterInfo == nil {
		return nil, false
	}
	return o.BuyNowPayLaterInfo, true
}

// HasBuyNowPayLaterInfo returns a boolean if a field has been set.
func (o *TransactionCaptured) HasBuyNowPayLaterInfo() bool {
	if o != nil && o.BuyNowPayLaterInfo != nil {
		return true
	}

	return false
}

// SetBuyNowPayLaterInfo gets a reference to the given BnplInfo and assigns it to the BuyNowPayLaterInfo field.
func (o *TransactionCaptured) SetBuyNowPayLaterInfo(v BnplInfo) {
	o.BuyNowPayLaterInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionCaptured) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionCaptured) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TransactionCaptured) SetId(v int64) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TransactionCaptured) GetMethod() int32 {
	if o == nil || o.Method == nil {
		var ret int32
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetMethodOk() (*int32, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TransactionCaptured) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given int32 and assigns it to the Method field.
func (o *TransactionCaptured) SetMethod(v int32) {
	o.Method = &v
}

// GetMethodDetails returns the MethodDetails field value if set, zero value otherwise.
func (o *TransactionCaptured) GetMethodDetails() int32 {
	if o == nil || o.MethodDetails == nil {
		var ret int32
		return ret
	}
	return *o.MethodDetails
}

// GetMethodDetailsOk returns a tuple with the MethodDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetMethodDetailsOk() (*int32, bool) {
	if o == nil || o.MethodDetails == nil {
		return nil, false
	}
	return o.MethodDetails, true
}

// HasMethodDetails returns a boolean if a field has been set.
func (o *TransactionCaptured) HasMethodDetails() bool {
	if o != nil && o.MethodDetails != nil {
		return true
	}

	return false
}

// SetMethodDetails gets a reference to the given int32 and assigns it to the MethodDetails field.
func (o *TransactionCaptured) SetMethodDetails(v int32) {
	o.MethodDetails = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TransactionCaptured) GetDate() int32 {
	if o == nil || o.Date == nil {
		var ret int32
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetDateOk() (*int32, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TransactionCaptured) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given int32 and assigns it to the Date field.
func (o *TransactionCaptured) SetDate(v int32) {
	o.Date = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *TransactionCaptured) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *TransactionCaptured) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *TransactionCaptured) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TransactionCaptured) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TransactionCaptured) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TransactionCaptured) SetComment(v string) {
	o.Comment = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionCaptured) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionCaptured) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *TransactionCaptured) SetStatus(v int32) {
	o.Status = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *TransactionCaptured) GetExecutionDate() int32 {
	if o == nil || o.ExecutionDate == nil {
		var ret int32
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetExecutionDateOk() (*int32, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *TransactionCaptured) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given int32 and assigns it to the ExecutionDate field.
func (o *TransactionCaptured) SetExecutionDate(v int32) {
	o.ExecutionDate = &v
}

// GetLemonWayCommission returns the LemonWayCommission field value if set, zero value otherwise.
func (o *TransactionCaptured) GetLemonWayCommission() LemonWayCommission {
	if o == nil || o.LemonWayCommission == nil {
		var ret LemonWayCommission
		return ret
	}
	return *o.LemonWayCommission
}

// GetLemonWayCommissionOk returns a tuple with the LemonWayCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetLemonWayCommissionOk() (*LemonWayCommission, bool) {
	if o == nil || o.LemonWayCommission == nil {
		return nil, false
	}
	return o.LemonWayCommission, true
}

// HasLemonWayCommission returns a boolean if a field has been set.
func (o *TransactionCaptured) HasLemonWayCommission() bool {
	if o != nil && o.LemonWayCommission != nil {
		return true
	}

	return false
}

// SetLemonWayCommission gets a reference to the given LemonWayCommission and assigns it to the LemonWayCommission field.
func (o *TransactionCaptured) SetLemonWayCommission(v LemonWayCommission) {
	o.LemonWayCommission = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TransactionCaptured) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TransactionCaptured) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TransactionCaptured) SetReference(v string) {
	o.Reference = &v
}

// GetRemainingAmount returns the RemainingAmount field value if set, zero value otherwise.
func (o *TransactionCaptured) GetRemainingAmount() int32 {
	if o == nil || o.RemainingAmount == nil {
		var ret int32
		return ret
	}
	return *o.RemainingAmount
}

// GetRemainingAmountOk returns a tuple with the RemainingAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetRemainingAmountOk() (*int32, bool) {
	if o == nil || o.RemainingAmount == nil {
		return nil, false
	}
	return o.RemainingAmount, true
}

// HasRemainingAmount returns a boolean if a field has been set.
func (o *TransactionCaptured) HasRemainingAmount() bool {
	if o != nil && o.RemainingAmount != nil {
		return true
	}

	return false
}

// SetRemainingAmount gets a reference to the given int32 and assigns it to the RemainingAmount field.
func (o *TransactionCaptured) SetRemainingAmount(v int32) {
	o.RemainingAmount = &v
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *TransactionCaptured) GetAuthorization() TransactionCapturedAuthorization {
	if o == nil || o.Authorization == nil {
		var ret TransactionCapturedAuthorization
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptured) GetAuthorizationOk() (*TransactionCapturedAuthorization, bool) {
	if o == nil || o.Authorization == nil {
		return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *TransactionCaptured) HasAuthorization() bool {
	if o != nil && o.Authorization != nil {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given TransactionCapturedAuthorization and assigns it to the Authorization field.
func (o *TransactionCaptured) SetAuthorization(v TransactionCapturedAuthorization) {
	o.Authorization = &v
}

func (o TransactionCaptured) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReceiverAccountId != nil {
		toSerialize["receiverAccountId"] = o.ReceiverAccountId
	}
	if o.CreditAmount != nil {
		toSerialize["creditAmount"] = o.CreditAmount
	}
	if o.ScheduledDate != nil {
		toSerialize["scheduledDate"] = o.ScheduledDate
	}
	if o.ScheduledNumber != nil {
		toSerialize["scheduledNumber"] = o.ScheduledNumber
	}
	if o.MaskedLabel != nil {
		toSerialize["maskedLabel"] = o.MaskedLabel
	}
	if o.PSP != nil {
		toSerialize["PSP"] = o.PSP
	}
	if o.Card != nil {
		toSerialize["card"] = o.Card
	}
	if o.BankStatus != nil {
		toSerialize["bankStatus"] = o.BankStatus
	}
	if o.RefundAmount != nil {
		toSerialize["refundAmount"] = o.RefundAmount
	}
	if o.BankReference != nil {
		toSerialize["bankReference"] = o.BankReference
	}
	if o.ChequeSendingAddressCorporateName != nil {
		toSerialize["ChequeSendingAddress_CorporateName"] = o.ChequeSendingAddressCorporateName
	}
	if o.ChequeSendingAddressStreet != nil {
		toSerialize["ChequeSendingAddress_Street"] = o.ChequeSendingAddressStreet
	}
	if o.ChequeSendingAddressCity != nil {
		toSerialize["ChequeSendingAddress_City"] = o.ChequeSendingAddressCity
	}
	if o.ChequeSendingAddressPostCode != nil {
		toSerialize["ChequeSendingAddress_PostCode"] = o.ChequeSendingAddressPostCode
	}
	if o.ThreeDs != nil {
		toSerialize["threeDS"] = o.ThreeDs
	}
	if o.BuyNowPayLaterInfo != nil {
		toSerialize["buyNowPayLaterInfo"] = o.BuyNowPayLaterInfo
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.MethodDetails != nil {
		toSerialize["methodDetails"] = o.MethodDetails
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.CommissionAmount != nil {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ExecutionDate != nil {
		toSerialize["executionDate"] = o.ExecutionDate
	}
	if o.LemonWayCommission != nil {
		toSerialize["lemonWayCommission"] = o.LemonWayCommission
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.RemainingAmount != nil {
		toSerialize["remainingAmount"] = o.RemainingAmount
	}
	if o.Authorization != nil {
		toSerialize["authorization"] = o.Authorization
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionCaptured struct {
	value *TransactionCaptured
	isSet bool
}

func (v NullableTransactionCaptured) Get() *TransactionCaptured {
	return v.value
}

func (v *NullableTransactionCaptured) Set(val *TransactionCaptured) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCaptured) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCaptured) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCaptured(val *TransactionCaptured) *NullableTransactionCaptured {
	return &NullableTransactionCaptured{value: val, isSet: true}
}

func (v NullableTransactionCaptured) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCaptured) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


