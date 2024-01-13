package store

import (
    "net/http"

    "github.com/gorilla/sessions"
)

type SessionManager struct {
    store sessions.Store
}

func NewSessionManager(secretKey string) *SessionManager {
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

func (sm *SessionManager) GetSession(r *http.Request, name string) (*sessions.Session, error) {
	return sm.store.Get(r, name)
}
