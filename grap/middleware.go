package grap

import (
   "context"
   "github.com/dgrijalva/jwt-go"
   "net"
   "time"
   "net/http"
   "strings"
   "errors"
   "log"
   DDBB "github.com/siulfe/gql/Database"
)

type UserAuth struct {
   UserID    string
   Roles     string
   IPAddress string
   Token     string
}

var userCtxKey = &contextKey{"user"}

type contextKey struct {
   name string
}
type UserClaims struct {
   UserId string `json:"user_id"`
   Roles string `json:"roles"`
   jwt.StandardClaims
}

func Middleware() func(http.Handler) http.Handler {
   return func(next http.Handler) http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         token := TokenFromHttpRequest(r)
         var userId string =""
         var roles string  =""

         if tokens[token]!= 0{
            if tokens[token] >= int(time.Now().Unix()){
               tokens[token] = int(time.Now().Add(5 * time.Minute).Unix())
               userId = UserIDFromToken(token)
               roles = RolesFromToken(token)
            }else{
              _,w := DDBB.GetDB().Exec(DDBB.DESLOGED,UserIDFromToken(token))
               
               if w == nil{
                  delete(tokens,token)   
               }
               
            }
         }

         ip, _, _ := net.SplitHostPort(r.RemoteAddr)
         userAuth := UserAuth{
            UserID:    userId,
            Roles: roles,
            IPAddress: ip,
            Token: token,
         }
         // put it in context
         ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
         // and call the next with our new context
         r = r.WithContext(ctx)
         next.ServeHTTP(w, r)
      })
   }
}
func TokenFromHttpRequest(r *http.Request) string {
   reqToken := r.Header.Get("Authorization")
   log.Println("Entrada: ",reqToken)
   var tokenString string
   splitToken := strings.Split(reqToken, "Patria ")
   if len(splitToken) > 1 {
      tokenString = splitToken[1]
   }
   return tokenString
}
func UserIDFromToken(tokenString string) string {

   token, err := jwtDecode(tokenString)
   if err != nil {
      return ""
   }
   if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
      if claims == nil {
         return ""
      }
      return claims.UserId
   } else {
      return ""
   }
}
func RolesFromToken(tokenString string) string{
	token, err := jwtDecode(tokenString)
	if err!= nil{
		return ""
	}

	if claims, ok :=token.Claims.(*UserClaims); ok && token.Valid{
		if claims == nil{
			return ""
		}
		return claims.Roles
	}

	return ""
}
func ForContext(ctx context.Context) *UserAuth {
   raw := ctx.Value(userCtxKey)
   if raw == nil {
      return nil
   }
   return raw.(*UserAuth)
}
func getAuth(ctx context.Context) *UserAuth {
   return ForContext(ctx)
}

func (auth *UserAuth)validar(rol string) error{
   var w int
   error := errors.New("Usuario Invalido")

   if auth.UserID == "" || auth.Roles == "" || auth.Roles != rol{
      return error
   }

   err := DDBB.GetDB().QueryRow(DDBB.SELECT_USER_NAME_ROL,auth.UserID, auth.Roles).Scan(&w)

   if err != nil{
      return err
   }

   if w != 1{
      return error
   }

   return nil
}