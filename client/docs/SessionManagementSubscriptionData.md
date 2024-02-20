# SessionManagementSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleNssai** | [**Snssai**](Snssai.md) |  | 
**DnnConfigurations** | Pointer to [**map[string]DnnConfiguration**](DnnConfiguration.md) | A map (list of key-value pairs where Dnn, or optionally the Wildcard DNN, serves as key) of DnnConfigurations | [optional] 
**InternalGroupIds** | Pointer to **[]string** |  | [optional] 
**SharedVnGroupDataIds** | Pointer to **map[string]string** | A map(list of key-value pairs) where GroupId serves as key of SharedDataId | [optional] 
**SharedDnnConfigurationsId** | Pointer to **string** |  | [optional] 
**OdbPacketServices** | Pointer to [**NullableOdbPacketServices**](OdbPacketServices.md) |  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**SharedTraceDataId** | Pointer to **string** |  | [optional] 
**ExpectedUeBehavioursList** | Pointer to [**map[string]ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) | A map(list of key-value pairs) where Dnn serves as key of ExpectedUeBehaviourData | [optional] 
**ExpectedUeBehaviourData** | Pointer to [**map[string]map[string]ExpectedUeBehaviourData**](map.md) | A map(list of key-value pairs) where DNN serves as key | [optional] 
**AppSpecificExpectedUeBehaviourData** | Pointer to [**map[string]map[string]AppSpecificExpectedUeBehaviourData**](map.md) | A map(list of key-value pairs) where DNN serves as key | [optional] 
**SuggestedPacketNumDlList** | Pointer to [**map[string]SuggestedPacketNumDl**](SuggestedPacketNumDl.md) | A map(list of key-value pairs) where Dnn serves as key of SuggestedPacketNumDl | [optional] 
**Var3gppChargingCharacteristics** | Pointer to **string** |  | [optional] 
**NsacMode** | Pointer to [**NsacAdmissionMode**](NsacAdmissionMode.md) |  | [optional] 
**SessInactTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**OnDemand** | Pointer to **bool** |  | [optional] [default to false]
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AdditionalSharedDnnConfigurationsIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSessionManagementSubscriptionData

`func NewSessionManagementSubscriptionData(singleNssai Snssai, ) *SessionManagementSubscriptionData`

NewSessionManagementSubscriptionData instantiates a new SessionManagementSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionManagementSubscriptionDataWithDefaults

`func NewSessionManagementSubscriptionDataWithDefaults() *SessionManagementSubscriptionData`

NewSessionManagementSubscriptionDataWithDefaults instantiates a new SessionManagementSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleNssai

