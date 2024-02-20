# SharedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedDataId** | **string** |  | 
**SharedAmData** | Pointer to [**AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md) |  | [optional] 
**SharedSmsSubsData** | Pointer to [**SmsSubscriptionData**](SmsSubscriptionData.md) |  | [optional] 
**SharedSmsMngSubsData** | Pointer to [**SmsManagementSubscriptionData**](SmsManagementSubscriptionData.md) |  | [optional] 
**SharedDnnConfigurations** | Pointer to [**map[string]DnnConfiguration**](DnnConfiguration.md) | A map(list of key-value pairs) where Dnn, or optionally the Wildcard DNN, serves as key of DnnConfiguration | [optional] 
**SharedTraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**SharedSnssaiInfos** | Pointer to [**map[string]SnssaiInfo**](SnssaiInfo.md) | A map(list of key-value pairs) where singleNssai serves as key of SnssaiInfo | [optional] 
**SharedVnGroupDatas** | Pointer to [**map[string]VnGroupData**](VnGroupData.md) | A map(list of key-value pairs) where GroupId serves as key of VnGroupData | [optional] 
**TreatmentInstructions** | Pointer to [**map[string]SharedDataTreatmentInstruction**](SharedDataTreatmentInstruction.md) | A map(list of key-value pairs) where JSON pointer pointing to an attribute within the SharedData serves as key of SharedDataTreatmentInstruction | [optional] 
**SharedSmSubsData** | Pointer to [**SessionManagementSubscriptionData**](SessionManagementSubscriptionData.md) |  | [optional] 
**SharedEcsAddrConfigInfo** | Pointer to [**NullableEcsAddrConfigInfo**](EcsAddrConfigInfo.md) |  | [optional] 

## Methods

### NewSharedData

`func NewSharedData(sharedDataId string, ) *SharedData`

NewSharedData instantiates a new SharedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDataWithDefaults

`func NewSharedDataWithDefaults() *SharedData`

NewSharedDataWithDefaults instantiates a new SharedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedDataId

`func (o *SharedData) GetSharedDataId() string`

GetSharedDataId returns the SharedDataId field if non-nil, zero value otherwise.

### GetSharedDataIdOk

`func (o *SharedData) GetSharedDataIdOk() (*string, bool)`

GetSharedDataIdOk returns a tuple with the SharedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDataId

`func (o *SharedData) SetSharedDataId(v string)`

SetSharedDataId sets SharedDataId field to given value.


### GetSharedAmData

`func (o *SharedData) GetSharedAmData() AccessAndMobilitySubscriptionData`

GetSharedAmData returns the SharedAmData field if non-nil, zero value otherwise.

### GetSharedAmDataOk

`func (o *SharedData) GetSharedAmDataOk() (*AccessAndMobilitySubscriptionData, bool)`

GetSharedAmDataOk returns a tuple with the SharedAmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAmData

`func (o *SharedData) SetSharedAmData(v AccessAndMobilitySubscriptionData)`

SetSharedAmData sets SharedAmData field to given value.

### HasSharedAmData

`func (o *SharedData) HasSharedAmData() bool`

HasSharedAmData returns a boolean if a field has been set.

### GetSharedSmsSubsData

`func (o *SharedData) GetSharedSmsSubsData() SmsSubscriptionData`

GetSharedSmsSubsData returns the SharedSmsSubsData field if non-nil, zero value otherwise.

### GetSharedSmsSubsDataOk

`func (o *SharedData) GetSharedSmsSubsDataOk() (*SmsSubscriptionData, bool)`

GetSharedSmsSubsDataOk returns a tuple with the SharedSmsSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmsSubsData

`func (o *SharedData) SetSharedSmsSubsData(v SmsSubscriptionData)`

SetSharedSmsSubsData sets SharedSmsSubsData field to given value.

### HasSharedSmsSubsData

`func (o *SharedData) HasSharedSmsSubsData() bool`

