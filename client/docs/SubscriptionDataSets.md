# SubscriptionDataSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmData** | Pointer to [**AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md) |  | [optional] 
**SmfSelData** | Pointer to [**SmfSelectionSubscriptionData**](SmfSelectionSubscriptionData.md) |  | [optional] 
**UecAmfData** | Pointer to [**UeContextInAmfData**](UeContextInAmfData.md) |  | [optional] 
**UecSmfData** | Pointer to [**UeContextInSmfData**](UeContextInSmfData.md) |  | [optional] 
**UecSmsfData** | Pointer to [**UeContextInSmsfData**](UeContextInSmsfData.md) |  | [optional] 
**SmsSubsData** | Pointer to [**SmsSubscriptionData**](SmsSubscriptionData.md) |  | [optional] 
**SmData** | Pointer to [**SmSubsData**](SmSubsData.md) |  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**SmsMngData** | Pointer to [**SmsManagementSubscriptionData**](SmsManagementSubscriptionData.md) |  | [optional] 
**LcsPrivacyData** | Pointer to [**LcsPrivacyData**](LcsPrivacyData.md) |  | [optional] 
**LcsMoData** | Pointer to [**LcsMoData**](LcsMoData.md) |  | [optional] 
**LcsSubscriptionData** | Pointer to [**LcsSubscriptionData**](LcsSubscriptionData.md) |  | [optional] 
**V2xData** | Pointer to [**V2xSubscriptionData**](V2xSubscriptionData.md) |  | [optional] 
**LcsBroadcastAssistanceTypesData** | Pointer to [**LcsBroadcastAssistanceTypesData**](LcsBroadcastAssistanceTypesData.md) |  | [optional] 
**ProseData** | Pointer to [**ProseSubscriptionData**](ProseSubscriptionData.md) |  | [optional] 
**MbsData** | Pointer to [**MbsSubscriptionData**](MbsSubscriptionData.md) |  | [optional] 
**UcData** | Pointer to [**UcSubscriptionData**](UcSubscriptionData.md) |  | [optional] 
**A2xData** | Pointer to [**A2xSubscriptionData**](A2xSubscriptionData.md) |  | [optional] 

## Methods

### NewSubscriptionDataSets

`func NewSubscriptionDataSets() *SubscriptionDataSets`

NewSubscriptionDataSets instantiates a new SubscriptionDataSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataSetsWithDefaults

`func NewSubscriptionDataSetsWithDefaults() *SubscriptionDataSets`

NewSubscriptionDataSetsWithDefaults instantiates a new SubscriptionDataSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmData

`func (o *SubscriptionDataSets) GetAmData() AccessAndMobilitySubscriptionData`

GetAmData returns the AmData field if non-nil, zero value otherwise.

### GetAmDataOk

`func (o *SubscriptionDataSets) GetAmDataOk() (*AccessAndMobilitySubscriptionData, bool)`

GetAmDataOk returns a tuple with the AmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmData

`func (o *SubscriptionDataSets) SetAmData(v AccessAndMobilitySubscriptionData)`

SetAmData sets AmData field to given value.

### HasAmData

`func (o *SubscriptionDataSets) HasAmData() bool`

HasAmData returns a boolean if a field has been set.

### GetSmfSelData

`func (o *SubscriptionDataSets) GetSmfSelData() SmfSelectionSubscriptionData`

GetSmfSelData returns the SmfSelData field if non-nil, zero value otherwise.

### GetSmfSelDataOk

`func (o *SubscriptionDataSets) GetSmfSelDataOk() (*SmfSelectionSubscriptionData, bool)`

GetSmfSelDataOk returns a tuple with the SmfSelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelData

`func (o *SubscriptionDataSets) SetSmfSelData(v SmfSelectionSubscriptionData)`

SetSmfSelData sets SmfSelData field to given value.

### HasSmfSelData

`func (o *SubscriptionDataSets) HasSmfSelData() bool`

HasSmfSelData returns a boolean if a field has been set.

### GetUecAmfData

`func (o *SubscriptionDataSets) GetUecAmfData() UeContextInAmfData`

GetUecAmfData returns the UecAmfData field if non-nil, zero value otherwise.

### GetUecAmfDataOk

`func (o *SubscriptionDataSets) GetUecAmfDataOk() (*UeContextInAmfData, bool)`

GetUecAmfDataOk returns a tuple with the UecAmfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUecAmfData

`func (o *SubscriptionDataSets) SetUecAmfData(v UeContextInAmfData)`

SetUecAmfData sets UecAmfData field to given value.

### HasUecAmfData

`func (o *SubscriptionDataSets) HasUecAmfData() bool`

HasUecAmfData returns a boolean if a field has been set.

### GetUecSmfData

`func (o *SubscriptionDataSets) GetUecSmfData() UeContextInSmfData`

GetUecSmfData returns the UecSmfData field if non-nil, zero value otherwise.

### GetUecSmfDataOk

`func (o *SubscriptionDataSets) GetUecSmfDataOk() (*UeContextInSmfData, bool)`

GetUecSmfDataOk returns a tuple with the UecSmfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUecSmfData

