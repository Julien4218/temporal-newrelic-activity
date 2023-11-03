package activities

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/newrelic/newrelic-client-go/v2/newrelic"

	"github.com/Julien4218/temporal-newrelic-activity/instrumentation"
)

func QueryNrql(ctx context.Context, param string) (string, error) {
	instrumentation.Log("QueryNrql")
	_, err := newrelic.New(newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")))
	if err != nil {
		message := fmt.Sprintf("error initializing client:%s", err.Error())
		instrumentation.Log(message)
		return "", errors.New(message)
	}
	result := "pass"
	return result, nil
}
