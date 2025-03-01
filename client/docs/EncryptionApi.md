# \EncryptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEncryptionPassword**](EncryptionApi.md#CreateEncryptionPassword) | **Post** /api/v1/encryptionPasswords | Add Encryption Password
[**DeleteEncryptionPassword**](EncryptionApi.md#DeleteEncryptionPassword) | **Delete** /api/v1/encryptionPasswords/{id} | Remove Encryption Password
[**GetAllEncryptionPasswords**](EncryptionApi.md#GetAllEncryptionPasswords) | **Get** /api/v1/encryptionPasswords | Get All Encryption Passwords
[**GetEncryptionPassword**](EncryptionApi.md#GetEncryptionPassword) | **Get** /api/v1/encryptionPasswords/{id} | Get Encryption Password
[**UpdateEncryptionPassword**](EncryptionApi.md#UpdateEncryptionPassword) | **Put** /api/v1/encryptionPasswords/{id} | Edit Encryption Password



## CreateEncryptionPassword

> EncryptionPasswordModel CreateEncryptionPassword(ctx).XApiVersion(xApiVersion).EncryptionPasswordSpec(encryptionPasswordSpec).Execute()

Add Encryption Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    encryptionPasswordSpec := *client.NewEncryptionPasswordSpec("Password_example", "Hint_example") // EncryptionPasswordSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.EncryptionApi.CreateEncryptionPassword(context.Background()).XApiVersion(xApiVersion).EncryptionPasswordSpec(encryptionPasswordSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.CreateEncryptionPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEncryptionPassword`: EncryptionPasswordModel
    fmt.Fprintf(os.Stdout, "Response from `EncryptionApi.CreateEncryptionPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEncryptionPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **encryptionPasswordSpec** | [**EncryptionPasswordSpec**](EncryptionPasswordSpec.md) |  | 

### Return type

[**EncryptionPasswordModel**](EncryptionPasswordModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEncryptionPassword

> map[string]interface{} DeleteEncryptionPassword(ctx, id).XApiVersion(xApiVersion).Execute()

Remove Encryption Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the encryption password.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.EncryptionApi.DeleteEncryptionPassword(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.DeleteEncryptionPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEncryptionPassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EncryptionApi.DeleteEncryptionPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the encryption password. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEncryptionPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllEncryptionPasswords

> EncryptionPasswordsResult GetAllEncryptionPasswords(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).Execute()

Get All Encryption Passwords



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    skip := int32(56) // int32 | Number of passwords to skip. (optional)
    limit := int32(56) // int32 | Maximum number of passwords to return. (optional)
    orderColumn := client.EEncryptionPasswordsFiltersOrderColumn("Hint") // EEncryptionPasswordsFiltersOrderColumn | Sorts passwords by one of the password parameters. (optional)
    orderAsc := true // bool | Sorts passwords in the ascending order by the `orderColumn` parameter. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.EncryptionApi.GetAllEncryptionPasswords(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.GetAllEncryptionPasswords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllEncryptionPasswords`: EncryptionPasswordsResult
    fmt.Fprintf(os.Stdout, "Response from `EncryptionApi.GetAllEncryptionPasswords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllEncryptionPasswordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of passwords to skip. | 
 **limit** | **int32** | Maximum number of passwords to return. | 
 **orderColumn** | [**EEncryptionPasswordsFiltersOrderColumn**](EEncryptionPasswordsFiltersOrderColumn.md) | Sorts passwords by one of the password parameters. | 
 **orderAsc** | **bool** | Sorts passwords in the ascending order by the &#x60;orderColumn&#x60; parameter. | 

### Return type

[**EncryptionPasswordsResult**](EncryptionPasswordsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionPassword

> EncryptionPasswordModel GetEncryptionPassword(ctx, id).XApiVersion(xApiVersion).Execute()

Get Encryption Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the encryption password.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.EncryptionApi.GetEncryptionPassword(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.GetEncryptionPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEncryptionPassword`: EncryptionPasswordModel
    fmt.Fprintf(os.Stdout, "Response from `EncryptionApi.GetEncryptionPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the encryption password. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**EncryptionPasswordModel**](EncryptionPasswordModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEncryptionPassword

> EncryptionPasswordModel UpdateEncryptionPassword(ctx, id).XApiVersion(xApiVersion).EncryptionPasswordModel(encryptionPasswordModel).Execute()

Edit Encryption Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the encryption password.
    encryptionPasswordModel := *client.NewEncryptionPasswordModel("Id_example", "Hint_example") // EncryptionPasswordModel |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.EncryptionApi.UpdateEncryptionPassword(context.Background(), id).XApiVersion(xApiVersion).EncryptionPasswordModel(encryptionPasswordModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.UpdateEncryptionPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEncryptionPassword`: EncryptionPasswordModel
    fmt.Fprintf(os.Stdout, "Response from `EncryptionApi.UpdateEncryptionPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the encryption password. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEncryptionPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **encryptionPasswordModel** | [**EncryptionPasswordModel**](EncryptionPasswordModel.md) |  | 

### Return type

[**EncryptionPasswordModel**](EncryptionPasswordModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

