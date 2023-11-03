package activities

import (
	"context"

	"github.com/Julien4218/temporal-newrelic-activity/instrumentation"
	newrelic "github.com/newrelic/newrelic-client-go/v2"
)

func QueryNrql(ctx context.Context, param string) (*string, error) {
	instrumentation.Log("QueryNrql")
	_ = newrelic.New([]newrelic.ConfigOption{})
	result := "pass"
	return &result, nil
}
