package activities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEventShouldSucceed(t *testing.T) {
	eventsService := NewMockNewRelicEventsService()
	ctx := NewNewRelicActivityContextWith(eventsService)
	input := CreateEventInput{
		AccountID:     -1,
		EventDataJson: "{\"eventType\":\"tableName\"}",
	}
	err := ctx.createEventImpl(context.Background(), input)
	require.NoError(t, err)
	require.Equal(t, eventsService.CreateEventWithContextCount, 1)
}

func TestCreateEventShouldErrorOnInvalidClient(t *testing.T) {
	// Set invalid API key and region
	t.Setenv("NEW_RELIC_API_KEY", "")
	t.Setenv("NEW_RELIC_REGION", "US")
	ctx := NewNewRelicActivityContext()
	input := CreateEventInput{
		AccountID:     -1,
		EventDataJson: "{\"eventType\":\"tableName\"}",
	}
	err := ctx.createEventImpl(context.Background(), input)
	require.Error(t, err)
}
