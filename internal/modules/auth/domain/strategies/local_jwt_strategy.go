package strategies

import "context"

type LocalJWTLoginStrategy struct {
}

func (s *LocalJWTLoginStrategy) Login(ctx context.Context, userID string, ipAddress string) error {
	return nil
}
