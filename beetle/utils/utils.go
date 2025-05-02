package utils

import (
	"context"
	"fmt"
)

const packageName = "packageName"

func GetRequestID(c context.Context) string {
	fmt.Println(c.Value("X-Request-ID"))
	if val, ok := c.Value("X-Request-ID").(string); ok {
		return val
	}

	return ""
}

func GetFromContext[T any](c context.Context, key string) T {
	if val, ok := c.Value(key).(T); ok {
		return val
	}

	var zero T

	return zero
}
