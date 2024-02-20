# AerialUeSubscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AerialUeInd** | [**AerialUeIndication**](AerialUeIndication.md) |  | 
**Var3gppUavId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 

## Methods

### NewAerialUeSubscriptionInfo

`func NewAerialUeSubscriptionInfo(aerialUeInd AerialUeIndication, ) *AerialUeSubscriptionInfo`

NewAerialUeSubscriptionInfo instantiates a new AerialUeSubscriptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAerialUeSubscriptionInfoWithDefaults

`func NewAerialUeSubscriptionInfoWithDefaults() *AerialUeSubscriptionInfo`

NewAerialUeSubscriptionInfoWithDefaults instantiates a new AerialUeSubscriptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAerialUeInd

`func (o *AerialUeSubscriptionInfo) GetAerialUeInd() AerialUeIndication`

GetAerialUeInd returns the AerialUeInd field if non-nil, zero value otherwise.

### GetAerialUeIndOk

`func (o *AerialUeSubscriptionInfo) GetAerialUeIndOk() (*AerialUeIndication, bool)`

GetAerialUeIndOk returns a tuple with the AerialUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerialUeInd

`func (o *AerialUeSubscriptionInfo) SetAerialUeInd(v AerialUeIndication)`

SetAerialUeInd sets AerialUeInd field to given value.


### GetVar3gppUavId

`func (o *AerialUeSubscriptionInfo) GetVar3gppUavId() string`

GetVar3gppUavId returns the Var3gppUavId field if non-nil, zero value otherwise.

### GetVar3gppUavIdOk

`func (o *AerialUeSubscriptionInfo) GetVar3gppUavIdOk() (*string, bool)`

GetVar3gppUavIdOk returns a tuple with the Var3gppUavId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppUavId

`func (o *AerialUeSubscriptionInfo) SetVar3gppUavId(v string)`

SetVar3gppUavId sets Var3gppUavId field to given value.

### HasVar3gppUavId

`func (o *AerialUeSubscriptionInfo) HasVar3gppUavId() bool`

HasVar3gppUavId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


