# InfoServiceApi

All URIs are relative to *http://localhost:2746*

Method | HTTP request | Description
------------- | ------------- | -------------
[**infoServiceGetInfo**](InfoServiceApi.md#infoServiceGetInfo) | **GET** /api/v1/info | 
[**infoServiceGetUserInfo**](InfoServiceApi.md#infoServiceGetUserInfo) | **GET** /api/v1/userinfo | 
[**infoServiceGetVersion**](InfoServiceApi.md#infoServiceGetVersion) | **GET** /api/v1/version | 


<a name="infoServiceGetInfo"></a>
# **infoServiceGetInfo**
> IonholuongutWorkflowV1alpha1InfoResponse infoServiceGetInfo()



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.InfoServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    InfoServiceApi apiInstance = new InfoServiceApi(defaultClient);
    try {
      IonholuongutWorkflowV1alpha1InfoResponse result = apiInstance.infoServiceGetInfo();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InfoServiceApi#infoServiceGetInfo");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IonholuongutWorkflowV1alpha1InfoResponse**](IonholuongutWorkflowV1alpha1InfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

<a name="infoServiceGetUserInfo"></a>
# **infoServiceGetUserInfo**
> IonholuongutWorkflowV1alpha1GetUserInfoResponse infoServiceGetUserInfo()



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.InfoServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    InfoServiceApi apiInstance = new InfoServiceApi(defaultClient);
    try {
      IonholuongutWorkflowV1alpha1GetUserInfoResponse result = apiInstance.infoServiceGetUserInfo();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InfoServiceApi#infoServiceGetUserInfo");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IonholuongutWorkflowV1alpha1GetUserInfoResponse**](IonholuongutWorkflowV1alpha1GetUserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

<a name="infoServiceGetVersion"></a>
# **infoServiceGetVersion**
> IonholuongutWorkflowV1alpha1Version infoServiceGetVersion()



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.InfoServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    InfoServiceApi apiInstance = new InfoServiceApi(defaultClient);
    try {
      IonholuongutWorkflowV1alpha1Version result = apiInstance.infoServiceGetVersion();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InfoServiceApi#infoServiceGetVersion");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IonholuongutWorkflowV1alpha1Version**](IonholuongutWorkflowV1alpha1Version.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

