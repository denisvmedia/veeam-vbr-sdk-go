# VCenterInventoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VmwareObjectSizeModel**](VmwareObjectSizeModel.md) | Array of vCenter Servers. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewVCenterInventoryResult

`func NewVCenterInventoryResult(data []VmwareObjectSizeModel, pagination PaginationResult, ) *VCenterInventoryResult`

NewVCenterInventoryResult instantiates a new VCenterInventoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCenterInventoryResultWithDefaults

`func NewVCenterInventoryResultWithDefaults() *VCenterInventoryResult`

NewVCenterInventoryResultWithDefaults instantiates a new VCenterInventoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VCenterInventoryResult) GetData() []VmwareObjectSizeModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VCenterInventoryResult) GetDataOk() (*[]VmwareObjectSizeModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VCenterInventoryResult) SetData(v []VmwareObjectSizeModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *VCenterInventoryResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VCenterInventoryResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VCenterInventoryResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


