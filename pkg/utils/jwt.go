package util

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// var screet = setting.FileConfigSetting.App.JwtSecret
// var jwtSecret = []byte("secreet")

// Claims :
type Claims struct {
	UserID    string `json:"user_id,omitempty"`
	RoleID    string `json:"role_id,omitempty"`
	CompanyID int    `json:"company_id,omitempty"`
	jwt.StandardClaims
}

// GenerateToken :
func GenerateToken(id string, role_id string, company_id int) (string, error) {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.json': %v", err)
	}

	var screet = viper.GetString(`jwt_secret`)
	var jwtSecret = []byte(screet)
	// Set custom claims
	claims := &Claims{
		UserID:    id,
		RoleID:    role_id,
		CompanyID: company_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken :
func ParseToken(token string) (*Claims, error) {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.json': %v", err)
	}

	var screet = viper.GetString(`jwt_secret`)
	var jwtSecret = []byte(screet)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
