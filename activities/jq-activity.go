package activities

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/itchyny/gojq"

	"github.com/Julien4218/temporal-newrelic-activity/instrumentation"
)

type JqInput struct {
	Json  string
	Query string
}

func JQ(ctx context.Context, input JqInput) ([]string, error) {
	instrumentation.Log(fmt.Sprintf("JQ with query:%s", input.Query))

	query, err := gojq.Parse(input.Query)
	if err != nil {
		message := fmt.Sprintf("error while parsing query detail:%s", err.Error())
		instrumentation.Log(message)
		return []string{}, errors.New(message)
	}

	jsonBytes := []byte(input.Json)

	var obj interface{}
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		message := fmt.Sprintf("error while deserializing input json detail:%s", err.Error())
		instrumentation.Log(message)
		return []string{}, errors.New(message)
	}

	result := []string{}
	iter := query.Run(obj)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			message := fmt.Sprintf("error while iterating next detail:%s", err.Error())
			instrumentation.Log(message)
			return []string{}, errors.New(message)
		}

		vj, err := json.Marshal(v)
		if err != nil {
			message := fmt.Sprintf("error while serializing next iteration value detail:%s", err.Error())
			instrumentation.Log(message)
			return []string{}, errors.New(message)
		}

		result = append(result, string(vj))
	}

	return result, nil
}
