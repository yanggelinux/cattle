package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/yanggelinux/cattle/global"
	"github.com/yanggelinux/cattle/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	UserID    int64  `json:"user_id"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.JwtSecret)
}

func GenerateToken(appKey, appSecret string, userID int64) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(global.JWTSetting.JwtExpire) * time.Second)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.JwtIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(GetJWTSecret())
	return
}

func ParseToken(token string) (claims *Claims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return
}
