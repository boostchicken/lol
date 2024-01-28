package sesssionstore // 

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionManager struct {
	store sessions.Store
}

func (*SessionManager) New(secretKey string) *SessionManager {
	return &SessionManager{
		store: sessions.NewCookieStore([]byte(secretKey)),
	}
}

func (sm *SessionManager) CreateSession(w http.ResponseWriter, r *http.Request, name string, values map[interface{}]interface{}) error {
	session, err := sm.store.Get(r, name)
	if err != nil {
		return err
	}

	for k, v := range values {
		session.Values[k] = v
	}

	return session.Save(r, w)
}

func (sm *SessionManager) DeleteSession(w http.ResponseWriter, r *http.Request, name string) error {
	session, err := sm.store.Get(r, name)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	return session.Save(r, w)
}

func (sm *SessionManager) PutValue(r *http.Request, name string, key interface{}, value interface{}) error {
	session, err := sm.store.Get(r, name)
	if err != nil {
		return err
	}

	session.Values[key] = value

	return session.Save(r, nil)
}

func (sm *SessionManager) GetSession(r *http.Request, name string) (*sessions.Session, error) {
	return sm.store.Get(r, name)
}
