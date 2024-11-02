

# IonholuongutEventsV1alpha1PulsarTrigger

PulsarTrigger refers to the specification of the Pulsar trigger.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authTokenSecret** | [**io.kubernetes.client.openapi.models.V1SecretKeySelector**](io.kubernetes.client.openapi.models.V1SecretKeySelector.md) |  |  [optional]
**connectionBackoff** | [**IonholuongutEventsV1alpha1Backoff**](IonholuongutEventsV1alpha1Backoff.md) |  |  [optional]
**parameters** | [**List&lt;IonholuongutEventsV1alpha1TriggerParameter&gt;**](IonholuongutEventsV1alpha1TriggerParameter.md) | Parameters is the list of parameters that is applied to resolved Kafka trigger object. |  [optional]
**payload** | [**List&lt;IonholuongutEventsV1alpha1TriggerParameter&gt;**](IonholuongutEventsV1alpha1TriggerParameter.md) | Payload is the list of key-value extracted from an event payload to construct the request payload. |  [optional]
**tls** | [**IonholuongutEventsV1alpha1TLSConfig**](IonholuongutEventsV1alpha1TLSConfig.md) |  |  [optional]
**tlsAllowInsecureConnection** | **Boolean** |  |  [optional]
**tlsTrustCertsSecret** | [**io.kubernetes.client.openapi.models.V1SecretKeySelector**](io.kubernetes.client.openapi.models.V1SecretKeySelector.md) |  |  [optional]
**tlsValidateHostname** | **Boolean** |  |  [optional]
**topic** | **String** |  |  [optional]
**url** | **String** |  |  [optional]



