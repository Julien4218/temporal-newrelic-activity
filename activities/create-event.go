package activities

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Julien4218/temporal-newrelic-activity/instrumentation"
	"github.com/newrelic/newrelic-client-go/v2/newrelic"
)

type CreateEventInput struct {
	AccountID     int
	EventDataJson string
}

func CreateEvent(ctx context.Context, input CreateEventInput) error {
	client, err := newrelic.New(
		newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")),
		newrelic.ConfigRegion(os.Getenv("NEW_RELIC_REGION")),
	)
	if err != nil {
		message := fmt.Sprintf("error initializing client:%s", err.Error())
		instrumentation.Log(message)
		return errors.New(message)
	}
	instrumentation.Log("CreateEvent")

	if !strings.Contains(input.EventDataJson, "eventType") {
		message := "error the EventDataJson payload requires the use of an `eventType` field that represents the custom event's type"
		instrumentation.Log(message)
		return errors.New(message)
	}

	eventJsonBytes := []byte(input.EventDataJson)

	var obj interface{}
	err = json.Unmarshal(eventJsonBytes, &obj)
	if err != nil {
		message := fmt.Sprintf("error while deserializing input EventDataJson detail:%s", err.Error())
		instrumentation.Log(message)
		return errors.New(message)
	}

	err = client.Events.CreateEventWithContext(ctx, input.AccountID, obj)
	if err != nil {
		message := fmt.Sprintf("error while creating event detail:%s", err.Error())
		instrumentation.Log(message)
		return errors.New(message)
	}

	return nil
}
