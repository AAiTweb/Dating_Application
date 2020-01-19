package EyosiHand
//
//import (
//	"github.com/biniyam112/TheDatingApp/Dating_Application/entity"
//	"html/template"
//	"net/http"
//)
//
//type contextKey string
//
//var ctxUserSessionKey = contextKey("signed_in_user_session")
//
//
//type UserHandler struct {
//	templ *template.Template
//	session *entity.Session
//}
//
//func NewUserHandler(t *template.Template,s *entity.Session) UserHandler{
//	return UserHandler{templ:t,session:s};
//}
//
//func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
//
//
//}
//
//func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		ok := uh.loggedIn(r)
//		if !ok {
//			http.Redirect(w, r, "/login", http.StatusSeeOther)
//			return
//		}
//		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.session)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	}
//	return http.HandlerFunc(fn)
//}