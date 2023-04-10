package session

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/pkg/v3/infrastucture/redis"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Store struct {
	CookieKey      string
	CookieSameSite string
	CookiePath     string
	CookieDomain   string
	CookieSecure   bool
	CookieHTTPOnly bool
	ExpirationWeb  time.Duration
	ExpirationApp  time.Duration

	redis *redis.RedisClient
}

var (
	ErrSessionNotFound = errors.New("session not found")
)

func NewStore(redis *redis.RedisClient) *Store {
	return &Store{
		CookieKey:      "session_id",
		CookieSameSite: "lax",
		CookiePath:     "/",
		CookieSecure:   false,
		CookieHTTPOnly: true,
		ExpirationWeb:  30 * time.Minute,
		ExpirationApp:  7 * 24 * time.Hour,
		redis:          redis,
	}
}

func (s *Store) GetSession(ctx context.Context, id uuid.UUID) (*Session, error) {
	keys, err := s.redis.Keys(ctx, fmt.Sprintf("_user_session:*:%s", id))
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, ErrSessionNotFound
	}

	result, err := s.redis.Get(ctx, keys[0])
	if err != nil {
		return nil, ErrSessionNotFound
	}

	var sessionData SessionData
	err = result.Scan(&sessionData)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	session := acquireSession()
	session.id = id
	session.ctx = ctx
	session.store = s
	session.Data = sessionData

	return session, nil
}

func (s *Store) GetSessions(ctx context.Context, uid string) ([]SessionData, error) {
	keys, err := s.redis.Keys(ctx, fmt.Sprintf("_user_session:%s:*", uid))
	if err != nil {
		return nil, err
	}

	sessionsData := make([]SessionData, 0)

	for _, key := range keys {
		result, err := s.redis.Get(ctx, key)
		if err != nil {
			return nil, err
		}

		var sessionData SessionData
		err = result.Scan(&sessionData)
		if err != nil {
			return nil, err
		}

		sessionsData = append(sessionsData, sessionData)
	}

	return sessionsData, nil
}

func (s *Store) DeleteSession(ctx context.Context, uid string, sessionID string) error {
	exist, err := s.redis.Exist(ctx, fmt.Sprintf("_user_session:%s:%s", uid, sessionID))
	if err != nil {
		return err
	}

	if !exist {
		return ErrSessionNotFound
	}

	return s.redis.Delete(ctx, fmt.Sprintf("_user_session:%s:%s", uid, sessionID))
}

func (s *Store) ApplySession(ctx context.Context, user *userv1.User) (*Session, error) {
	var err error
	fresh := false

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("session is not ready")
	}

	id := s.getSessionID(md)

	if id == uuid.Nil {
		fresh = true
		id = uuid.New()
		if err != nil {
			return nil, err
		}
	}

	session := acquireSession()
	session.id = id
	session.ctx = ctx
	session.store = s
	session.fresh = fresh

	if !fresh {
		key := fmt.Sprintf("_user_session:%s:%s", user.Uid, id)
		exist, err := s.redis.Exist(ctx, key)
		if err != nil {
			return nil, err
		}

		if exist {
			result, err := s.redis.Get(ctx, key)
			if err != nil {
				return nil, err
			}

			err = result.Scan(&session.Data)
			if err != nil {
				return nil, err
			}
		}
		userID, err := user.Id.ToUUID()
		if err != nil {
			return nil, err
		}

		session.Data.UserID = userID
		session.Data.UID = user.Uid
		session.Data.UserAgent = md.Get("grpcgateway-user-agent")[0]
		session.Data.AuthenticatedAt = time.Now()
	}

	session.Data.ID = id

	header := metadata.Pairs(
		s.CookieKey,
		session.id.String(),
	)

	err = grpc.SendHeader(ctx, header)
	if err != nil {
		return nil, err
	}

	err = session.Save(ctx)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Store) getSessionID(md metadata.MD) uuid.UUID {
	if len(md.Get(s.CookieKey)) == 0 {
		return uuid.Nil
	}

	id, err := uuid.Parse(md.Get("session_id")[0])
	if err != nil {
		return uuid.Nil
	}

	return id
}
