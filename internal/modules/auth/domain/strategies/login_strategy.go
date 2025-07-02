package strategies

import "context"

type LoginStrategy interface {
	Login(ctx context.Context, userID string, ipAddress string) error
}
