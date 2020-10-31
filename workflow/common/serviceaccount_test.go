package common

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/argoproj/argo/test/util"
)

// TestGetServiceAccountTokenName verifies service account token retrieved by service account name.
func TestGetServiceAccountTokenName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	clientset := fake.NewSimpleClientset()
	_, err := util.CreateServiceAccountWithToken(clientset, "", "test", "test-token")
	assert.NoError(t, err)
	tokenName, err := GetServiceAccountTokenName(ctx, clientset, "", "test")
	assert.NoError(t, err)
	assert.Equal(t, "test-token", tokenName)
}
