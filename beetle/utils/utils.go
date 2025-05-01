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
