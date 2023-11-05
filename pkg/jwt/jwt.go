package jwt

import (
	"context"
	"encoding/json"
	"errors"
	jwt_go "github.com/dgrijalva/jwt-go"
	"time"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/consts"
)

// LoginTokenInfo 登录令牌信息
type LoginTokenInfo struct {
	AccessToken string `json:"access_token"` // 访问令牌
	TokenType   string `json:"token_type"`   // 令牌类型
	ExpiresAt   int64  `json:"expires_at"`   // 令牌到期时间戳
}

var Auth *JWTAuth

// JWTAuth jwt认证
type JWTAuth struct {
	signingMethod jwt_go.SigningMethod
	signingKey    interface{}
	keyfunc       jwt_go.Keyfunc
	expired       int
	tokenType     string
}

var defaultOptions = &JWTAuth{
	tokenType:     "Bearer",
	expired:       30 * 24 * 3600,
	signingMethod: jwt_go.SigningMethodHS512,
	signingKey:    []byte(consts.AUTH_SIGNING_KEY),
	keyfunc: func(t *jwt_go.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt_go.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(consts.AUTH_SIGNING_KEY), nil
	},
}

type SubjectInfo struct {
	UserID string         `json:"user_id"` // 用户ID
	Phone  string         `json:"phone"`   // 手机号
	Role   model.UserRole `json:"role"`
}

func Init() {
	Auth = &JWTAuth{
		signingMethod: defaultOptions.signingMethod,
		signingKey:    []byte(consts.AUTH_SIGNING_KEY), // notice: []byte
		keyfunc: func(t *jwt_go.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt_go.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid")
			}
			return []byte(consts.AUTH_SIGNING_KEY), nil
		},
		expired:   defaultOptions.expired,
		tokenType: defaultOptions.tokenType,
	}
}

// GenerateToken 生成令牌
func (jwt *JWTAuth) GenerateToken(ctx context.Context, subject *SubjectInfo) (*LoginTokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(jwt.expired) * time.Second).Unix()

	str, _ := json.Marshal(subject)
	token := jwt_go.NewWithClaims(jwt.signingMethod, &jwt_go.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   string(str),
	})

	tokenString, err := token.SignedString(jwt.signingKey)
	if err != nil {
		return nil, err
	}

	tokenInfo := &LoginTokenInfo{
		ExpiresAt:   expiresAt,
		TokenType:   jwt.tokenType,
		AccessToken: tokenString,
	}
	return tokenInfo, nil
}

// ParseUserID 解析用户ID
func (jwt *JWTAuth) ParseUserID(ctx context.Context, tokenString string) (string, error) {
	if len(tokenString) == 0 {
		return "", errors.New("empty token")
	}
	token, err := jwt_go.ParseWithClaims(tokenString, &jwt_go.StandardClaims{}, jwt.keyfunc)
	if err != nil {
		return "", err
	}
	return token.Claims.(*jwt_go.StandardClaims).Subject, nil
}
