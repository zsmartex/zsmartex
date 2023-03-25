package session

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/mileusna/useragent"
)

type SessionData struct {
	ID              uuid.UUID `json:"id,omitempty"`
	UserID          int64     `json:"user_id,omitempty"`
	UID             string    `json:"uid,omitempty"`
	UserIP          string    `json:"user_ip,omitempty"`
	UserIPCountry   string    `json:"user_ip_country,omitempty"`
	UserAgent       string    `json:"user_agent,omitempty"`
	CrfsToken       string    `json:"crfs_token,omitempty"`
	AuthenticatedAt time.Time `json:"authenticated_at,omitempty"`
	CurrentSession  bool      `json:"-"`
}

func (i SessionData) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (m *SessionData) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type Session struct {
	id    uuid.UUID
	fresh bool // if new session
	ctx   context.Context
	exp   time.Duration
	store *Store
	Data  SessionData // data of session
}

func acquireSession() *Session {
	s := new(Session)

	return s
}

func (s *Session) SetUserID(userID int64) {
	s.Data.UserID = userID
}

func (s *Session) SetUID(uid string) {
	s.Data.UID = uid
}

func (s *Session) SetUserIP(userIP string) {
	s.Data.UserIP = userIP
}

func (s *Session) SetUserIPCountry(UserIPCountry string) {
	s.Data.UserIPCountry = UserIPCountry
}

func (s *Session) SetUserAgent(UserAgent string) {
	s.Data.UserAgent = UserAgent
}

func (s *Session) SetCrfsToken(CrfsToken string) {
	s.Data.CrfsToken = CrfsToken
}

func (s *Session) SetAuthenticatedAt(AuthenticatedAt time.Time) {
	s.Data.AuthenticatedAt = AuthenticatedAt
}

func (s *Session) Save(ctx context.Context) error {
	key := fmt.Sprintf("_user_session:%s:%s", s.Data.UID, s.id)

	ua := useragent.Parse(s.Data.UserAgent)
	if ua.Mobile || ua.Tablet || ua.IsIOS() || ua.IsAndroid() {
		s.exp = s.store.ExpirationApp
	} else {
		s.exp = s.store.ExpirationWeb
	}

	return s.store.redis.Set(ctx, key, s.Data, s.exp)
}

func (s *Session) Destroy(ctx context.Context) error {
	key := fmt.Sprintf("_user_session:%s:%s", s.Data.UID, s.id)
	err := s.store.redis.Delete(ctx, key)
	if err != nil {
		return nil
	}

	return nil
}

func (s *Session) SetSession(ctx context.Context, w http.ResponseWriter) {
	ua := useragent.Parse(s.Data.UserAgent)
	if ua.Mobile || ua.Tablet || ua.IsIOS() || ua.IsAndroid() {
		s.exp = s.store.ExpirationApp
	} else {
		s.exp = s.store.ExpirationWeb
	}

	cookieSameSite := http.SameSiteDefaultMode
	switch s.store.CookieSameSite {
	case "lax":
		cookieSameSite = http.SameSiteLaxMode
	case "strict":
		cookieSameSite = http.SameSiteStrictMode
	case "none":
		cookieSameSite = http.SameSiteNoneMode
	}

	http.SetCookie(w, &http.Cookie{
		Name:     s.store.CookieKey,
		Value:    s.id.String(),
		Path:     "/",
		Domain:   s.store.CookieDomain,
		HttpOnly: s.store.CookieHTTPOnly,
		Secure:   s.store.CookieSecure,
		SameSite: cookieSameSite,
		Expires:  time.Now().Add(s.exp),
		MaxAge:   int(s.exp),
	})
}

func (s *Session) DelSession(ctx context.Context, w http.ResponseWriter) {
	cookieSameSite := http.SameSiteDefaultMode
	switch s.store.CookieSameSite {
	case "lax":
		cookieSameSite = http.SameSiteLaxMode
	case "strict":
		cookieSameSite = http.SameSiteStrictMode
	case "none":
		cookieSameSite = http.SameSiteNoneMode
	}

	http.SetCookie(w, &http.Cookie{
		Name:     s.store.CookieKey,
		Value:    "",
		Path:     "/",
		Domain:   s.store.CookieDomain,
		HttpOnly: s.store.CookieHTTPOnly,
		Secure:   s.store.CookieSecure,
		SameSite: cookieSameSite,
		MaxAge:   -1,
	})
}
