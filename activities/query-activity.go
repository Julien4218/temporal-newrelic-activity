package activities

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	"github.com/newrelic/newrelic-client-go/v2/pkg/nrdb"

	"github.com/Julien4218/temporal-newrelic-activity/instrumentation"
)

type QueryNrqlInput struct {
	AccountID int
	Query     string
}

func QueryNrql(ctx context.Context, input QueryNrqlInput) (string, error) {
	instrumentation.Log("QueryNrql")

	client, err := newrelic.New(
		newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")),
		newrelic.ConfigRegion(os.Getenv("NEW_RELIC_REGION")),
	)
	if err != nil {
		message := fmt.Sprintf("error initializing client:%s", err.Error())
		instrumentation.Log(message)
		return "", errors.New(message)
	}
	err = client.TestEndpoints()
	if err != nil {
		message := fmt.Sprintf("error testing client connection:%s", err.Error())
		instrumentation.Log(message)
		return "", errors.New(message)
	}
	instrumentation.Log("NewRelic endpoints are good")
	instrumentation.Log(fmt.Sprintf("Querying on accountID:%d with:%s", input.AccountID, nrdb.NRQL(input.Query)))
	result, err := client.Nrdb.Query(input.AccountID, nrdb.NRQL(input.Query))
	if err != nil {
		message := fmt.Sprintf("error while querying NRQL detail:%s", err.Error())
		instrumentation.Log(message)
		return "", errors.New(message)
	}
	instrumentation.Log(fmt.Sprintf("Got %d current results", len(result.CurrentResults)))
	return "OK", nil
}
