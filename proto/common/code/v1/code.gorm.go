package codev1

import "time"

func (c CodeORM) IsOutAttempt() bool {
	return c.AttemptCount >= 5
}

func (c CodeORM) IsValidated() bool {
	return c.ValidatedAt.Valid
}

func (c CodeORM) IsExpired() bool {
	if c.IsOutAttempt() || c.IsValidated() {
		return true
	}

	return time.Now().After(*c.ExpiredAt)
}
