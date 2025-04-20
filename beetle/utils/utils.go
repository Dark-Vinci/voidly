package utils

import "context"

const packageName = "packageName"

func GetRequestID(c context.Context) string {
	if val := c.Value("requestID"); val != nil {
		return val.(string)
	}

	return ""
}
