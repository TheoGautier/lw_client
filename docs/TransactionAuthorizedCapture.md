# TransactionAuthorizedCapture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Transaction ID | [optional] 
**Date** | Pointer to **int32** | Transaction initialization date, UTC Unix timestamp | [optional] 
**CommissionAmount** | Pointer to **int32** | Your fee  Amounts are given as integer numbers in cents | [optional] 
**Comment** | Pointer to **string** | Comment | [optional] 
**LemonWayCommission** | Pointer to [**LemonWayCommission**](LemonWayCommission.md) |  | [optional] 
**Amount** | Pointer to **int32** | Transaction Amount  Amounts represented in integer (cents) | [optional] 
**RefundAmount** | Pointer to **float64** |  | [optional] 
**ExecutionDate** | Pointer to **int32** | Transaction execution date | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**PSP** | Pointer to [**PSP**](PSP.md) |  | [optional] 
**Reference** | Pointer to **string** | Unique ID generated by your server | [optional] 

## Methods

### NewTransactionAuthorizedCapture

`func NewTransactionAuthorizedCapture() *TransactionAuthorizedCapture`

NewTransactionAuthorizedCapture instantiates a new TransactionAuthorizedCapture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAuthorizedCaptureWithDefaults

`func NewTransactionAuthorizedCaptureWithDefaults() *TransactionAuthorizedCapture`

NewTransactionAuthorizedCaptureWithDefaults instantiates a new TransactionAuthorizedCapture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionAuthorizedCapture) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionAuthorizedCapture) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionAuthorizedCapture) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionAuthorizedCapture) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *TransactionAuthorizedCapture) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionAuthorizedCapture) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionAuthorizedCapture) SetDate(v int32)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionAuthorizedCapture) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCommissionAmount

`func (o *TransactionAuthorizedCapture) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *TransactionAuthorizedCapture) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *TransactionAuthorizedCapture) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.

### HasCommissionAmount

`func (o *TransactionAuthorizedCapture) HasCommissionAmount() bool`

HasCommissionAmount returns a boolean if a field has been set.

### GetComment

`func (o *TransactionAuthorizedCapture) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TransactionAuthorizedCapture) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TransactionAuthorizedCapture) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TransactionAuthorizedCapture) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLemonWayCommission

`func (o *TransactionAuthorizedCapture) GetLemonWayCommission() LemonWayCommission`

GetLemonWayCommission returns the LemonWayCommission field if non-nil, zero value otherwise.

### GetLemonWayCommissionOk

`func (o *TransactionAuthorizedCapture) GetLemonWayCommissionOk() (*LemonWayCommission, bool)`

GetLemonWayCommissionOk returns a tuple with the LemonWayCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLemonWayCommission

`func (o *TransactionAuthorizedCapture) SetLemonWayCommission(v LemonWayCommission)`

SetLemonWayCommission sets LemonWayCommission field to given value.

### HasLemonWayCommission

`func (o *TransactionAuthorizedCapture) HasLemonWayCommission() bool`

HasLemonWayCommission returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionAuthorizedCapture) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionAuthorizedCapture) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionAuthorizedCapture) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionAuthorizedCapture) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRefundAmount

`func (o *TransactionAuthorizedCapture) GetRefundAmount() float64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *TransactionAuthorizedCapture) GetRefundAmountOk() (*float64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *TransactionAuthorizedCapture) SetRefundAmount(v float64)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *TransactionAuthorizedCapture) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetExecutionDate

`func (o *TransactionAuthorizedCapture) GetExecutionDate() int32`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *TransactionAuthorizedCapture) GetExecutionDateOk() (*int32, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *TransactionAuthorizedCapture) SetExecutionDate(v int32)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *TransactionAuthorizedCapture) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionAuthorizedCapture) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionAuthorizedCapture) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionAuthorizedCapture) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionAuthorizedCapture) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPSP

`func (o *TransactionAuthorizedCapture) GetPSP() PSP`

GetPSP returns the PSP field if non-nil, zero value otherwise.

### GetPSPOk

`func (o *TransactionAuthorizedCapture) GetPSPOk() (*PSP, bool)`

GetPSPOk returns a tuple with the PSP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSP

`func (o *TransactionAuthorizedCapture) SetPSP(v PSP)`

SetPSP sets PSP field to given value.

### HasPSP

`func (o *TransactionAuthorizedCapture) HasPSP() bool`

HasPSP returns a boolean if a field has been set.

### GetReference

`func (o *TransactionAuthorizedCapture) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransactionAuthorizedCapture) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransactionAuthorizedCapture) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TransactionAuthorizedCapture) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


