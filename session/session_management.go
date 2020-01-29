package session

import (
	"context"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)
type contextKey string
var ctxUserSessionKey = contextKey("signed_in_user_session")

func IsAuthenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie("token")
		tknStr := c.Value
		claims := &entity.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return entity.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.Write([]byte("access not autorized"))
				return
			}
			w.Write([]byte("access not autorized"))
			return
		}
		if !tkn.Valid {
			w.Write([]byte("access not autorized"))
			return
		}

		//next.ServeHTTP(w, r)
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func Generate(Id int, UserName, ProfilePicture string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &entity.Claims{
		Username:       UserName,
		Id:             Id,
		ProfilePicture: ProfilePicture,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//log.Println(*claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(entity.JwtKey)
	return tokenString, err
	//http.SetCookie(writer, &http.Cookie{
	//	Name:    "token",
	//	Value:   tokenString,
	//	Expires: expirationTime,
	//})
	//http.Redirect(writer,request,"/",http.StatusSeeOther)
}

func GetSessionData(w http.ResponseWriter, r *http.Request) *entity.Claims {
	c, err := r.Cookie("token")
	if err!=nil{
		return nil

	}
	tknStr := c.Value
	claims := &entity.Claims{}
	jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return entity.JwtKey, nil
	})
	return claims

}

func RenewSession(username string, id int,profilePicture string, w http.ResponseWriter){
	RemoveSession(w)
	tokenString,err := Generate(id,username,profilePicture)
	log.Println(err)
	c := &http.Cookie{
		Name:       "token",
		Value:      tokenString,
		Path:"/",
	}
	w.Header().Set("Set-Cookie", c.String())
}

func RemoveSession(w http.ResponseWriter) {
	c := http.Cookie{
		Name:    "token",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
		Value:   "",
		Path:"/",
	}
	http.SetCookie(w, &c)
}