`func (o *SessionManagementSubscriptionData) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *SessionManagementSubscriptionData) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *SessionManagementSubscriptionData) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.


### GetDnnConfigurations

`func (o *SessionManagementSubscriptionData) GetDnnConfigurations() map[string]DnnConfiguration`

GetDnnConfigurations returns the DnnConfigurations field if non-nil, zero value otherwise.

### GetDnnConfigurationsOk

`func (o *SessionManagementSubscriptionData) GetDnnConfigurationsOk() (*map[string]DnnConfiguration, bool)`

GetDnnConfigurationsOk returns a tuple with the DnnConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnConfigurations

`func (o *SessionManagementSubscriptionData) SetDnnConfigurations(v map[string]DnnConfiguration)`

SetDnnConfigurations sets DnnConfigurations field to given value.

### HasDnnConfigurations

`func (o *SessionManagementSubscriptionData) HasDnnConfigurations() bool`

HasDnnConfigurations returns a boolean if a field has been set.

### GetInternalGroupIds

`func (o *SessionManagementSubscriptionData) GetInternalGroupIds() []string`

GetInternalGroupIds returns the InternalGroupIds field if non-nil, zero value otherwise.

### GetInternalGroupIdsOk

`func (o *SessionManagementSubscriptionData) GetInternalGroupIdsOk() (*[]string, bool)`

GetInternalGroupIdsOk returns a tuple with the InternalGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIds

`func (o *SessionManagementSubscriptionData) SetInternalGroupIds(v []string)`

SetInternalGroupIds sets InternalGroupIds field to given value.

### HasInternalGroupIds

`func (o *SessionManagementSubscriptionData) HasInternalGroupIds() bool`

HasInternalGroupIds returns a boolean if a field has been set.

### GetSharedVnGroupDataIds

`func (o *SessionManagementSubscriptionData) GetSharedVnGroupDataIds() map[string]string`

GetSharedVnGroupDataIds returns the SharedVnGroupDataIds field if non-nil, zero value otherwise.

### GetSharedVnGroupDataIdsOk

`func (o *SessionManagementSubscriptionData) GetSharedVnGroupDataIdsOk() (*map[string]string, bool)`

GetSharedVnGroupDataIdsOk returns a tuple with the SharedVnGroupDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVnGroupDataIds

`func (o *SessionManagementSubscriptionData) SetSharedVnGroupDataIds(v map[string]string)`

SetSharedVnGroupDataIds sets SharedVnGroupDataIds field to given value.

### HasSharedVnGroupDataIds

`func (o *SessionManagementSubscriptionData) HasSharedVnGroupDataIds() bool`

HasSharedVnGroupDataIds returns a boolean if a field has been set.

### GetSharedDnnConfigurationsId

`func (o *SessionManagementSubscriptionData) GetSharedDnnConfigurationsId() string`

GetSharedDnnConfigurationsId returns the SharedDnnConfigurationsId field if non-nil, zero value otherwise.

### GetSharedDnnConfigurationsIdOk

`func (o *SessionManagementSubscriptionData) GetSharedDnnConfigurationsIdOk() (*string, bool)`

GetSharedDnnConfigurationsIdOk returns a tuple with the SharedDnnConfigurationsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDnnConfigurationsId

`func (o *SessionManagementSubscriptionData) SetSharedDnnConfigurationsId(v string)`

SetSharedDnnConfigurationsId sets SharedDnnConfigurationsId field to given value.

### HasSharedDnnConfigurationsId

`func (o *SessionManagementSubscriptionData) HasSharedDnnConfigurationsId() bool`

HasSharedDnnConfigurationsId returns a boolean if a field has been set.

### GetOdbPacketServices

`func (o *SessionManagementSubscriptionData) GetOdbPacketServices() OdbPacketServices`

GetOdbPacketServices returns the OdbPacketServices field if non-nil, zero value otherwise.

### GetOdbPacketServicesOk

`func (o *SessionManagementSubscriptionData) GetOdbPacketServicesOk() (*OdbPacketServices, bool)`

GetOdbPacketServicesOk returns a tuple with the OdbPacketServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdbPacketServices

`func (o *SessionManagementSubscriptionData) SetOdbPacketServices(v OdbPacketServices)`

SetOdbPacketServices sets OdbPacketServices field to given value.

### HasOdbPacketServices

`func (o *SessionManagementSubscriptionData) HasOdbPacketServices() bool`

HasOdbPacketServices returns a boolean if a field has been set.

### SetOdbPacketServicesNil

`func (o *SessionManagementSubscriptionData) SetOdbPacketServicesNil(b bool)`

 SetOdbPacketServicesNil sets the value for OdbPacketServices to be an explicit nil

### UnsetOdbPacketServices
`func (o *SessionManagementSubscriptionData) UnsetOdbPacketServices()`

UnsetOdbPacketServices ensures that no value is present for OdbPacketServices, not even an explicit nil
### GetTraceData

`func (o *SessionManagementSubscriptionData) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SessionManagementSubscriptionData) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SessionManagementSubscriptionData) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SessionManagementSubscriptionData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SessionManagementSubscriptionData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SessionManagementSubscriptionData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetSharedTraceDataId

`func (o *SessionManagementSubscriptionData) GetSharedTraceDataId() string`

GetSharedTraceDataId returns the SharedTraceDataId field if non-nil, zero value otherwise.

### GetSharedTraceDataIdOk

`func (o *SessionManagementSubscriptionData) GetSharedTraceDataIdOk() (*string, bool)`

GetSharedTraceDataIdOk returns a tuple with the SharedTraceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTraceDataId

`func (o *SessionManagementSubscriptionData) SetSharedTraceDataId(v string)`

SetSharedTraceDataId sets SharedTraceDataId field to given value.

### HasSharedTraceDataId

`func (o *SessionManagementSubscriptionData) HasSharedTraceDataId() bool`

HasSharedTraceDataId returns a boolean if a field has been set.

### GetExpectedUeBehavioursList

