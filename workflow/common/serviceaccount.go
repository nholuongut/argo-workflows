package common

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/nholuongut/argo-workflows/v3/errors"
)

// GetServiceAccountTokenName returns the name of the first referenced ServiceAccountToken secret of the service account.
func GetServiceAccountTokenName(ctx context.Context, clientset kubernetes.Interface, namespace, name string) (string, error) {
	serviceAccount, err := clientset.CoreV1().ServiceAccounts(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	if len(serviceAccount.Secrets) == 0 {
		return "", errors.Errorf("", "Service account %s/%s does not have any token", serviceAccount.Namespace, serviceAccount.Name)
	}
	return serviceAccount.Secrets[0].Name, nil
}
