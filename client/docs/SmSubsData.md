# SmSubsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSmSubsDataIds** | **[]string** |  | 
**IndividualSmSubsData** | Pointer to [**[]SessionManagementSubscriptionData**](SessionManagementSubscriptionData.md) |  | [optional] 

## Methods

### NewSmSubsData

`func NewSmSubsData(sharedSmSubsDataIds []string, ) *SmSubsData`

NewSmSubsData instantiates a new SmSubsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmSubsDataWithDefaults

`func NewSmSubsDataWithDefaults() *SmSubsData`

NewSmSubsDataWithDefaults instantiates a new SmSubsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSmSubsDataIds

`func (o *SmSubsData) GetSharedSmSubsDataIds() []string`

GetSharedSmSubsDataIds returns the SharedSmSubsDataIds field if non-nil, zero value otherwise.

### GetSharedSmSubsDataIdsOk

`func (o *SmSubsData) GetSharedSmSubsDataIdsOk() (*[]string, bool)`

GetSharedSmSubsDataIdsOk returns a tuple with the SharedSmSubsDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmSubsDataIds

`func (o *SmSubsData) SetSharedSmSubsDataIds(v []string)`

SetSharedSmSubsDataIds sets SharedSmSubsDataIds field to given value.


### GetIndividualSmSubsData

`func (o *SmSubsData) GetIndividualSmSubsData() []SessionManagementSubscriptionData`

GetIndividualSmSubsData returns the IndividualSmSubsData field if non-nil, zero value otherwise.

### GetIndividualSmSubsDataOk

`func (o *SmSubsData) GetIndividualSmSubsDataOk() (*[]SessionManagementSubscriptionData, bool)`

GetIndividualSmSubsDataOk returns a tuple with the IndividualSmSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualSmSubsData

`func (o *SmSubsData) SetIndividualSmSubsData(v []SessionManagementSubscriptionData)`

SetIndividualSmSubsData sets IndividualSmSubsData field to given value.

### HasIndividualSmSubsData

`func (o *SmSubsData) HasIndividualSmSubsData() bool`

HasIndividualSmSubsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


