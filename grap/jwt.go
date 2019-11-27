package grap

import "github.com/dgrijalva/jwt-go"


var mySigninKey = []byte("grap.io")


func jwtDecode(token string)(*jwt.Token,error){
	return jwt.ParseWithClaims(token, &UserClaims{},func(token *jwt.Token)(interface{}, error){
		return mySigninKey, nil
	})
}


func jwtCreate(userID string,expiredAt int64,role string) string {
	claims := UserClaims{
		userID,
		role,
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer: "grap",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	ss, _ := token.SignedString(mySigninKey)
	return ss
}