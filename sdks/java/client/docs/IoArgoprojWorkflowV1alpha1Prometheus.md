

# IonholuongutWorkflowV1alpha1Prometheus

Prometheus is a prometheus metric to be emitted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**counter** | [**IonholuongutWorkflowV1alpha1Counter**](IonholuongutWorkflowV1alpha1Counter.md) |  |  [optional]
**gauge** | [**IonholuongutWorkflowV1alpha1Gauge**](IonholuongutWorkflowV1alpha1Gauge.md) |  |  [optional]
**help** | **String** | Help is a string that describes the metric | 
**histogram** | [**IonholuongutWorkflowV1alpha1Histogram**](IonholuongutWorkflowV1alpha1Histogram.md) |  |  [optional]
**labels** | [**List&lt;IonholuongutWorkflowV1alpha1MetricLabel&gt;**](IonholuongutWorkflowV1alpha1MetricLabel.md) | Labels is a list of metric labels |  [optional]
**name** | **String** | Name is the name of the metric | 
**when** | **String** | When is a conditional statement that decides when to emit the metric |  [optional]



