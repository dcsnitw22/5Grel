# ExtendedSmSubsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSmSubsDataIds** | **[]string** |  | 
**IndividualSmSubsData** | Pointer to [**[]SessionManagementSubscriptionData**](SessionManagementSubscriptionData.md) |  | [optional] 

## Methods

### NewExtendedSmSubsData

`func NewExtendedSmSubsData(sharedSmSubsDataIds []string, ) *ExtendedSmSubsData`

NewExtendedSmSubsData instantiates a new ExtendedSmSubsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedSmSubsDataWithDefaults

`func NewExtendedSmSubsDataWithDefaults() *ExtendedSmSubsData`

NewExtendedSmSubsDataWithDefaults instantiates a new ExtendedSmSubsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSmSubsDataIds

`func (o *ExtendedSmSubsData) GetSharedSmSubsDataIds() []string`

GetSharedSmSubsDataIds returns the SharedSmSubsDataIds field if non-nil, zero value otherwise.

### GetSharedSmSubsDataIdsOk

`func (o *ExtendedSmSubsData) GetSharedSmSubsDataIdsOk() (*[]string, bool)`

GetSharedSmSubsDataIdsOk returns a tuple with the SharedSmSubsDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmSubsDataIds

`func (o *ExtendedSmSubsData) SetSharedSmSubsDataIds(v []string)`

SetSharedSmSubsDataIds sets SharedSmSubsDataIds field to given value.


### GetIndividualSmSubsData

`func (o *ExtendedSmSubsData) GetIndividualSmSubsData() []SessionManagementSubscriptionData`

GetIndividualSmSubsData returns the IndividualSmSubsData field if non-nil, zero value otherwise.

### GetIndividualSmSubsDataOk

`func (o *ExtendedSmSubsData) GetIndividualSmSubsDataOk() (*[]SessionManagementSubscriptionData, bool)`

GetIndividualSmSubsDataOk returns a tuple with the IndividualSmSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualSmSubsData

`func (o *ExtendedSmSubsData) SetIndividualSmSubsData(v []SessionManagementSubscriptionData)`

SetIndividualSmSubsData sets IndividualSmSubsData field to given value.

### HasIndividualSmSubsData

`func (o *ExtendedSmSubsData) HasIndividualSmSubsData() bool`

HasIndividualSmSubsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