`func (o *SubscriptionDataSets) SetUecSmfData(v UeContextInSmfData)`

SetUecSmfData sets UecSmfData field to given value.

### HasUecSmfData

`func (o *SubscriptionDataSets) HasUecSmfData() bool`

HasUecSmfData returns a boolean if a field has been set.

### GetUecSmsfData

`func (o *SubscriptionDataSets) GetUecSmsfData() UeContextInSmsfData`

GetUecSmsfData returns the UecSmsfData field if non-nil, zero value otherwise.

### GetUecSmsfDataOk

`func (o *SubscriptionDataSets) GetUecSmsfDataOk() (*UeContextInSmsfData, bool)`

GetUecSmsfDataOk returns a tuple with the UecSmsfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUecSmsfData

`func (o *SubscriptionDataSets) SetUecSmsfData(v UeContextInSmsfData)`

SetUecSmsfData sets UecSmsfData field to given value.

### HasUecSmsfData

`func (o *SubscriptionDataSets) HasUecSmsfData() bool`

HasUecSmsfData returns a boolean if a field has been set.

### GetSmsSubsData

`func (o *SubscriptionDataSets) GetSmsSubsData() SmsSubscriptionData`

GetSmsSubsData returns the SmsSubsData field if non-nil, zero value otherwise.

### GetSmsSubsDataOk

`func (o *SubscriptionDataSets) GetSmsSubsDataOk() (*SmsSubscriptionData, bool)`

GetSmsSubsDataOk returns a tuple with the SmsSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSubsData

`func (o *SubscriptionDataSets) SetSmsSubsData(v SmsSubscriptionData)`

SetSmsSubsData sets SmsSubsData field to given value.

### HasSmsSubsData

`func (o *SubscriptionDataSets) HasSmsSubsData() bool`

HasSmsSubsData returns a boolean if a field has been set.

### GetSmData

`func (o *SubscriptionDataSets) GetSmData() SmSubsData`

GetSmData returns the SmData field if non-nil, zero value otherwise.

### GetSmDataOk

`func (o *SubscriptionDataSets) GetSmDataOk() (*SmSubsData, bool)`

GetSmDataOk returns a tuple with the SmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmData

`func (o *SubscriptionDataSets) SetSmData(v SmSubsData)`

SetSmData sets SmData field to given value.

### HasSmData

`func (o *SubscriptionDataSets) HasSmData() bool`

HasSmData returns a boolean if a field has been set.

### GetTraceData

