package auth

import (
	"testing"
	"time"

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

func Test_WrongAlgorithm(t *testing.T) {
	for idx, tc := range testEachKC(t) {
		idx := idx
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			signingKCidx := (idx + 1) % len(testEachKC(t)) // get the next KC with wrapping
			signingKC := testEachKC(t)[signingKCidx].kc
			p, err := NewJWTProvider(signingKC, "issuer", 10*time.Second, 24*time.Hour)
			require.NoError(t, err)
			token, err := p.CreateToken(NewAccessTokenClaims("user_id", "user_name", "email", []string{"permission1", "permission2"}))
			require.NoError(t, err)

			_, err = jwt.Parse(token.Token, tc.kc.VerificationKey)
			require.Error(t, err)
		})
	}
}

func Test_VerificationKey(t *testing.T) {
	for _, tc := range testEachKC(t) {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			p, err := NewJWTProvider(tc.kc, "issuer", 10*time.Second, 24*time.Hour)
			require.NoError(t, err)
			token, err := p.CreateToken(NewAccessTokenClaims("user_id", "user_name", "email", []string{"permission1", "permission2"}))
			require.NoError(t, err)

			_, err = jwt.Parse(token.Token, tc.kc.VerificationKey)
			require.NoError(t, err)
		})
	}
}

func Test_NewHS512KeyContainerErrors(t *testing.T) {
	tests := []struct {
		name       string
		signingKey string
		shouldErr  bool
	}{
		{
			name:       "empty",
			signingKey: "",
			shouldErr:  true,
		},
		{
			name:       "valid",
			signingKey: "some_secret",
			shouldErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewHS512KeyContainer(tt.signingKey)
			if tt.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_NewRS512KeyContainerErrors(t *testing.T) {
	tests := []struct {
		name                string
		signingKeyPath      string
		verificationKeyPath string
		shouldErr           bool
	}{
		{
			name:                "signing key not found",
			signingKeyPath:      "test/jwtRS512.key.notfound",
			verificationKeyPath: "test/jwtRS512.key.pub",
			shouldErr:           true,
		},
		{
			name:                "verification key not found",
			signingKeyPath:      "test/jwtRS512.key",
			verificationKeyPath: "test/jwtRS512.key.pub.notfound",
			shouldErr:           true,
		},
		{
			name:                "signing key invalid",
			signingKeyPath:      "test/jwtRS512.key.pub",
			verificationKeyPath: "test/jwtRS512.key.pub",
			shouldErr:           true,
		},
		{
			name:                "verification key invalid",
			signingKeyPath:      "test/jwtRS512.key",
			verificationKeyPath: "test/jwtRS512.key",
			shouldErr:           true,
		},
		{
			name:                "signing key empty",
			signingKeyPath:      "",
			verificationKeyPath: "test/jwtRS512.key.pub",
			shouldErr:           true,
		},
		{
			name:                "verification key empty",
			signingKeyPath:      "test/jwtRS512.key",
			verificationKeyPath: "",
			shouldErr:           true,
		},
		{
			name:                "valid",
			signingKeyPath:      "test/jwtRS512.key",
			verificationKeyPath: "test/jwtRS512.key.pub",
			shouldErr:           false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewRS512KeyContainer(tt.signingKeyPath, tt.verificationKeyPath)
			if tt.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