HasSharedSmsSubsData returns a boolean if a field has been set.

### GetSharedSmsMngSubsData

`func (o *SharedData) GetSharedSmsMngSubsData() SmsManagementSubscriptionData`

GetSharedSmsMngSubsData returns the SharedSmsMngSubsData field if non-nil, zero value otherwise.

### GetSharedSmsMngSubsDataOk

`func (o *SharedData) GetSharedSmsMngSubsDataOk() (*SmsManagementSubscriptionData, bool)`

GetSharedSmsMngSubsDataOk returns a tuple with the SharedSmsMngSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmsMngSubsData

`func (o *SharedData) SetSharedSmsMngSubsData(v SmsManagementSubscriptionData)`

SetSharedSmsMngSubsData sets SharedSmsMngSubsData field to given value.

### HasSharedSmsMngSubsData

`func (o *SharedData) HasSharedSmsMngSubsData() bool`

HasSharedSmsMngSubsData returns a boolean if a field has been set.

### GetSharedDnnConfigurations

`func (o *SharedData) GetSharedDnnConfigurations() map[string]DnnConfiguration`

GetSharedDnnConfigurations returns the SharedDnnConfigurations field if non-nil, zero value otherwise.

### GetSharedDnnConfigurationsOk

`func (o *SharedData) GetSharedDnnConfigurationsOk() (*map[string]DnnConfiguration, bool)`

GetSharedDnnConfigurationsOk returns a tuple with the SharedDnnConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDnnConfigurations

`func (o *SharedData) SetSharedDnnConfigurations(v map[string]DnnConfiguration)`

SetSharedDnnConfigurations sets SharedDnnConfigurations field to given value.

### HasSharedDnnConfigurations

`func (o *SharedData) HasSharedDnnConfigurations() bool`

HasSharedDnnConfigurations returns a boolean if a field has been set.

### GetSharedTraceData

`func (o *SharedData) GetSharedTraceData() TraceData`

GetSharedTraceData returns the SharedTraceData field if non-nil, zero value otherwise.

### GetSharedTraceDataOk

`func (o *SharedData) GetSharedTraceDataOk() (*TraceData, bool)`

GetSharedTraceDataOk returns a tuple with the SharedTraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTraceData

`func (o *SharedData) SetSharedTraceData(v TraceData)`

SetSharedTraceData sets SharedTraceData field to given value.

### HasSharedTraceData

`func (o *SharedData) HasSharedTraceData() bool`

HasSharedTraceData returns a boolean if a field has been set.

### SetSharedTraceDataNil

`func (o *SharedData) SetSharedTraceDataNil(b bool)`

 SetSharedTraceDataNil sets the value for SharedTraceData to be an explicit nil

### UnsetSharedTraceData
`func (o *SharedData) UnsetSharedTraceData()`

UnsetSharedTraceData ensures that no value is present for SharedTraceData, not even an explicit nil
### GetSharedSnssaiInfos

`func (o *SharedData) GetSharedSnssaiInfos() map[string]SnssaiInfo`

GetSharedSnssaiInfos returns the SharedSnssaiInfos field if non-nil, zero value otherwise.

### GetSharedSnssaiInfosOk

`func (o *SharedData) GetSharedSnssaiInfosOk() (*map[string]SnssaiInfo, bool)`

GetSharedSnssaiInfosOk returns a tuple with the SharedSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSnssaiInfos

`func (o *SharedData) SetSharedSnssaiInfos(v map[string]SnssaiInfo)`

SetSharedSnssaiInfos sets SharedSnssaiInfos field to given value.

### HasSharedSnssaiInfos

`func (o *SharedData) HasSharedSnssaiInfos() bool`

HasSharedSnssaiInfos returns a boolean if a field has been set.

### GetSharedVnGroupDatas

`func (o *SharedData) GetSharedVnGroupDatas() map[string]VnGroupData`