`func (o *SessionManagementSubscriptionData) GetExpectedUeBehavioursList() map[string]ExpectedUeBehaviourData`

GetExpectedUeBehavioursList returns the ExpectedUeBehavioursList field if non-nil, zero value otherwise.

### GetExpectedUeBehavioursListOk

`func (o *SessionManagementSubscriptionData) GetExpectedUeBehavioursListOk() (*map[string]ExpectedUeBehaviourData, bool)`

GetExpectedUeBehavioursListOk returns a tuple with the ExpectedUeBehavioursList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehavioursList

`func (o *SessionManagementSubscriptionData) SetExpectedUeBehavioursList(v map[string]ExpectedUeBehaviourData)`

SetExpectedUeBehavioursList sets ExpectedUeBehavioursList field to given value.

### HasExpectedUeBehavioursList

`func (o *SessionManagementSubscriptionData) HasExpectedUeBehavioursList() bool`

HasExpectedUeBehavioursList returns a boolean if a field has been set.

### GetExpectedUeBehaviourData

`func (o *SessionManagementSubscriptionData) GetExpectedUeBehaviourData() map[string]map[string]ExpectedUeBehaviourData`

GetExpectedUeBehaviourData returns the ExpectedUeBehaviourData field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourDataOk

`func (o *SessionManagementSubscriptionData) GetExpectedUeBehaviourDataOk() (*map[string]map[string]ExpectedUeBehaviourData, bool)`

GetExpectedUeBehaviourDataOk returns a tuple with the ExpectedUeBehaviourData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourData

`func (o *SessionManagementSubscriptionData) SetExpectedUeBehaviourData(v map[string]map[string]ExpectedUeBehaviourData)`

SetExpectedUeBehaviourData sets ExpectedUeBehaviourData field to given value.

### HasExpectedUeBehaviourData

`func (o *SessionManagementSubscriptionData) HasExpectedUeBehaviourData() bool`

HasExpectedUeBehaviourData returns a boolean if a field has been set.

### GetAppSpecificExpectedUeBehaviourData

`func (o *SessionManagementSubscriptionData) GetAppSpecificExpectedUeBehaviourData() map[string]map[string]AppSpecificExpectedUeBehaviourData`

GetAppSpecificExpectedUeBehaviourData returns the AppSpecificExpectedUeBehaviourData field if non-nil, zero value otherwise.

### GetAppSpecificExpectedUeBehaviourDataOk

`func (o *SessionManagementSubscriptionData) GetAppSpecificExpectedUeBehaviourDataOk() (*map[string]map[string]AppSpecificExpectedUeBehaviourData, bool)`

GetAppSpecificExpectedUeBehaviourDataOk returns a tuple with the AppSpecificExpectedUeBehaviourData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpecificExpectedUeBehaviourData

`func (o *SessionManagementSubscriptionData) SetAppSpecificExpectedUeBehaviourData(v map[string]map[string]AppSpecificExpectedUeBehaviourData)`

SetAppSpecificExpectedUeBehaviourData sets AppSpecificExpectedUeBehaviourData field to given value.

### HasAppSpecificExpectedUeBehaviourData

`func (o *SessionManagementSubscriptionData) HasAppSpecificExpectedUeBehaviourData() bool`

HasAppSpecificExpectedUeBehaviourData returns a boolean if a field has been set.

### GetSuggestedPacketNumDlList

`func (o *SessionManagementSubscriptionData) GetSuggestedPacketNumDlList() map[string]SuggestedPacketNumDl`

GetSuggestedPacketNumDlList returns the SuggestedPacketNumDlList field if non-nil, zero value otherwise.

### GetSuggestedPacketNumDlListOk

`func (o *SessionManagementSubscriptionData) GetSuggestedPacketNumDlListOk() (*map[string]SuggestedPacketNumDl, bool)`

GetSuggestedPacketNumDlListOk returns a tuple with the SuggestedPacketNumDlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPacketNumDlList

`func (o *SessionManagementSubscriptionData) SetSuggestedPacketNumDlList(v map[string]SuggestedPacketNumDl)`

SetSuggestedPacketNumDlList sets SuggestedPacketNumDlList field to given value.

### HasSuggestedPacketNumDlList

`func (o *SessionManagementSubscriptionData) HasSuggestedPacketNumDlList() bool`

HasSuggestedPacketNumDlList returns a boolean if a field has been set.

