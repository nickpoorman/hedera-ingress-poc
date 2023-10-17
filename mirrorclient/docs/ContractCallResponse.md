# ContractCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to ***os.File** | Result in hexadecimal from executed contract call. | [optional] 

## Methods

### NewContractCallResponse

`func NewContractCallResponse() *ContractCallResponse`

NewContractCallResponse instantiates a new ContractCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallResponseWithDefaults

`func NewContractCallResponseWithDefaults() *ContractCallResponse`

NewContractCallResponseWithDefaults instantiates a new ContractCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ContractCallResponse) GetResult() *os.File`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ContractCallResponse) GetResultOk() (**os.File, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ContractCallResponse) SetResult(v *os.File)`

SetResult sets Result field to given value.

### HasResult

`func (o *ContractCallResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


