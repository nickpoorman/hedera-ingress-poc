# ErrorStatusMessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **Nullable*os.File** | Error message in hexadecimal | [optional] 
**Detail** | Pointer to **NullableString** | Detailed error message | [optional] 
**Message** | Pointer to **string** | Error message | [optional] 

## Methods

### NewErrorStatusMessagesInner

`func NewErrorStatusMessagesInner() *ErrorStatusMessagesInner`

NewErrorStatusMessagesInner instantiates a new ErrorStatusMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorStatusMessagesInnerWithDefaults

`func NewErrorStatusMessagesInnerWithDefaults() *ErrorStatusMessagesInner`

NewErrorStatusMessagesInnerWithDefaults instantiates a new ErrorStatusMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ErrorStatusMessagesInner) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ErrorStatusMessagesInner) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ErrorStatusMessagesInner) SetData(v *os.File)`

SetData sets Data field to given value.

### HasData

`func (o *ErrorStatusMessagesInner) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ErrorStatusMessagesInner) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ErrorStatusMessagesInner) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDetail

`func (o *ErrorStatusMessagesInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorStatusMessagesInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorStatusMessagesInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorStatusMessagesInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ErrorStatusMessagesInner) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ErrorStatusMessagesInner) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetMessage

`func (o *ErrorStatusMessagesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorStatusMessagesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorStatusMessagesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorStatusMessagesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


