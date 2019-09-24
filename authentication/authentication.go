package authentication

import "net/http"

type user struct {
	Login        string `json:"login"`
	PasswordHash string `json:"hash"`
}

var user1 user

func DataInit() {
	user1.Login = "test"
	user1.PasswordHash = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
}

//TODO Перевести на валидацию JWT
var Authentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, hash, ok := r.BasicAuth()
		if !ok || !checkUsernameAndPassword(user, hash) {
			http.Error(w, "Wrong login or password", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func checkUsernameAndPassword(login, hash string) bool {
	return login == user1.Login && hash == user1.PasswordHash
}
