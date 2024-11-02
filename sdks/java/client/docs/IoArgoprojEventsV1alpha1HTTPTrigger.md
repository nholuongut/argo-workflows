

# IonholuongutEventsV1alpha1HTTPTrigger


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**basicAuth** | [**IonholuongutEventsV1alpha1BasicAuth**](IonholuongutEventsV1alpha1BasicAuth.md) |  |  [optional]
**headers** | **Map&lt;String, String&gt;** |  |  [optional]
**method** | **String** |  |  [optional]
**parameters** | [**List&lt;IonholuongutEventsV1alpha1TriggerParameter&gt;**](IonholuongutEventsV1alpha1TriggerParameter.md) | Parameters is the list of key-value extracted from event&#39;s payload that are applied to the HTTP trigger resource. |  [optional]
**payload** | [**List&lt;IonholuongutEventsV1alpha1TriggerParameter&gt;**](IonholuongutEventsV1alpha1TriggerParameter.md) |  |  [optional]
**secureHeaders** | [**List&lt;IonholuongutEventsV1alpha1SecureHeader&gt;**](IonholuongutEventsV1alpha1SecureHeader.md) |  |  [optional]
**timeout** | **String** |  |  [optional]
**tls** | [**IonholuongutEventsV1alpha1TLSConfig**](IonholuongutEventsV1alpha1TLSConfig.md) |  |  [optional]
**url** | **String** | URL refers to the URL to send HTTP request to. |  [optional]



