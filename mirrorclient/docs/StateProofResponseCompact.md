# StateProofResponseCompact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressBooks** | **[]string** | The network address book valid at the time of the transaction | 
**RecordFile** | [**StateProofResponseCompactRecordFile**](StateProofResponseCompactRecordFile.md) |  | 
**SignatureFiles** | **map[string]string** | The nodes&#39; signature files for the record file | 
**Version** | **int32** | The record file format version, either 5 or 6 | 

## Methods

### NewStateProofResponseCompact

`func NewStateProofResponseCompact(addressBooks []string, recordFile StateProofResponseCompactRecordFile, signatureFiles map[string]string, version int32, ) *StateProofResponseCompact`

NewStateProofResponseCompact instantiates a new StateProofResponseCompact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateProofResponseCompactWithDefaults

`func NewStateProofResponseCompactWithDefaults() *StateProofResponseCompact`

NewStateProofResponseCompactWithDefaults instantiates a new StateProofResponseCompact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressBooks

`func (o *StateProofResponseCompact) GetAddressBooks() []string`

GetAddressBooks returns the AddressBooks field if non-nil, zero value otherwise.

### GetAddressBooksOk

`func (o *StateProofResponseCompact) GetAddressBooksOk() (*[]string, bool)`

GetAddressBooksOk returns a tuple with the AddressBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBooks

`func (o *StateProofResponseCompact) SetAddressBooks(v []string)`

SetAddressBooks sets AddressBooks field to given value.


### GetRecordFile

`func (o *StateProofResponseCompact) GetRecordFile() StateProofResponseCompactRecordFile`

GetRecordFile returns the RecordFile field if non-nil, zero value otherwise.

### GetRecordFileOk

`func (o *StateProofResponseCompact) GetRecordFileOk() (*StateProofResponseCompactRecordFile, bool)`

GetRecordFileOk returns a tuple with the RecordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFile

`func (o *StateProofResponseCompact) SetRecordFile(v StateProofResponseCompactRecordFile)`

SetRecordFile sets RecordFile field to given value.


### GetSignatureFiles

`func (o *StateProofResponseCompact) GetSignatureFiles() map[string]string`

GetSignatureFiles returns the SignatureFiles field if non-nil, zero value otherwise.

### GetSignatureFilesOk

`func (o *StateProofResponseCompact) GetSignatureFilesOk() (*map[string]string, bool)`

GetSignatureFilesOk returns a tuple with the SignatureFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureFiles

`func (o *StateProofResponseCompact) SetSignatureFiles(v map[string]string)`

SetSignatureFiles sets SignatureFiles field to given value.


### GetVersion

`func (o *StateProofResponseCompact) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StateProofResponseCompact) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StateProofResponseCompact) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


