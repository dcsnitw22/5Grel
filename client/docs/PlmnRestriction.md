# PlmnRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**ForbiddenAreas** | Pointer to [**[]Area**](Area.md) |  | [optional] 
**ServiceAreaRestriction** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**CoreNetworkTypeRestrictions** | Pointer to [**[]CoreNetworkType**](CoreNetworkType.md) |  | [optional] 
**AccessTypeRestrictions** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**PrimaryRatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**SecondaryRatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 

## Methods

### NewPlmnRestriction

`func NewPlmnRestriction() *PlmnRestriction`

NewPlmnRestriction instantiates a new PlmnRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlmnRestrictionWithDefaults

`func NewPlmnRestrictionWithDefaults() *PlmnRestriction`

NewPlmnRestrictionWithDefaults instantiates a new PlmnRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatRestrictions

`func (o *PlmnRestriction) GetRatRestrictions() []RatType`

GetRatRestrictions returns the RatRestrictions field if non-nil, zero value otherwise.

### GetRatRestrictionsOk

`func (o *PlmnRestriction) GetRatRestrictionsOk() (*[]RatType, bool)`

GetRatRestrictionsOk returns a tuple with the RatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatRestrictions

`func (o *PlmnRestriction) SetRatRestrictions(v []RatType)`

SetRatRestrictions sets RatRestrictions field to given value.

### HasRatRestrictions

`func (o *PlmnRestriction) HasRatRestrictions() bool`

HasRatRestrictions returns a boolean if a field has been set.

### GetForbiddenAreas

`func (o *PlmnRestriction) GetForbiddenAreas() []Area`

GetForbiddenAreas returns the ForbiddenAreas field if non-nil, zero value otherwise.

### GetForbiddenAreasOk

`func (o *PlmnRestriction) GetForbiddenAreasOk() (*[]Area, bool)`

GetForbiddenAreasOk returns a tuple with the ForbiddenAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAreas

`func (o *PlmnRestriction) SetForbiddenAreas(v []Area)`

SetForbiddenAreas sets ForbiddenAreas field to given value.

### HasForbiddenAreas

`func (o *PlmnRestriction) HasForbiddenAreas() bool`

HasForbiddenAreas returns a boolean if a field has been set.

### GetServiceAreaRestriction

`func (o *PlmnRestriction) GetServiceAreaRestriction() ServiceAreaRestriction`

GetServiceAreaRestriction returns the ServiceAreaRestriction field if non-nil, zero value otherwise.

### GetServiceAreaRestrictionOk

`func (o *PlmnRestriction) GetServiceAreaRestrictionOk() (*ServiceAreaRestriction, bool)`

GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreaRestriction

`func (o *PlmnRestriction) SetServiceAreaRestriction(v ServiceAreaRestriction)`

SetServiceAreaRestriction sets ServiceAreaRestriction field to given value.

### HasServiceAreaRestriction

`func (o *PlmnRestriction) HasServiceAreaRestriction() bool`

HasServiceAreaRestriction returns a boolean if a field has been set.

### GetCoreNetworkTypeRestrictions

`func (o *PlmnRestriction) GetCoreNetworkTypeRestrictions() []CoreNetworkType`

GetCoreNetworkTypeRestrictions returns the CoreNetworkTypeRestrictions field if non-nil, zero value otherwise.

### GetCoreNetworkTypeRestrictionsOk

`func (o *PlmnRestriction) GetCoreNetworkTypeRestrictionsOk() (*[]CoreNetworkType, bool)`

GetCoreNetworkTypeRestrictionsOk returns a tuple with the CoreNetworkTypeRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreNetworkTypeRestrictions

`func (o *PlmnRestriction) SetCoreNetworkTypeRestrictions(v []CoreNetworkType)`

SetCoreNetworkTypeRestrictions sets CoreNetworkTypeRestrictions field to given value.

### HasCoreNetworkTypeRestrictions

`func (o *PlmnRestriction) HasCoreNetworkTypeRestrictions() bool`

HasCoreNetworkTypeRestrictions returns a boolean if a field has been set.

### GetAccessTypeRestrictions

`func (o *PlmnRestriction) GetAccessTypeRestrictions() []AccessType`

GetAccessTypeRestrictions returns the AccessTypeRestrictions field if non-nil, zero value otherwise.

### GetAccessTypeRestrictionsOk

`func (o *PlmnRestriction) GetAccessTypeRestrictionsOk() (*[]AccessType, bool)`

GetAccessTypeRestrictionsOk returns a tuple with the AccessTypeRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypeRestrictions

`func (o *PlmnRestriction) SetAccessTypeRestrictions(v []AccessType)`

SetAccessTypeRestrictions sets AccessTypeRestrictions field to given value.

### HasAccessTypeRestrictions

`func (o *PlmnRestriction) HasAccessTypeRestrictions() bool`

HasAccessTypeRestrictions returns a boolean if a field has been set.

### GetPrimaryRatRestrictions

`func (o *PlmnRestriction) GetPrimaryRatRestrictions() []RatType`

GetPrimaryRatRestrictions returns the PrimaryRatRestrictions field if non-nil, zero value otherwise.

### GetPrimaryRatRestrictionsOk

`func (o *PlmnRestriction) GetPrimaryRatRestrictionsOk() (*[]RatType, bool)`

GetPrimaryRatRestrictionsOk returns a tuple with the PrimaryRatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRatRestrictions

`func (o *PlmnRestriction) SetPrimaryRatRestrictions(v []RatType)`

SetPrimaryRatRestrictions sets PrimaryRatRestrictions field to given value.

### HasPrimaryRatRestrictions

`func (o *PlmnRestriction) HasPrimaryRatRestrictions() bool`

HasPrimaryRatRestrictions returns a boolean if a field has been set.

### GetSecondaryRatRestrictions

`func (o *PlmnRestriction) GetSecondaryRatRestrictions() []RatType`

GetSecondaryRatRestrictions returns the SecondaryRatRestrictions field if non-nil, zero value otherwise.

### GetSecondaryRatRestrictionsOk

`func (o *PlmnRestriction) GetSecondaryRatRestrictionsOk() (*[]RatType, bool)`

GetSecondaryRatRestrictionsOk returns a tuple with the SecondaryRatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatRestrictions

`func (o *PlmnRestriction) SetSecondaryRatRestrictions(v []RatType)`

SetSecondaryRatRestrictions sets SecondaryRatRestrictions field to given value.

### HasSecondaryRatRestrictions

`func (o *PlmnRestriction) HasSecondaryRatRestrictions() bool`

HasSecondaryRatRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


