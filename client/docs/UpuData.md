# UpuData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecPacket** | Pointer to **string** | Contains a secure packet. | [optional] 
**DefaultConfNssai** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RoutingId** | Pointer to **string** | Represents a routing indicator. | [optional] 

## Methods

### NewUpuData

`func NewUpuData() *UpuData`

NewUpuData instantiates a new UpuData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpuDataWithDefaults

`func NewUpuDataWithDefaults() *UpuData`

NewUpuDataWithDefaults instantiates a new UpuData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecPacket

`func (o *UpuData) GetSecPacket() string`

GetSecPacket returns the SecPacket field if non-nil, zero value otherwise.

### GetSecPacketOk

`func (o *UpuData) GetSecPacketOk() (*string, bool)`

GetSecPacketOk returns a tuple with the SecPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecPacket

`func (o *UpuData) SetSecPacket(v string)`

SetSecPacket sets SecPacket field to given value.

### HasSecPacket

`func (o *UpuData) HasSecPacket() bool`

HasSecPacket returns a boolean if a field has been set.

### GetDefaultConfNssai

`func (o *UpuData) GetDefaultConfNssai() []Snssai`

GetDefaultConfNssai returns the DefaultConfNssai field if non-nil, zero value otherwise.

### GetDefaultConfNssaiOk

`func (o *UpuData) GetDefaultConfNssaiOk() (*[]Snssai, bool)`

GetDefaultConfNssaiOk returns a tuple with the DefaultConfNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfNssai

`func (o *UpuData) SetDefaultConfNssai(v []Snssai)`

SetDefaultConfNssai sets DefaultConfNssai field to given value.

### HasDefaultConfNssai

`func (o *UpuData) HasDefaultConfNssai() bool`

HasDefaultConfNssai returns a boolean if a field has been set.

### GetRoutingId

`func (o *UpuData) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### GetRoutingIdOk

`func (o *UpuData) GetRoutingIdOk() (*string, bool)`

GetRoutingIdOk returns a tuple with the RoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingId

`func (o *UpuData) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### HasRoutingId

`func (o *UpuData) HasRoutingId() bool`

HasRoutingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


