

# IonholuongutEventsV1alpha1AMQPEventSource


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auth** | [**IonholuongutEventsV1alpha1BasicAuth**](IonholuongutEventsV1alpha1BasicAuth.md) |  |  [optional]
**connectionBackoff** | [**IonholuongutEventsV1alpha1Backoff**](IonholuongutEventsV1alpha1Backoff.md) |  |  [optional]
**consume** | [**IonholuongutEventsV1alpha1AMQPConsumeConfig**](IonholuongutEventsV1alpha1AMQPConsumeConfig.md) |  |  [optional]
**exchangeDeclare** | [**IonholuongutEventsV1alpha1AMQPExchangeDeclareConfig**](IonholuongutEventsV1alpha1AMQPExchangeDeclareConfig.md) |  |  [optional]
**exchangeName** | **String** |  |  [optional]
**exchangeType** | **String** |  |  [optional]
**jsonBody** | **Boolean** |  |  [optional]
**metadata** | **Map&lt;String, String&gt;** |  |  [optional]
**queueBind** | [**IonholuongutEventsV1alpha1AMQPQueueBindConfig**](IonholuongutEventsV1alpha1AMQPQueueBindConfig.md) |  |  [optional]
**queueDeclare** | [**IonholuongutEventsV1alpha1AMQPQueueDeclareConfig**](IonholuongutEventsV1alpha1AMQPQueueDeclareConfig.md) |  |  [optional]
**routingKey** | **String** |  |  [optional]
**tls** | [**IonholuongutEventsV1alpha1TLSConfig**](IonholuongutEventsV1alpha1TLSConfig.md) |  |  [optional]
**url** | **String** |  |  [optional]
**urlSecret** | [**io.kubernetes.client.openapi.models.V1SecretKeySelector**](io.kubernetes.client.openapi.models.V1SecretKeySelector.md) |  |  [optional]



