# TokenAllowancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowances** | Pointer to [**[]TokenAllowance**](TokenAllowance.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewTokenAllowancesResponse

`func NewTokenAllowancesResponse() *TokenAllowancesResponse`

NewTokenAllowancesResponse instantiates a new TokenAllowancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAllowancesResponseWithDefaults

`func NewTokenAllowancesResponseWithDefaults() *TokenAllowancesResponse`

NewTokenAllowancesResponseWithDefaults instantiates a new TokenAllowancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowances

`func (o *TokenAllowancesResponse) GetAllowances() []TokenAllowance`

GetAllowances returns the Allowances field if non-nil, zero value otherwise.

### GetAllowancesOk

`func (o *TokenAllowancesResponse) GetAllowancesOk() (*[]TokenAllowance, bool)`

GetAllowancesOk returns a tuple with the Allowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowances

`func (o *TokenAllowancesResponse) SetAllowances(v []TokenAllowance)`

SetAllowances sets Allowances field to given value.

### HasAllowances

`func (o *TokenAllowancesResponse) HasAllowances() bool`

HasAllowances returns a boolean if a field has been set.

### GetLinks

`func (o *TokenAllowancesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokenAllowancesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokenAllowancesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokenAllowancesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


