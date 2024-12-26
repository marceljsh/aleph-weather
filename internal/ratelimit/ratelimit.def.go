package ratelimit

import "context"

func (l *tokenBucketLimiter) Allow() bool {
	return l.limiter.Allow()
}

func (l *tokenBucketLimiter) Wait(ctx context.Context) error {
	return l.limiter.Wait(ctx)
}
