package util

import (
	"crypto/md5"
	"crypto/sha256"
	"dealljobs/domain/user"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var ApplicationName = "My Simple JWT App"
var LoginExpirationDuration = time.Duration(5) * time.Minute
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("the secret of kalimdor")

type MyClaims struct {
	jwt.RegisteredClaims
	Data interface{} `json:"data"`
}

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JwtSigningMethod {
				return nil, fmt.Errorf("signing method invalid")
			}

			return JwtSignatureKey, nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userInfo", claims["data"])
		c.Next()
	}
}

func AuthenticateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := GetUserInfo(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}

		if u["Role"] != string(user.RoleAdmin) {
			c.JSON(http.StatusForbidden, gin.H{"error": "you are forbidden to access this"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func EncryptPassword(password string) string {
	md5pass := md5.Sum([]byte(password))
	sha256pass := sha256.Sum256(md5pass[:])

	return string(sha256pass[:])
}

func GetUserInfo(c *gin.Context) (map[string]interface{}, error) {
	userInfo, exist := c.Get("userInfo")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo from auth invalid"})
		c.Abort()
		return nil, errors.New("userInfo from auth invalid")
	}

	u, ok := userInfo.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo from auth invalid"})
		c.Abort()
		return nil, errors.New("userInfo from auth invalid")
	}

	return u, nil
}

func TokenizeData(data interface{}) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    ApplicationName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LoginExpirationDuration)),
		},
		Data: data,
	}

	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
