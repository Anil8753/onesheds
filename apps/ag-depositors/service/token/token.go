package token

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	User   string
	UserId string
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func GenerateTokenPair(u *UserData) (*TokenPair, error) {

	secret := os.Getenv("API_SECRET")
	accessTokenTimespan, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_MIN_LIFESPAN"))
	if err != nil {
		accessTokenTimespan = 5
		log.Default().Println("env ACCESS_TOKEN_MIN_LIFESPAN not set. using default value 5 minutes")
	}

	refreshTokenTimespan, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		refreshTokenTimespan = 6
		log.Default().Println("env ACCESS_TOKEN_MIN_LIFESPAN not set. using default value 6 hours")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["authorized"] = true
	claims["user"] = u.User
	claims["userUniqueId"] = u.UserId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(accessTokenTimespan)).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["user"] = u.User
	rtClaims["exp"] = time.Now().Add(time.Hour * time.Duration(refreshTokenTimespan)).Unix()

	rt, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{AccessToken: t, RefreshToken: rt}, nil
}

func ExtractUserData(c *gin.Context) (*UserData, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user, ok := claims["user"]
		if !ok {
			return nil, errors.New("user not found")
		}

		userUniqueId, ok := claims["userUniqueId"]
		if !ok {
			return nil, errors.New("userUniqueId not found")
		}

		return &UserData{User: user.(string), UserId: userUniqueId.(string)}, nil
	}

	return nil, errors.New("user token is not valid")
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func GetUserFromRefreshToken(tokenStr string) (string, error) {
	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		u, ok := claims["user"]
		if !ok {
			return "", errors.New("user key not found in the claim")
		}

		return u.(string), nil
	}

	return "", err
}
