package middleware

import (
	"gorestapi/modules/datainit"
	"net/http"
)

type user struct {
	Login        string `json:"login"`
	PasswordHash string `json:"hash"`
}

//TODO Подключить СУБД в качестве источника

//Init test data
var user1 user

func init() {
	user1.Login = "test"
	user1.PasswordHash = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
}

//TODO Перевести на валидацию JWT
func Authentication(next http.Handler) http.Handler {
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
	for i := range datainit.UsersData {
		return login == datainit.UsersData[i].Login && hash == datainit.UsersData[i].Password
	}
	return false
}
