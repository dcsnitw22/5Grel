# UeContextInSmfDataSubFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnnList** | Pointer to **[]string** |  | [optional] 
**SnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**EmergencyInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUeContextInSmfDataSubFilter

`func NewUeContextInSmfDataSubFilter() *UeContextInSmfDataSubFilter`

NewUeContextInSmfDataSubFilter instantiates a new UeContextInSmfDataSubFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextInSmfDataSubFilterWithDefaults

`func NewUeContextInSmfDataSubFilterWithDefaults() *UeContextInSmfDataSubFilter`

NewUeContextInSmfDataSubFilterWithDefaults instantiates a new UeContextInSmfDataSubFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnnList

`func (o *UeContextInSmfDataSubFilter) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *UeContextInSmfDataSubFilter) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *UeContextInSmfDataSubFilter) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *UeContextInSmfDataSubFilter) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetSnssaiList

`func (o *UeContextInSmfDataSubFilter) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *UeContextInSmfDataSubFilter) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *UeContextInSmfDataSubFilter) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.

### HasSnssaiList

`func (o *UeContextInSmfDataSubFilter) HasSnssaiList() bool`

HasSnssaiList returns a boolean if a field has been set.

### GetEmergencyInd

`func (o *UeContextInSmfDataSubFilter) GetEmergencyInd() bool`

GetEmergencyInd returns the EmergencyInd field if non-nil, zero value otherwise.

### GetEmergencyIndOk

`func (o *UeContextInSmfDataSubFilter) GetEmergencyIndOk() (*bool, bool)`

GetEmergencyIndOk returns a tuple with the EmergencyInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyInd

`func (o *UeContextInSmfDataSubFilter) SetEmergencyInd(v bool)`

SetEmergencyInd sets EmergencyInd field to given value.

### HasEmergencyInd

`func (o *UeContextInSmfDataSubFilter) HasEmergencyInd() bool`

HasEmergencyInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


