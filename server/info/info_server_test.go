package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/nholuongut/argo-workflows/v3/server/auth"
	"github.com/nholuongut/argo-workflows/v3/server/auth/types"
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimsKey, &types.Claims{Claims: jwt.Claims{Issuer: "my-iss", Subject: "my-sub"}, Groups: []string{"my-group"}, Email: "my@email", EmailVerified: true, ServiceAccountName: "my-sa"})
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
		assert.Equal(t, []string{"my-group"}, info.Groups)
		assert.Equal(t, "my@email", info.Email)
		assert.True(t, info.EmailVerified)
		assert.Equal(t, "my-sa", info.ServiceAccountName)
	}
}
