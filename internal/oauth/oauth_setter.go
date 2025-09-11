package oauth

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func GothicStoreSetter(s *sessions.CookieStore) {
	gothic.Store = s
}
