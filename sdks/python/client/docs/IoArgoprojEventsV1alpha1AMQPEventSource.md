# IonholuongutEventsV1alpha1AMQPEventSource


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auth** | [**IonholuongutEventsV1alpha1BasicAuth**](IonholuongutEventsV1alpha1BasicAuth.md) |  | [optional] 
**connection_backoff** | [**IonholuongutEventsV1alpha1Backoff**](IonholuongutEventsV1alpha1Backoff.md) |  | [optional] 
**consume** | [**IonholuongutEventsV1alpha1AMQPConsumeConfig**](IonholuongutEventsV1alpha1AMQPConsumeConfig.md) |  | [optional] 
**exchange_declare** | [**IonholuongutEventsV1alpha1AMQPExchangeDeclareConfig**](IonholuongutEventsV1alpha1AMQPExchangeDeclareConfig.md) |  | [optional] 
**exchange_name** | **str** |  | [optional] 
**exchange_type** | **str** |  | [optional] 
**json_body** | **bool** |  | [optional] 
**metadata** | **{str: (str,)}** |  | [optional] 
**queue_bind** | [**IonholuongutEventsV1alpha1AMQPQueueBindConfig**](IonholuongutEventsV1alpha1AMQPQueueBindConfig.md) |  | [optional] 
**queue_declare** | [**IonholuongutEventsV1alpha1AMQPQueueDeclareConfig**](IonholuongutEventsV1alpha1AMQPQueueDeclareConfig.md) |  | [optional] 
**routing_key** | **str** |  | [optional] 
**tls** | [**IonholuongutEventsV1alpha1TLSConfig**](IonholuongutEventsV1alpha1TLSConfig.md) |  | [optional] 
**url** | **str** |  | [optional] 
**url_secret** | [**SecretKeySelector**](SecretKeySelector.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


