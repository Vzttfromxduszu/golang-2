package jwt

import (
	"encoding/json"
	"log"

	"github.com/dgrijalva/jwt-go"
)

// 解码token
type DecodedToken struct {
	Iat      int    `json:"iat"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Iss      string `json:"iss"`
	IsAdmin  bool   `json:"isAdmin"`
}

// 生成token
func GenerateToken(claims *jwt.Token, secret string) (token string) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ = claims.SignedString(hmacSecret)

	return
}

// 验证token
func VerifyToken(token string, secret string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})
	// jwt.Parse() 解析 Token 并验证签名,如果签名验证成功，decoded.Valid 为 true；否则为 false
	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)         // 序列化为 JSON 字符串
	jsonErr := json.Unmarshal(jsonString, &decodedToken) // 反序列化为 DecodedToken 结构体
	if jsonErr != nil {
		log.Print(jsonErr)
	}

	return &decodedToken
}
