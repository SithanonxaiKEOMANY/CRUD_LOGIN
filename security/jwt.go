// package security

// import (
// 	"github.com/golang-jwt/jwt/v4"
// 	"time"
// )

// var (
// 	JwtSecretKey     = []byte("ceit_")
// 	JwtPartnerSecret = []byte("jiv313a2")
// 	//jwtSigningMethod = jwt.SigningMethodHS256.Name
// )

// func NewAccessToken(userId string) (string, error) {
// 	claims := jwt.StandardClaims{
// 		Id:        userId,
// 		Issuer:    userId,
// 		IssuedAt:  time.Now().Unix(),
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 	}
// 	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedString, err := withClaims.SignedString(JwtSecretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedString, nil
// }


package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	JwtSecretKey     = []byte("ceit_")
	JwtPartnerSecret = []byte("jiv313a2")
)

func NewAccessToken(userId string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userId,
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := withClaims.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}


func CheckToken(tokenString string) (bool, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		if claims.ExpiresAt > time.Now().Unix() {
			return true, nil
		} else {
			return false, errors.New("token has expired")
		}
	} else {
		return false, errors.New("invalid token")
	}
}





// func createToken(username string) (string, error) {
//  token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
//         jwt.MapClaims{ 
//         "username": username, 
//         "exp": time.Now().Add(time.Hour * 24).Unix(), 
//         })

//    tokenString, err := token.SignedString(secretKey)
//    if err != nil {
//    return "", err
//    }

//  return tokenString, nil
// }

// func verifyToken(tokenString string) error {
//    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//       return secretKey, nil
//    })
  
//    if err != nil {
//       return err
//    }
  
//    if !token.Valid {
//       return fmt.Errorf("invalid token")
//    }
  
//    return nil
// }
