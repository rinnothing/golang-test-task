package test

import (
	"log"
	"math/rand/v2"
	"net/http"
	"slices"
	"strings"
	"testing"

	"github.com/rinnothing/golang-test-task/config"
	"github.com/stretchr/testify/require"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.New("../config/prod.yaml")
	if err != nil {
		log.Fatalf("can't read config: %s", err.Error())
	}

	m.Run()
}

func TestPath(t *testing.T) {
	ctx := t.Context()

	client, err := newTestClient(cfg.HTTP.Port)
	require.NoError(t, err)

	// add 100 random numbers
	var vals []int
	for range 100 {
		n := rand.Int() % 1000
		vals = append(vals, n)
		slices.Sort(vals)

		resp, err := client.AddNumber(ctx, n)
		require.NoError(t, err)

		require.Equal(t, http.StatusCreated, resp.StatusCode())
		require.Equal(t, vals, *resp.JSON201)
	}

	// try to put string into
	resp, err := client.client.PostIntegerAddWithBodyWithResponse(ctx, "application/json", strings.NewReader("hello"))
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode())
}
