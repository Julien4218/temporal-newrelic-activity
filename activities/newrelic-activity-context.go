package activities

import (
	"fmt"
	"os"

	"github.com/newrelic/newrelic-client-go/v2/pkg/config"
	"github.com/newrelic/newrelic-client-go/v2/pkg/events"

	"github.com/Julien4218/temporal-newrelic-activity/instrumentation"
)

type NewRelicActivityContext struct {
	eventsService events.EventsAPI
}

type NewRelicActivityContextAPI interface {
	GetEventsService() events.EventsAPI
}

func NewNewRelicActivityContext() *NewRelicActivityContext {
	return &NewRelicActivityContext{}
}

func NewNewRelicActivityContextWith(eventsService events.EventsAPI) *NewRelicActivityContext {
	return &NewRelicActivityContext{
		eventsService: eventsService,
	}
}

func (c *NewRelicActivityContext) GetEventsService() (events.EventsAPI, error) {
	if c.eventsService == nil {
		service, err := events.NewEventsService(
			config.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")),
			config.ConfigRegion(os.Getenv("NEW_RELIC_REGION")),
		)
		if err != nil {
			message := fmt.Sprintf("error initializing client:%s", err.Error())
			instrumentation.Log(message)
			return nil, err
		}
		c.eventsService = service
	}
	return c.eventsService, nil
}
