package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func testServer(t *testing.T, actionCount *atomic.Int64) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	mux.HandleFunc("/jmap/session", func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"apiUrl":          "/jmap",
			"primaryAccounts": map[string]string{"urn:ietf:params:jmap:core": "a"},
		})
	})
	mux.HandleFunc("/jmap", func(w http.ResponseWriter, r *http.Request) {
		var req request
		_ = json.NewDecoder(r.Body).Decode(&req)
		if len(req.MethodCalls) == 1 && req.MethodCalls[0].Name == "x:Action/set" {
			actionCount.Add(1)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"methodResponses": []any{
				[]any{"x:Action/set", map[string]any{
					"created": map[string]any{"new": map[string]any{"id": "x"}},
				}, "c0"},
			},
		})
	})

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	return server
}

func newTestClient(t *testing.T, server *httptest.Server, autoReload bool) *Client {
	t.Helper()

	c, err := New(context.Background(), Config{
		Endpoint:   server.URL,
		Username:   "admin",
		Password:   "secret",
		AutoReload: autoReload,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return c
}

func TestReloadInvokesAction(t *testing.T) {
	t.Parallel()

	var actions atomic.Int64
	c := newTestClient(t, testServer(t, &actions), true)

	if err := c.Reload(context.Background(), "ReloadSettings"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actions.Load() != 1 {
		t.Fatalf("expected 1 action, got %d", actions.Load())
	}
}

func TestReloadSkipsWhenDisabled(t *testing.T) {
	t.Parallel()

	var actions atomic.Int64
	c := newTestClient(t, testServer(t, &actions), false)

	if err := c.Reload(context.Background(), "ReloadSettings"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actions.Load() != 0 {
		t.Fatalf("expected no actions, got %d", actions.Load())
	}
}

func TestReloadSkipsEmptyAction(t *testing.T) {
	t.Parallel()

	var actions atomic.Int64
	c := newTestClient(t, testServer(t, &actions), true)

	if err := c.Reload(context.Background(), ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actions.Load() != 0 {
		t.Fatalf("expected no actions, got %d", actions.Load())
	}
}

func TestReloadCoalescesConcurrentRequests(t *testing.T) {
	t.Parallel()

	var actions atomic.Int64
	c := newTestClient(t, testServer(t, &actions), true)

	requested := time.Now()
	c.reloadMutex.Lock()
	c.lastReload["ReloadSettings"] = requested.Add(time.Second)
	c.reloadMutex.Unlock()

	if err := c.Reload(context.Background(), "ReloadSettings"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actions.Load() != 0 {
		t.Fatalf("expected reload to be skipped, got %d actions", actions.Load())
	}

	c.reloadMutex.Lock()
	c.lastReload["ReloadSettings"] = requested.Add(-time.Second)
	c.reloadMutex.Unlock()

	if err := c.Reload(context.Background(), "ReloadSettings"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actions.Load() != 1 {
		t.Fatalf("expected reload to run, got %d actions", actions.Load())
	}
}
