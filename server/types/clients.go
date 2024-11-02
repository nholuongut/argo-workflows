package types

import (
	eventsource "github.com/nholuongut/argo-events/pkg/client/eventsource/clientset/versioned"
	sensor "github.com/nholuongut/argo-events/pkg/client/sensor/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	workflow "github.com/nholuongut/argo-workflows/v3/pkg/client/clientset/versioned"
)

type Clients struct {
	Dynamic     dynamic.Interface
	Workflow    workflow.Interface
	Sensor      sensor.Interface
	EventSource eventsource.Interface
	Kubernetes  kubernetes.Interface
}
