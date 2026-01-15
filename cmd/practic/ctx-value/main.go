package main

import (
	"context"
	"fmt"
)

type requestID string

const RequestIDKey = requestID("request_id")

func main() {
	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, RequestIDKey, "123-ABC")

	if userEmail, ok := ctxWithValue.Value(RequestIDKey).(string); ok {
		fmt.Println(userEmail)
	} else {
		fmt.Println("no value")
	}
}
