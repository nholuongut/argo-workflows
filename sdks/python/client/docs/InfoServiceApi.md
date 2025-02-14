# argo_workflows.InfoServiceApi

All URIs are relative to *http://localhost:2746*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_info**](InfoServiceApi.md#get_info) | **GET** /api/v1/info | 
[**get_user_info**](InfoServiceApi.md#get_user_info) | **GET** /api/v1/userinfo | 
[**get_version**](InfoServiceApi.md#get_version) | **GET** /api/v1/version | 


# **get_info**
> IonholuongutWorkflowV1alpha1InfoResponse get_info()



### Example

```python
import time
import argo_workflows
from argo_workflows.api import info_service_api
from argo_workflows.model.grpc_gateway_runtime_error import GrpcGatewayRuntimeError
from argo_workflows.model.io_nholuongut_workflow_v1alpha1_info_response import IonholuongutWorkflowV1alpha1InfoResponse
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:2746
# See configuration.py for a list of all supported configuration parameters.
configuration = argo_workflows.Configuration(
    host = "http://localhost:2746"
)


# Enter a context with an instance of the API client
with argo_workflows.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = info_service_api.InfoServiceApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.get_info()
        pprint(api_response)
    except argo_workflows.ApiException as e:
        print("Exception when calling InfoServiceApi->get_info: %s\n" % e)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user_info**
> IonholuongutWorkflowV1alpha1GetUserInfoResponse get_user_info()



### Example

```python
import time
import argo_workflows
from argo_workflows.api import info_service_api
from argo_workflows.model.grpc_gateway_runtime_error import GrpcGatewayRuntimeError
from argo_workflows.model.io_nholuongut_workflow_v1alpha1_get_user_info_response import IonholuongutWorkflowV1alpha1GetUserInfoResponse
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:2746
# See configuration.py for a list of all supported configuration parameters.
configuration = argo_workflows.Configuration(
    host = "http://localhost:2746"
)


# Enter a context with an instance of the API client
with argo_workflows.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = info_service_api.InfoServiceApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.get_user_info()
        pprint(api_response)
    except argo_workflows.ApiException as e:
        print("Exception when calling InfoServiceApi->get_user_info: %s\n" % e)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_version**
> IonholuongutWorkflowV1alpha1Version get_version()



### Example

```python
import time
import argo_workflows
from argo_workflows.api import info_service_api
from argo_workflows.model.grpc_gateway_runtime_error import GrpcGatewayRuntimeError
from argo_workflows.model.io_nholuongut_workflow_v1alpha1_version import IonholuongutWorkflowV1alpha1Version
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:2746
# See configuration.py for a list of all supported configuration parameters.
configuration = argo_workflows.Configuration(
    host = "http://localhost:2746"
)


# Enter a context with an instance of the API client
with argo_workflows.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = info_service_api.InfoServiceApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.get_version()
        pprint(api_response)
    except argo_workflows.ApiException as e:
        print("Exception when calling InfoServiceApi->get_version: %s\n" % e)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