GetSharedVnGroupDatas returns the SharedVnGroupDatas field if non-nil, zero value otherwise.

### GetSharedVnGroupDatasOk

`func (o *SharedData) GetSharedVnGroupDatasOk() (*map[string]VnGroupData, bool)`

GetSharedVnGroupDatasOk returns a tuple with the SharedVnGroupDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVnGroupDatas

`func (o *SharedData) SetSharedVnGroupDatas(v map[string]VnGroupData)`

SetSharedVnGroupDatas sets SharedVnGroupDatas field to given value.

### HasSharedVnGroupDatas

`func (o *SharedData) HasSharedVnGroupDatas() bool`

HasSharedVnGroupDatas returns a boolean if a field has been set.

### GetTreatmentInstructions

`func (o *SharedData) GetTreatmentInstructions() map[string]SharedDataTreatmentInstruction`

GetTreatmentInstructions returns the TreatmentInstructions field if non-nil, zero value otherwise.

### GetTreatmentInstructionsOk

`func (o *SharedData) GetTreatmentInstructionsOk() (*map[string]SharedDataTreatmentInstruction, bool)`

GetTreatmentInstructionsOk returns a tuple with the TreatmentInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentInstructions

`func (o *SharedData) SetTreatmentInstructions(v map[string]SharedDataTreatmentInstruction)`

SetTreatmentInstructions sets TreatmentInstructions field to given value.

### HasTreatmentInstructions

`func (o *SharedData) HasTreatmentInstructions() bool`

HasTreatmentInstructions returns a boolean if a field has been set.

### GetSharedSmSubsData

`func (o *SharedData) GetSharedSmSubsData() SessionManagementSubscriptionData`

GetSharedSmSubsData returns the SharedSmSubsData field if non-nil, zero value otherwise.

### GetSharedSmSubsDataOk

`func (o *SharedData) GetSharedSmSubsDataOk() (*SessionManagementSubscriptionData, bool)`

GetSharedSmSubsDataOk returns a tuple with the SharedSmSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmSubsData

`func (o *SharedData) SetSharedSmSubsData(v SessionManagementSubscriptionData)`

SetSharedSmSubsData sets SharedSmSubsData field to given value.

### HasSharedSmSubsData

`func (o *SharedData) HasSharedSmSubsData() bool`

HasSharedSmSubsData returns a boolean if a field has been set.

### GetSharedEcsAddrConfigInfo

`func (o *SharedData) GetSharedEcsAddrConfigInfo() EcsAddrConfigInfo`

GetSharedEcsAddrConfigInfo returns the SharedEcsAddrConfigInfo field if non-nil, zero value otherwise.

### GetSharedEcsAddrConfigInfoOk

`func (o *SharedData) GetSharedEcsAddrConfigInfoOk() (*EcsAddrConfigInfo, bool)`

GetSharedEcsAddrConfigInfoOk returns a tuple with the SharedEcsAddrConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedEcsAddrConfigInfo

`func (o *SharedData) SetSharedEcsAddrConfigInfo(v EcsAddrConfigInfo)`

SetSharedEcsAddrConfigInfo sets SharedEcsAddrConfigInfo field to given value.

### HasSharedEcsAddrConfigInfo

`func (o *SharedData) HasSharedEcsAddrConfigInfo() bool`

HasSharedEcsAddrConfigInfo returns a boolean if a field has been set.

### SetSharedEcsAddrConfigInfoNil

`func (o *SharedData) SetSharedEcsAddrConfigInfoNil(b bool)`

 SetSharedEcsAddrConfigInfoNil sets the value for SharedEcsAddrConfigInfo to be an explicit nil

### UnsetSharedEcsAddrConfigInfo
`func (o *SharedData) UnsetSharedEcsAddrConfigInfo()`

UnsetSharedEcsAddrConfigInfo ensures that no value is present for SharedEcsAddrConfigInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