`func (o *SubscriptionDataSets) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SubscriptionDataSets) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SubscriptionDataSets) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SubscriptionDataSets) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SubscriptionDataSets) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SubscriptionDataSets) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetSmsMngData

`func (o *SubscriptionDataSets) GetSmsMngData() SmsManagementSubscriptionData`

GetSmsMngData returns the SmsMngData field if non-nil, zero value otherwise.

### GetSmsMngDataOk

`func (o *SubscriptionDataSets) GetSmsMngDataOk() (*SmsManagementSubscriptionData, bool)`

GetSmsMngDataOk returns a tuple with the SmsMngData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMngData

`func (o *SubscriptionDataSets) SetSmsMngData(v SmsManagementSubscriptionData)`

SetSmsMngData sets SmsMngData field to given value.

### HasSmsMngData

`func (o *SubscriptionDataSets) HasSmsMngData() bool`

HasSmsMngData returns a boolean if a field has been set.

### GetLcsPrivacyData

`func (o *SubscriptionDataSets) GetLcsPrivacyData() LcsPrivacyData`

GetLcsPrivacyData returns the LcsPrivacyData field if non-nil, zero value otherwise.

### GetLcsPrivacyDataOk

`func (o *SubscriptionDataSets) GetLcsPrivacyDataOk() (*LcsPrivacyData, bool)`

GetLcsPrivacyDataOk returns a tuple with the LcsPrivacyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsPrivacyData

`func (o *SubscriptionDataSets) SetLcsPrivacyData(v LcsPrivacyData)`

SetLcsPrivacyData sets LcsPrivacyData field to given value.

### HasLcsPrivacyData

`func (o *SubscriptionDataSets) HasLcsPrivacyData() bool`

HasLcsPrivacyData returns a boolean if a field has been set.

### GetLcsMoData

`func (o *SubscriptionDataSets) GetLcsMoData() LcsMoData`

GetLcsMoData returns the LcsMoData field if non-nil, zero value otherwise.

### GetLcsMoDataOk

`func (o *SubscriptionDataSets) GetLcsMoDataOk() (*LcsMoData, bool)`

GetLcsMoDataOk returns a tuple with the LcsMoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsMoData

`func (o *SubscriptionDataSets) SetLcsMoData(v LcsMoData)`

SetLcsMoData sets LcsMoData field to given value.

### HasLcsMoData

`func (o *SubscriptionDataSets) HasLcsMoData() bool`

HasLcsMoData returns a boolean if a field has been set.

### GetLcsSubscriptionData

`func (o *SubscriptionDataSets) GetLcsSubscriptionData() LcsSubscriptionData`

GetLcsSubscriptionData returns the LcsSubscriptionData field if non-nil, zero value otherwise.

### GetLcsSubscriptionDataOk

`func (o *SubscriptionDataSets) GetLcsSubscriptionDataOk() (*LcsSubscriptionData, bool)`

GetLcsSubscriptionDataOk returns a tuple with the LcsSubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsSubscriptionData

`func (o *SubscriptionDataSets) SetLcsSubscriptionData(v LcsSubscriptionData)`

SetLcsSubscriptionData sets LcsSubscriptionData field to given value.

### HasLcsSubscriptionData

`func (o *SubscriptionDataSets) HasLcsSubscriptionData() bool`

HasLcsSubscriptionData returns a boolean if a field has been set.

### GetV2xData

`func (o *SubscriptionDataSets) GetV2xData() V2xSubscriptionData`

GetV2xData returns the V2xData field if non-nil, zero value otherwise.

### GetV2xDataOk

`func (o *SubscriptionDataSets) GetV2xDataOk() (*V2xSubscriptionData, bool)`

GetV2xDataOk returns a tuple with the V2xData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xData

`func (o *SubscriptionDataSets) SetV2xData(v V2xSubscriptionData)`

SetV2xData sets V2xData field to given value.

### HasV2xData

`func (o *SubscriptionDataSets) HasV2xData() bool`

HasV2xData returns a boolean if a field has been set.

### GetLcsBroadcastAssistanceTypesData

`func (o *SubscriptionDataSets) GetLcsBroadcastAssistanceTypesData() LcsBroadcastAssistanceTypesData`

GetLcsBroadcastAssistanceTypesData returns the LcsBroadcastAssistanceTypesData field if non-nil, zero value otherwise.

### GetLcsBroadcastAssistanceTypesDataOk

`func (o *SubscriptionDataSets) GetLcsBroadcastAssistanceTypesDataOk() (*LcsBroadcastAssistanceTypesData, bool)`

GetLcsBroadcastAssistanceTypesDataOk returns a tuple with the LcsBroadcastAssistanceTypesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsBroadcastAssistanceTypesData

`func (o *SubscriptionDataSets) SetLcsBroadcastAssistanceTypesData(v LcsBroadcastAssistanceTypesData)`

SetLcsBroadcastAssistanceTypesData sets LcsBroadcastAssistanceTypesData field to given value.

### HasLcsBroadcastAssistanceTypesData

`func (o *SubscriptionDataSets) HasLcsBroadcastAssistanceTypesData() bool`

HasLcsBroadcastAssistanceTypesData returns a boolean if a field has been set.

### GetProseData

`func (o *SubscriptionDataSets) GetProseData() ProseSubscriptionData`

GetProseData returns the ProseData field if non-nil, zero value otherwise.

### GetProseDataOk

`func (o *SubscriptionDataSets) GetProseDataOk() (*ProseSubscriptionData, bool)`

GetProseDataOk returns a tuple with the ProseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseData

`func (o *SubscriptionDataSets) SetProseData(v ProseSubscriptionData)`

SetProseData sets ProseData field to given value.

### HasProseData

`func (o *SubscriptionDataSets) HasProseData() bool`

HasProseData returns a boolean if a field has been set.

### GetMbsData

`func (o *SubscriptionDataSets) GetMbsData() MbsSubscriptionData`

GetMbsData returns the MbsData field if non-nil, zero value otherwise.

### GetMbsDataOk

`func (o *SubscriptionDataSets) GetMbsDataOk() (*MbsSubscriptionData, bool)`

GetMbsDataOk returns a tuple with the MbsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsData

`func (o *SubscriptionDataSets) SetMbsData(v MbsSubscriptionData)`

SetMbsData sets MbsData field to given value.

### HasMbsData

`func (o *SubscriptionDataSets) HasMbsData() bool`

HasMbsData returns a boolean if a field has been set.

### GetUcData

`func (o *SubscriptionDataSets) GetUcData() UcSubscriptionData`

GetUcData returns the UcData field if non-nil, zero value otherwise.

### GetUcDataOk

`func (o *SubscriptionDataSets) GetUcDataOk() (*UcSubscriptionData, bool)`

GetUcDataOk returns a tuple with the UcData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcData

`func (o *SubscriptionDataSets) SetUcData(v UcSubscriptionData)`

SetUcData sets UcData field to given value.

### HasUcData

`func (o *SubscriptionDataSets) HasUcData() bool`

HasUcData returns a boolean if a field has been set.

### GetA2xData

`func (o *SubscriptionDataSets) GetA2xData() A2xSubscriptionData`

GetA2xData returns the A2xData field if non-nil, zero value otherwise.

### GetA2xDataOk

`func (o *SubscriptionDataSets) GetA2xDataOk() (*A2xSubscriptionData, bool)`

GetA2xDataOk returns a tuple with the A2xData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2xData

`func (o *SubscriptionDataSets) SetA2xData(v A2xSubscriptionData)`

SetA2xData sets A2xData field to given value.

### HasA2xData

`func (o *SubscriptionDataSets) HasA2xData() bool`

HasA2xData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


