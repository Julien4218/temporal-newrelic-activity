package activities

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Julien4218/temporal-newrelic-activity/instrumentation"
)

type CreateEventInput struct {
	AccountID     int
	EventDataJson string
}

func NewCreateEventActivity(ctx context.Context, input CreateEventInput) error {
	return NewNewRelicActivityContext().createEventImpl(ctx, input)
}

func (c *NewRelicActivityContext) createEventImpl(ctx context.Context, input CreateEventInput) error {
	instrumentation.Log("CreateEvent")

	if !strings.Contains(input.EventDataJson, "eventType") {
		message := "error the EventDataJson payload requires the use of an `eventType` field that represents the custom event's type"
		instrumentation.Log(message)
		return errors.New(message)
	}

	eventJsonBytes := []byte(input.EventDataJson)

	var obj interface{}
	err := json.Unmarshal(eventJsonBytes, &obj)
	if err != nil {
		message := fmt.Sprintf("error while deserializing input EventDataJson detail:%s", err.Error())
		instrumentation.Log(message)
		return errors.New(message)
	}

	service, err := c.GetEventsService()
	if err != nil {
		message := fmt.Sprintf("error while getting newrelic Events service detail:%s", err.Error())
		instrumentation.Log(message)
		return errors.New(message)
	}

	err = service.CreateEventWithContext(ctx, input.AccountID, obj)
	if err != nil {
		message := fmt.Sprintf("error while creating event detail:%s", err.Error())
		instrumentation.Log(message)
		return errors.New(message)
	}

	return nil
}
