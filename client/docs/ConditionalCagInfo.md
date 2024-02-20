# ConditionalCagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCagList** | **[]string** |  | 
**CagOnlyIndicator** | Pointer to **bool** |  | [optional] 
**ValidTimePeriod** | Pointer to [**ValidTimePeriod**](ValidTimePeriod.md) |  | [optional] 

## Methods

### NewConditionalCagInfo

`func NewConditionalCagInfo(allowedCagList []string, ) *ConditionalCagInfo`

NewConditionalCagInfo instantiates a new ConditionalCagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalCagInfoWithDefaults

`func NewConditionalCagInfoWithDefaults() *ConditionalCagInfo`

NewConditionalCagInfoWithDefaults instantiates a new ConditionalCagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedCagList

`func (o *ConditionalCagInfo) GetAllowedCagList() []string`

GetAllowedCagList returns the AllowedCagList field if non-nil, zero value otherwise.

### GetAllowedCagListOk

`func (o *ConditionalCagInfo) GetAllowedCagListOk() (*[]string, bool)`

GetAllowedCagListOk returns a tuple with the AllowedCagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCagList

`func (o *ConditionalCagInfo) SetAllowedCagList(v []string)`

SetAllowedCagList sets AllowedCagList field to given value.


### GetCagOnlyIndicator

`func (o *ConditionalCagInfo) GetCagOnlyIndicator() bool`

GetCagOnlyIndicator returns the CagOnlyIndicator field if non-nil, zero value otherwise.

### GetCagOnlyIndicatorOk

`func (o *ConditionalCagInfo) GetCagOnlyIndicatorOk() (*bool, bool)`

GetCagOnlyIndicatorOk returns a tuple with the CagOnlyIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagOnlyIndicator

`func (o *ConditionalCagInfo) SetCagOnlyIndicator(v bool)`

SetCagOnlyIndicator sets CagOnlyIndicator field to given value.

### HasCagOnlyIndicator

`func (o *ConditionalCagInfo) HasCagOnlyIndicator() bool`

HasCagOnlyIndicator returns a boolean if a field has been set.

### GetValidTimePeriod

`func (o *ConditionalCagInfo) GetValidTimePeriod() ValidTimePeriod`

GetValidTimePeriod returns the ValidTimePeriod field if non-nil, zero value otherwise.

### GetValidTimePeriodOk

`func (o *ConditionalCagInfo) GetValidTimePeriodOk() (*ValidTimePeriod, bool)`

GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTimePeriod

`func (o *ConditionalCagInfo) SetValidTimePeriod(v ValidTimePeriod)`

SetValidTimePeriod sets ValidTimePeriod field to given value.

### HasValidTimePeriod

`func (o *ConditionalCagInfo) HasValidTimePeriod() bool`

HasValidTimePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


