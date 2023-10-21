package middleware

import (
	"encoding/json"
	"errors"
	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/jwt"
)

func UserAuthMiddleware(skipPrefix ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if !matchPrefix(path, skipPrefix) {
			if err := verifyToken(c); err != nil {
				c.Abort()
			}
		}

		c.Next()
	}
}

func verifyToken(c *gin.Context) error {
	subjectJsonStr, err := jwt.Auth.ParseUserID(c, GetToken(c))
	if vErr, ok := err.(*jwt_go.ValidationError); ok {
		if vErr.Errors&(jwt_go.ValidationErrorExpired|jwt_go.ValidationErrorNotValidYet) != 0 {
			return errors.New("token expired")
		}
	} else if err != nil {
		return err
	}
	var subjectInfo *jwt.SubjectInfo
	err = json.Unmarshal([]byte(subjectJsonStr), &subjectInfo)
	if err != nil {
		return err
	}

	c.Set(consts.CONTEXT_HEADER_USER_OPEN_ID, subjectInfo.OpenID)
	c.Set(consts.CONTEXT_HEADER_USERNAME, subjectInfo.UserName)
	c.Set(consts.CONTEXT_HEADER_USER_ROLE, subjectInfo.Role)
	c.Set(consts.CONTEXT_HEADER_USER_ID, subjectInfo.UserID)
	c.Set(consts.CONTEXT_HEADER_USER_AUTH_INFO, subjectInfo)
	return nil
}

func matchPrefix(path string, prefixes []string) bool {
	if len(prefixes) == 0 {
		return false
	}
	pathLen := len(path)
	for _, p := range prefixes {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return true
		}
	}
	return false
}

func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}
