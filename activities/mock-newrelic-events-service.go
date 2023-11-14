package activities

import "context"

type MockNewRelicEventsService struct {
	CreateEventWithContextCount int
	CreateEventCount            int
}

func NewMockNewRelicEventsService() *MockNewRelicEventsService {
	return &MockNewRelicEventsService{}
}

func (c *MockNewRelicEventsService) CreateEventWithContext(ctx context.Context, accountID int, event interface{}) error {
	c.CreateEventWithContextCount++
	return nil
}

func (c *MockNewRelicEventsService) CreateEvent(accountID int, event interface{}) error {
	c.CreateEventCount++
	return nil
}