### GetVar3gppChargingCharacteristics

`func (o *SessionManagementSubscriptionData) GetVar3gppChargingCharacteristics() string`

GetVar3gppChargingCharacteristics returns the Var3gppChargingCharacteristics field if non-nil, zero value otherwise.

### GetVar3gppChargingCharacteristicsOk

`func (o *SessionManagementSubscriptionData) GetVar3gppChargingCharacteristicsOk() (*string, bool)`

GetVar3gppChargingCharacteristicsOk returns a tuple with the Var3gppChargingCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppChargingCharacteristics

`func (o *SessionManagementSubscriptionData) SetVar3gppChargingCharacteristics(v string)`

SetVar3gppChargingCharacteristics sets Var3gppChargingCharacteristics field to given value.

### HasVar3gppChargingCharacteristics

`func (o *SessionManagementSubscriptionData) HasVar3gppChargingCharacteristics() bool`

HasVar3gppChargingCharacteristics returns a boolean if a field has been set.

### GetNsacMode

`func (o *SessionManagementSubscriptionData) GetNsacMode() NsacAdmissionMode`

GetNsacMode returns the NsacMode field if non-nil, zero value otherwise.

### GetNsacModeOk

`func (o *SessionManagementSubscriptionData) GetNsacModeOk() (*NsacAdmissionMode, bool)`

GetNsacModeOk returns a tuple with the NsacMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacMode

`func (o *SessionManagementSubscriptionData) SetNsacMode(v NsacAdmissionMode)`

SetNsacMode sets NsacMode field to given value.

### HasNsacMode

`func (o *SessionManagementSubscriptionData) HasNsacMode() bool`

HasNsacMode returns a boolean if a field has been set.

### GetSessInactTimer

`func (o *SessionManagementSubscriptionData) GetSessInactTimer() int32`

GetSessInactTimer returns the SessInactTimer field if non-nil, zero value otherwise.

### GetSessInactTimerOk

`func (o *SessionManagementSubscriptionData) GetSessInactTimerOk() (*int32, bool)`

GetSessInactTimerOk returns a tuple with the SessInactTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessInactTimer

`func (o *SessionManagementSubscriptionData) SetSessInactTimer(v int32)`

SetSessInactTimer sets SessInactTimer field to given value.

### HasSessInactTimer

`func (o *SessionManagementSubscriptionData) HasSessInactTimer() bool`

HasSessInactTimer returns a boolean if a field has been set.

### GetOnDemand

`func (o *SessionManagementSubscriptionData) GetOnDemand() bool`

GetOnDemand returns the OnDemand field if non-nil, zero value otherwise.

### GetOnDemandOk

`func (o *SessionManagementSubscriptionData) GetOnDemandOk() (*bool, bool)`

GetOnDemandOk returns a tuple with the OnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemand

`func (o *SessionManagementSubscriptionData) SetOnDemand(v bool)`

SetOnDemand sets OnDemand field to given value.

### HasOnDemand

`func (o *SessionManagementSubscriptionData) HasOnDemand() bool`

HasOnDemand returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SessionManagementSubscriptionData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SessionManagementSubscriptionData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SessionManagementSubscriptionData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SessionManagementSubscriptionData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAdditionalSharedDnnConfigurationsIds

`func (o *SessionManagementSubscriptionData) GetAdditionalSharedDnnConfigurationsIds() []string`

GetAdditionalSharedDnnConfigurationsIds returns the AdditionalSharedDnnConfigurationsIds field if non-nil, zero value otherwise.

### GetAdditionalSharedDnnConfigurationsIdsOk

`func (o *SessionManagementSubscriptionData) GetAdditionalSharedDnnConfigurationsIdsOk() (*[]string, bool)`

GetAdditionalSharedDnnConfigurationsIdsOk returns a tuple with the AdditionalSharedDnnConfigurationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSharedDnnConfigurationsIds

`func (o *SessionManagementSubscriptionData) SetAdditionalSharedDnnConfigurationsIds(v []string)`

SetAdditionalSharedDnnConfigurationsIds sets AdditionalSharedDnnConfigurationsIds field to given value.

### HasAdditionalSharedDnnConfigurationsIds

`func (o *SessionManagementSubscriptionData) HasAdditionalSharedDnnConfigurationsIds() bool`

HasAdditionalSharedDnnConfigurationsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


