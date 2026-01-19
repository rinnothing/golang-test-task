package test

import (
	"context"
	"fmt"

	"github.com/rinnothing/golang-test-task/api/gen"
)

type testClient struct {
	client gen.ClientWithResponsesInterface
}

func newTestClient(port string) (*testClient, error) {
	client, err := gen.NewClientWithResponses(fmt.Sprintf("http://localhost:%s", port))
	if err != nil {
		return nil, err
	}
	return &testClient{client: client}, nil
}

func (c *testClient) AddNumber(ctx context.Context, num int) (*gen.PostIntegerAddResponse, error) {
	return c.client.PostIntegerAddWithResponse(ctx, gen.PostIntegerAddJSONRequestBody(num))
}
