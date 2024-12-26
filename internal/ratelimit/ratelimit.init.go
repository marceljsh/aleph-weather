package ratelimit

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type (
	RateLimiter interface {
		Allow() bool
		Wait(context.Context) error
	}

	tokenBucketLimiter struct {
		limiter *rate.Limiter
	}
)

func NewTokenBucket(rpm int) RateLimiter {
	return &tokenBucketLimiter{
		limiter: rate.NewLimiter(rate.Every(time.Minute/time.Duration(rpm)), rpm),
	}
}
