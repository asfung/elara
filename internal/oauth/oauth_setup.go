package oauth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func InitProviders() {
	key := os.Getenv("SESSION_KEY")
	if key == "" {
		key = "dev-session-key-please-change"
	}

	store := sessions.NewCookieStore([]byte(key))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30,
		HttpOnly: true,
	}

	goth.UseProviders(
		google.New(
			os.Getenv("OAUTH_GOOGLE_CLIENT_ID"),
			os.Getenv("OAUTH_GOOGLE_CLIENT_SECRET"),
			os.Getenv("OAUTH_GOOGLE_CALLBACK_URL"),
			"email", "profile",
		),

		// more provider

	)

	gothicStore := store
	GothicStoreSetter(gothicStore)
}
