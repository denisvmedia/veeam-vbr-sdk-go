# AmazonSnowballEdgeBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server used to connect to the AWS Snowball Edge device. | [optional] 
**ConnectionPoint** | Pointer to **string** | Service point address and port number of the AWS Snowball Edge device. | [optional] 
**Regions** | Pointer to [**[]AmazonSnowballEdgeRegionBrowserModel**](AmazonSnowballEdgeRegionBrowserModel.md) | Array of regions. | [optional] 

## Methods

### NewAmazonSnowballEdgeBrowserModel

`func NewAmazonSnowballEdgeBrowserModel() *AmazonSnowballEdgeBrowserModel`

NewAmazonSnowballEdgeBrowserModel instantiates a new AmazonSnowballEdgeBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeBrowserModelWithDefaults

`func NewAmazonSnowballEdgeBrowserModelWithDefaults() *AmazonSnowballEdgeBrowserModel`

NewAmazonSnowballEdgeBrowserModelWithDefaults instantiates a new AmazonSnowballEdgeBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AmazonSnowballEdgeBrowserModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AmazonSnowballEdgeBrowserModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AmazonSnowballEdgeBrowserModel) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AmazonSnowballEdgeBrowserModel) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetConnectionPoint

`func (o *AmazonSnowballEdgeBrowserModel) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *AmazonSnowballEdgeBrowserModel) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *AmazonSnowballEdgeBrowserModel) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.

### HasConnectionPoint

`func (o *AmazonSnowballEdgeBrowserModel) HasConnectionPoint() bool`

HasConnectionPoint returns a boolean if a field has been set.

### GetRegions

`func (o *AmazonSnowballEdgeBrowserModel) GetRegions() []AmazonSnowballEdgeRegionBrowserModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AmazonSnowballEdgeBrowserModel) GetRegionsOk() (*[]AmazonSnowballEdgeRegionBrowserModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AmazonSnowballEdgeBrowserModel) SetRegions(v []AmazonSnowballEdgeRegionBrowserModel)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AmazonSnowballEdgeBrowserModel) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


