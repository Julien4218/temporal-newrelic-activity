package activities

import (
	"context"
)

func (s *UnitTestSuite) Test_ShouldSelectAllSingleField() {

	input := JqInput{
		Json:  "[{\"id\":\"abc\"},{\"id\":\"def\"},{\"id\":\"123\"}]",
		Query: ".[] | .id",
	}
	results, err := JQ(context.Background(), input)
	s.NoError(err)
	s.Equal(3, len(results))
}

func (s *UnitTestSuite) Test_ShouldSelectFirst() {

	input := JqInput{
		Json:  "[{\"id\":\"abc\"},{\"id\":\"def\"},{\"id\":\"123\"}]",
		Query: "(first(.[])) | .id",
	}
	results, err := JQ(context.Background(), input)
	s.NoError(err)
	s.Equal(1, len(results))
	s.Equal("\"abc\"", results[0])
}
