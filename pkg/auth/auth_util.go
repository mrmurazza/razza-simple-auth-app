package auth

import (
	"crypto/md5"
	"crypto/sha256"
	"dealljobs/config"
	"dealljobs/domain/user"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var JwtSigningMethod = jwt.SigningMethodHS256

type MyClaims struct {
	jwt.RegisteredClaims
	Data interface{} `json:"data"`
}

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()
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

			key := []byte(cfg.JWTSignatureKey)
			return key, nil
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

	str := base64.StdEncoding.EncodeToString(sha256pass[:])
	return str
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
	cfg := config.Get()
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWTExpiryDuration)),
		},
		Data: data,
	}

	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)

	key := []byte(cfg.JWTSignatureKey)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
