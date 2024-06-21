package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	users "github.com/jvpereirarocha/jwt-auth/internal/domain/users"
)

func subToUser(user *users.User) string {
	if user.Email != "" {
		return user.Email
	}

	return user.Username
}

func timeToExpiration(minutes int) int64 {
	return time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
}

func createToken(user *users.User, secretKey string, app string, minutesToExpire int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subToUser(user),
		"iss": app,
		"exp": timeToExpiration(minutesToExpire),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
