package auth

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func testKCHS512(t *testing.T) KeyContainer {
	kc, err := NewHS512KeyContainer("some_secret")
	require.NoError(t, err)
	return kc
}

func testKCRS512(t *testing.T) KeyContainer {
	kc, err := NewRS512KeyContainer(
		"test/jwtRS512.key",
		"test/jwtRS512.key.pub",
	)
	require.NoError(t, err)
	return kc
}

type keyContainerTestCase struct {
	name                  string
	kc                    KeyContainer
	expectedSigningMethod jwt.SigningMethod
	wrongSigningMethod    jwt.SigningMethod
}

func testEachKC(t *testing.T) []keyContainerTestCase {
	return []keyContainerTestCase{
		{
			name:                  "HS512",
			kc:                    testKCHS512(t),
			expectedSigningMethod: jwt.SigningMethodHS512,
			wrongSigningMethod:    jwt.SigningMethodRS512,
		},
		{
			name:                  "RS512",
			kc:                    testKCRS512(t),
			expectedSigningMethod: jwt.SigningMethodRS512,
			wrongSigningMethod:    jwt.SigningMethodHS512,
		},
	}
}

func Test_SignedToken(t *testing.T) {
	for _, tc := range testEachKC(t) {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			token, err := tc.kc.SignedToken(Claims{})
			require.NoError(t, err)
			require.NotEmpty(t, token)
		})
	}
}

func Test_WrongAlgorithm(t *testing.T) {
	for idx, tc := range testEachKC(t) {
		idx := idx
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			signingKCidx := (idx + 1) % len(testEachKC(t)) // get the next KC with wrapping
			signingKC := testEachKC(t)[signingKCidx].kc
			token, err := signingKC.SignedToken(Claims{})
			require.NoError(t, err)

			_, err = jwt.Parse(token, tc.kc.VerificationKeyFunc)
			require.Error(t, err)
		})
	}
}
