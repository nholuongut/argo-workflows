# CronWorkflowServiceApi

All URIs are relative to *http://localhost:2746*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cronWorkflowServiceCreateCronWorkflow**](CronWorkflowServiceApi.md#cronWorkflowServiceCreateCronWorkflow) | **POST** /api/v1/cron-workflows/{namespace} | 
[**cronWorkflowServiceDeleteCronWorkflow**](CronWorkflowServiceApi.md#cronWorkflowServiceDeleteCronWorkflow) | **DELETE** /api/v1/cron-workflows/{namespace}/{name} | 
[**cronWorkflowServiceGetCronWorkflow**](CronWorkflowServiceApi.md#cronWorkflowServiceGetCronWorkflow) | **GET** /api/v1/cron-workflows/{namespace}/{name} | 
[**cronWorkflowServiceLintCronWorkflow**](CronWorkflowServiceApi.md#cronWorkflowServiceLintCronWorkflow) | **POST** /api/v1/cron-workflows/{namespace}/lint | 
[**cronWorkflowServiceListCronWorkflows**](CronWorkflowServiceApi.md#cronWorkflowServiceListCronWorkflows) | **GET** /api/v1/cron-workflows/{namespace} | 
[**cronWorkflowServiceResumeCronWorkflow**](CronWorkflowServiceApi.md#cronWorkflowServiceResumeCronWorkflow) | **PUT** /api/v1/cron-workflows/{namespace}/{name}/resume | 
[**cronWorkflowServiceSuspendCronWorkflow**](CronWorkflowServiceApi.md#cronWorkflowServiceSuspendCronWorkflow) | **PUT** /api/v1/cron-workflows/{namespace}/{name}/suspend | 
[**cronWorkflowServiceUpdateCronWorkflow**](CronWorkflowServiceApi.md#cronWorkflowServiceUpdateCronWorkflow) | **PUT** /api/v1/cron-workflows/{namespace}/{name} | 


<a name="cronWorkflowServiceCreateCronWorkflow"></a>
# **cronWorkflowServiceCreateCronWorkflow**
> IonholuongutWorkflowV1alpha1CronWorkflow cronWorkflowServiceCreateCronWorkflow(namespace, body)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    IonholuongutWorkflowV1alpha1CreateCronWorkflowRequest body = new IonholuongutWorkflowV1alpha1CreateCronWorkflowRequest(); // IonholuongutWorkflowV1alpha1CreateCronWorkflowRequest | 
    try {
      IonholuongutWorkflowV1alpha1CronWorkflow result = apiInstance.cronWorkflowServiceCreateCronWorkflow(namespace, body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceCreateCronWorkflow");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **body** | [**IonholuongutWorkflowV1alpha1CreateCronWorkflowRequest**](IonholuongutWorkflowV1alpha1CreateCronWorkflowRequest.md)|  |

### Return type

[**IonholuongutWorkflowV1alpha1CronWorkflow**](IonholuongutWorkflowV1alpha1CronWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

<a name="cronWorkflowServiceDeleteCronWorkflow"></a>
# **cronWorkflowServiceDeleteCronWorkflow**
> Object cronWorkflowServiceDeleteCronWorkflow(namespace, name, deleteOptionsGracePeriodSeconds, deleteOptionsPreconditionsUid, deleteOptionsPreconditionsResourceVersion, deleteOptionsOrphanDependents, deleteOptionsPropagationPolicy, deleteOptionsDryRun)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    String name = "name_example"; // String | 
    String deleteOptionsGracePeriodSeconds = "deleteOptionsGracePeriodSeconds_example"; // String | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. +optional.
    String deleteOptionsPreconditionsUid = "deleteOptionsPreconditionsUid_example"; // String | Specifies the target UID. +optional.
    String deleteOptionsPreconditionsResourceVersion = "deleteOptionsPreconditionsResourceVersion_example"; // String | Specifies the target ResourceVersion +optional.
    Boolean deleteOptionsOrphanDependents = true; // Boolean | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. +optional.
    String deleteOptionsPropagationPolicy = "deleteOptionsPropagationPolicy_example"; // String | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. +optional.
    List<String> deleteOptionsDryRun = Arrays.asList(); // List<String> | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed +optional.
    try {
      Object result = apiInstance.cronWorkflowServiceDeleteCronWorkflow(namespace, name, deleteOptionsGracePeriodSeconds, deleteOptionsPreconditionsUid, deleteOptionsPreconditionsResourceVersion, deleteOptionsOrphanDependents, deleteOptionsPropagationPolicy, deleteOptionsDryRun);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceDeleteCronWorkflow");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **name** | **String**|  |
 **deleteOptionsGracePeriodSeconds** | **String**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. +optional. | [optional]
 **deleteOptionsPreconditionsUid** | **String**| Specifies the target UID. +optional. | [optional]
 **deleteOptionsPreconditionsResourceVersion** | **String**| Specifies the target ResourceVersion +optional. | [optional]
 **deleteOptionsOrphanDependents** | **Boolean**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. +optional. | [optional]
 **deleteOptionsPropagationPolicy** | **String**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: &#39;Orphan&#39; - orphan the dependents; &#39;Background&#39; - allow the garbage collector to delete the dependents in the background; &#39;Foreground&#39; - a cascading policy that deletes all dependents in the foreground. +optional. | [optional]
 **deleteOptionsDryRun** | [**List&lt;String&gt;**](String.md)| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed +optional. | [optional]

### Return type

**Object**

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

<a name="cronWorkflowServiceGetCronWorkflow"></a>
# **cronWorkflowServiceGetCronWorkflow**
> IonholuongutWorkflowV1alpha1CronWorkflow cronWorkflowServiceGetCronWorkflow(namespace, name, getOptionsResourceVersion)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    String name = "name_example"; // String | 
    String getOptionsResourceVersion = "getOptionsResourceVersion_example"; // String | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional
    try {
      IonholuongutWorkflowV1alpha1CronWorkflow result = apiInstance.cronWorkflowServiceGetCronWorkflow(namespace, name, getOptionsResourceVersion);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceGetCronWorkflow");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **name** | **String**|  |
 **getOptionsResourceVersion** | **String**| resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | [optional]

### Return type

[**IonholuongutWorkflowV1alpha1CronWorkflow**](IonholuongutWorkflowV1alpha1CronWorkflow.md)

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

<a name="cronWorkflowServiceLintCronWorkflow"></a>
# **cronWorkflowServiceLintCronWorkflow**
> IonholuongutWorkflowV1alpha1CronWorkflow cronWorkflowServiceLintCronWorkflow(namespace, body)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    IonholuongutWorkflowV1alpha1LintCronWorkflowRequest body = new IonholuongutWorkflowV1alpha1LintCronWorkflowRequest(); // IonholuongutWorkflowV1alpha1LintCronWorkflowRequest | 
    try {
      IonholuongutWorkflowV1alpha1CronWorkflow result = apiInstance.cronWorkflowServiceLintCronWorkflow(namespace, body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceLintCronWorkflow");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **body** | [**IonholuongutWorkflowV1alpha1LintCronWorkflowRequest**](IonholuongutWorkflowV1alpha1LintCronWorkflowRequest.md)|  |

### Return type

[**IonholuongutWorkflowV1alpha1CronWorkflow**](IonholuongutWorkflowV1alpha1CronWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

<a name="cronWorkflowServiceListCronWorkflows"></a>
# **cronWorkflowServiceListCronWorkflows**
> IonholuongutWorkflowV1alpha1CronWorkflowList cronWorkflowServiceListCronWorkflows(namespace, listOptionsLabelSelector, listOptionsFieldSelector, listOptionsWatch, listOptionsAllowWatchBookmarks, listOptionsResourceVersion, listOptionsResourceVersionMatch, listOptionsTimeoutSeconds, listOptionsLimit, listOptionsContinue)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    String listOptionsLabelSelector = "listOptionsLabelSelector_example"; // String | A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional.
    String listOptionsFieldSelector = "listOptionsFieldSelector_example"; // String | A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional.
    Boolean listOptionsWatch = true; // Boolean | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional.
    Boolean listOptionsAllowWatchBookmarks = true; // Boolean | allowWatchBookmarks requests watch events with type \"BOOKMARK\". Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server's discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional.
    String listOptionsResourceVersion = "listOptionsResourceVersion_example"; // String | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional
    String listOptionsResourceVersionMatch = "listOptionsResourceVersionMatch_example"; // String | resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional
    String listOptionsTimeoutSeconds = "listOptionsTimeoutSeconds_example"; // String | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional.
    String listOptionsLimit = "listOptionsLimit_example"; // String | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
    String listOptionsContinue = "listOptionsContinue_example"; // String | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
    try {
      IonholuongutWorkflowV1alpha1CronWorkflowList result = apiInstance.cronWorkflowServiceListCronWorkflows(namespace, listOptionsLabelSelector, listOptionsFieldSelector, listOptionsWatch, listOptionsAllowWatchBookmarks, listOptionsResourceVersion, listOptionsResourceVersionMatch, listOptionsTimeoutSeconds, listOptionsLimit, listOptionsContinue);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceListCronWorkflows");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **listOptionsLabelSelector** | **String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional. | [optional]
 **listOptionsFieldSelector** | **String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional. | [optional]
 **listOptionsWatch** | **Boolean**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional. | [optional]
 **listOptionsAllowWatchBookmarks** | **Boolean**| allowWatchBookmarks requests watch events with type \&quot;BOOKMARK\&quot;. Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server&#39;s discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional. | [optional]
 **listOptionsResourceVersion** | **String**| resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | [optional]
 **listOptionsResourceVersionMatch** | **String**| resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | [optional]
 **listOptionsTimeoutSeconds** | **String**| Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional. | [optional]
 **listOptionsLimit** | **String**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional]
 **listOptionsContinue** | **String**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional]

### Return type

[**IonholuongutWorkflowV1alpha1CronWorkflowList**](IonholuongutWorkflowV1alpha1CronWorkflowList.md)

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

<a name="cronWorkflowServiceResumeCronWorkflow"></a>
# **cronWorkflowServiceResumeCronWorkflow**
> IonholuongutWorkflowV1alpha1CronWorkflow cronWorkflowServiceResumeCronWorkflow(namespace, name, body)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    String name = "name_example"; // String | 
    IonholuongutWorkflowV1alpha1CronWorkflowResumeRequest body = new IonholuongutWorkflowV1alpha1CronWorkflowResumeRequest(); // IonholuongutWorkflowV1alpha1CronWorkflowResumeRequest | 
    try {
      IonholuongutWorkflowV1alpha1CronWorkflow result = apiInstance.cronWorkflowServiceResumeCronWorkflow(namespace, name, body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceResumeCronWorkflow");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **name** | **String**|  |
 **body** | [**IonholuongutWorkflowV1alpha1CronWorkflowResumeRequest**](IonholuongutWorkflowV1alpha1CronWorkflowResumeRequest.md)|  |

### Return type

[**IonholuongutWorkflowV1alpha1CronWorkflow**](IonholuongutWorkflowV1alpha1CronWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

<a name="cronWorkflowServiceSuspendCronWorkflow"></a>
# **cronWorkflowServiceSuspendCronWorkflow**
> IonholuongutWorkflowV1alpha1CronWorkflow cronWorkflowServiceSuspendCronWorkflow(namespace, name, body)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    String name = "name_example"; // String | 
    IonholuongutWorkflowV1alpha1CronWorkflowSuspendRequest body = new IonholuongutWorkflowV1alpha1CronWorkflowSuspendRequest(); // IonholuongutWorkflowV1alpha1CronWorkflowSuspendRequest | 
    try {
      IonholuongutWorkflowV1alpha1CronWorkflow result = apiInstance.cronWorkflowServiceSuspendCronWorkflow(namespace, name, body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceSuspendCronWorkflow");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **name** | **String**|  |
 **body** | [**IonholuongutWorkflowV1alpha1CronWorkflowSuspendRequest**](IonholuongutWorkflowV1alpha1CronWorkflowSuspendRequest.md)|  |

### Return type

[**IonholuongutWorkflowV1alpha1CronWorkflow**](IonholuongutWorkflowV1alpha1CronWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

<a name="cronWorkflowServiceUpdateCronWorkflow"></a>
# **cronWorkflowServiceUpdateCronWorkflow**
> IonholuongutWorkflowV1alpha1CronWorkflow cronWorkflowServiceUpdateCronWorkflow(namespace, name, body)



### Example
```java
// Import classes:
import io.nholuongut.workflow.ApiClient;
import io.nholuongut.workflow.ApiException;
import io.nholuongut.workflow.Configuration;
import io.nholuongut.workflow.models.*;
import io.nholuongut.workflow.apis.CronWorkflowServiceApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:2746");

    CronWorkflowServiceApi apiInstance = new CronWorkflowServiceApi(defaultClient);
    String namespace = "namespace_example"; // String | 
    String name = "name_example"; // String | DEPRECATED: This field is ignored.
    IonholuongutWorkflowV1alpha1UpdateCronWorkflowRequest body = new IonholuongutWorkflowV1alpha1UpdateCronWorkflowRequest(); // IonholuongutWorkflowV1alpha1UpdateCronWorkflowRequest | 
    try {
      IonholuongutWorkflowV1alpha1CronWorkflow result = apiInstance.cronWorkflowServiceUpdateCronWorkflow(namespace, name, body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CronWorkflowServiceApi#cronWorkflowServiceUpdateCronWorkflow");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**|  |
 **name** | **String**| DEPRECATED: This field is ignored. |
 **body** | [**IonholuongutWorkflowV1alpha1UpdateCronWorkflowRequest**](IonholuongutWorkflowV1alpha1UpdateCronWorkflowRequest.md)|  |

### Return type

[**IonholuongutWorkflowV1alpha1CronWorkflow**](IonholuongutWorkflowV1alpha1CronWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

