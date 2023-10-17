# TokenRelationshipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | Pointer to [**[]TokenRelationship**](TokenRelationship.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewTokenRelationshipResponse

`func NewTokenRelationshipResponse() *TokenRelationshipResponse`

NewTokenRelationshipResponse instantiates a new TokenRelationshipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRelationshipResponseWithDefaults

`func NewTokenRelationshipResponseWithDefaults() *TokenRelationshipResponse`

NewTokenRelationshipResponseWithDefaults instantiates a new TokenRelationshipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *TokenRelationshipResponse) GetTokens() []TokenRelationship`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TokenRelationshipResponse) GetTokensOk() (*[]TokenRelationship, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TokenRelationshipResponse) SetTokens(v []TokenRelationship)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *TokenRelationshipResponse) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetLinks

`func (o *TokenRelationshipResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokenRelationshipResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokenRelationshipResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokenRelationshipResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


