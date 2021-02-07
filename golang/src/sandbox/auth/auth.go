package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"golang.org/x/xerrors"
)

const privateKey string = `-----BEGIN PRIVATE KEY-----
MIIBOwIBAAJBALfeXtQl/mYJ+mMuOSmbbKOuOScigQW+ZTpzKG/LrjS3HpjVwdfh
2XwQyyjglzfkmjZCjXMuXdDHsQ5M7OymO4sCAwEAAQJBAJb2TIWCbEz7BElKOkSf
B6Ob9/DZs3UzzYkf46NmZ7F/A16Gog60DZo6aTBsd0mtAftVk8SMJKzg+oWRTjfg
Y4ECIQDiIcK3knWaLvIxpiA9sxwH51+Y9E18uJLMALd9oH8+pwIhANAnjg0jcXDw
+SGQ9ZlHtnma0iLh1o8RVCVYbpghrrx9AiAFIXz4i1Mtx2jpZXqPy9OT/lT19H0Z
mlqCVHXVXOHmOwIgXdBF/HPp784bal5r4n5oplv5s1D3o5laxE4b4iLkHv0CIQCQ
SvxjVqFkHyA5VvY0zs1c/dbgl+3+Mhkhwq2cfojg/Q==
-----END PRIVATE KEY-----`

const verifyKey string = `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALfeXtQl/mYJ+mMuOSmbbKOuOScigQW+
ZTpzKG/LrjS3HpjVwdfh2XwQyyjglzfkmjZCjXMuXdDHsQ5M7OymO4sCAwEAAQ==
-----END PUBLIC KEY-----`

type Auth struct {
	r *http.Request
}

func NewAuth(r *http.Request) *Auth {
	return &Auth{r}
}

func (a *Auth) Valid() error {
	v, err := jwt.ParseRSAPublicKeyFromPEM([]byte(verifyKey))
	if err != nil {
		return xerrors.Errorf("verify key is odd: %w", err)
	}

	token, err := request.ParseFromRequest(a.r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, xerrors.Errorf("unexpected signing method: %w", token.Header["alg"])
		}
		return v, nil
	})

	if err != nil {
		return xerrors.Errorf("fail to parse request: %w", err)
	}
	if !token.Valid {
		return xerrors.New("token is invalid")
	}
	return nil
}

func GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"name":  "hoge",
		"sub":   "54546557354",
		"admin": true,
		"iat":   time.Now(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	x, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", xerrors.Errorf("fail parse pem key: %w", err)
	}

	t, err := token.SignedString(x)
	if err != nil {
		return "", xerrors.Errorf("fail sign token: %w", err)
	}
	return t, nil
}
