package middlewares

import (
	"github.com/golang-jwt/jwt"
)

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		   jwt.MapClaims{ 
		   "username": username, 
		   "exp": time.Now().Add(time.Hour * 48).Unix(), 
		   })
   
	  tokenString, err := token.SignedString(secretKey)
	  if err != nil {
	  return "", err
	  }
   
	return tokenString, nil
}


func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secretKey, nil
	})
   
	if err != nil {
	   return err
	}
   
	if !token.Valid {
	   return fmt.Errorf("invalid token")
	}
   
	return nil
 }