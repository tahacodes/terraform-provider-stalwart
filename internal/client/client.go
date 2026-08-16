package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const CoreCapability = "urn:ietf:params:jmap:core"

type Config struct {
	Endpoint string
	Username string
	Password string
	Token    string
	Insecure bool
	Timeout  time.Duration
}

type Client struct {
	http      *http.Client
	endpoint  string
	apiURL    string
	accountID string
	username  string
	password  string
	token     string
}

type MethodCall struct {
	Name   string
	Args   map[string]any
	CallID string
}

func (m MethodCall) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{m.Name, m.Args, m.CallID})
}

func (m *MethodCall) UnmarshalJSON(b []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if len(raw) != 3 {
		return fmt.Errorf("method call must have 3 elements, got %d", len(raw))
	}
	if err := json.Unmarshal(raw[0], &m.Name); err != nil {
		return err
	}
	if err := json.Unmarshal(raw[1], &m.Args); err != nil {
		return err
	}
	return json.Unmarshal(raw[2], &m.CallID)
}

type request struct {
	Using       []string     `json:"using"`
	MethodCalls []MethodCall `json:"methodCalls"`
}

type response struct {
	MethodResponses []MethodCall `json:"methodResponses"`
	SessionState    string       `json:"sessionState"`
}

type session struct {
	APIURL          string            `json:"apiUrl"`
	PrimaryAccounts map[string]string `json:"primaryAccounts"`
	Accounts        map[string]any    `json:"accounts"`
}

func New(ctx context.Context, cfg Config) (*Client, error) {
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	if cfg.Insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	c := &Client{
		http:     &http.Client{Timeout: timeout, Transport: transport},
		endpoint: strings.TrimRight(cfg.Endpoint, "/"),
		username: cfg.Username,
		password: cfg.Password,
		token:    cfg.Token,
	}

	if err := c.discover(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) AccountID() string {
	return c.accountID
}

func (c *Client) discover(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint+"/jmap/session", nil)
	if err != nil {
		return err
	}
	c.authenticate(req)

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("fetching JMAP session: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("JMAP session returned %d: %s", resp.StatusCode, truncate(body))
	}

	var s session
	if err := json.Unmarshal(body, &s); err != nil {
		return fmt.Errorf("decoding JMAP session: %w", err)
	}

	c.apiURL = s.APIURL
	if c.apiURL == "" {
		c.apiURL = c.endpoint + "/jmap"
	}

	for _, id := range s.PrimaryAccounts {
		c.accountID = id
		break
	}
	if c.accountID == "" {
		for id := range s.Accounts {
			c.accountID = id
			break
		}
	}
	if c.accountID == "" {
		return fmt.Errorf("no account found in JMAP session")
	}

	return nil
}

func (c *Client) authenticate(req *http.Request) {
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
		return
	}
	req.SetBasicAuth(c.username, c.password)
}

func (c *Client) Call(ctx context.Context, calls []MethodCall) ([]MethodCall, error) {
	payload, err := json.Marshal(request{Using: []string{CoreCapability}, MethodCalls: calls})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.apiURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	c.authenticate(req)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("calling JMAP API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("JMAP API returned %d: %s", resp.StatusCode, truncate(body))
	}

	var out response
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("decoding JMAP response: %w", err)
	}

	return out.MethodResponses, nil
}

func (c *Client) Get(ctx context.Context, objectType string, ids []string) ([]map[string]any, error) {
	args := map[string]any{"accountId": c.accountID}
	if ids != nil {
		args["ids"] = ids
	}

	responses, err := c.Call(ctx, []MethodCall{{Name: objectType + "/get", Args: args, CallID: "c0"}})
	if err != nil {
		return nil, err
	}
	if len(responses) == 0 {
		return nil, fmt.Errorf("empty response for %s/get", objectType)
	}
	if err := methodError(responses[0]); err != nil {
		return nil, err
	}

	return decodeList(responses[0].Args["list"])
}

func (c *Client) Query(ctx context.Context, objectType string) ([]string, error) {
	responses, err := c.Call(ctx, []MethodCall{{
		Name:   objectType + "/query",
		Args:   map[string]any{"accountId": c.accountID},
		CallID: "c0",
	}})
	if err != nil {
		return nil, err
	}
	if len(responses) == 0 {
		return nil, fmt.Errorf("empty response for %s/query", objectType)
	}
	if err := methodError(responses[0]); err != nil {
		return nil, err
	}

	raw, _ := responses[0].Args["ids"].([]any)
	ids := make([]string, 0, len(raw))
	for _, v := range raw {
		if s, ok := v.(string); ok {
			ids = append(ids, s)
		}
	}

	return ids, nil
}

func (c *Client) Create(ctx context.Context, objectType string, object map[string]any) (map[string]any, error) {
	const ref = "new"

	args := map[string]any{
		"accountId": c.accountID,
		"create":    map[string]any{ref: object},
	}

	result, err := c.set(ctx, objectType, args)
	if err != nil {
		return nil, err
	}

	if notCreated, ok := result["notCreated"].(map[string]any); ok {
		if detail, exists := notCreated[ref]; exists {
			return nil, fmt.Errorf("%s not created: %s", objectType, encode(detail))
		}
	}

	created, ok := result["created"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s/set returned no created objects", objectType)
	}
	object, ok = created[ref].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s/set created no object for %q", objectType, ref)
	}

	return object, nil
}

func (c *Client) Update(ctx context.Context, objectType, id string, patch map[string]any) error {
	args := map[string]any{
		"accountId": c.accountID,
		"update":    map[string]any{id: patch},
	}

	result, err := c.set(ctx, objectType, args)
	if err != nil {
		return err
	}

	if notUpdated, ok := result["notUpdated"].(map[string]any); ok {
		if detail, exists := notUpdated[id]; exists {
			return fmt.Errorf("%s %q not updated: %s", objectType, id, encode(detail))
		}
	}

	return nil
}

func (c *Client) Destroy(ctx context.Context, objectType, id string) error {
	args := map[string]any{
		"accountId": c.accountID,
		"destroy":   []string{id},
	}

	result, err := c.set(ctx, objectType, args)
	if err != nil {
		return err
	}

	if notDestroyed, ok := result["notDestroyed"].(map[string]any); ok {
		if detail, exists := notDestroyed[id]; exists {
			return fmt.Errorf("%s %q not destroyed: %s", objectType, id, encode(detail))
		}
	}

	return nil
}

func (c *Client) set(ctx context.Context, objectType string, args map[string]any) (map[string]any, error) {
	responses, err := c.Call(ctx, []MethodCall{{Name: objectType + "/set", Args: args, CallID: "c0"}})
	if err != nil {
		return nil, err
	}
	if len(responses) == 0 {
		return nil, fmt.Errorf("empty response for %s/set", objectType)
	}
	if err := methodError(responses[0]); err != nil {
		return nil, err
	}

	return responses[0].Args, nil
}

func methodError(m MethodCall) error {
	if m.Name != "error" {
		return nil
	}

	return fmt.Errorf("JMAP error: %s", encode(m.Args))
}

func decodeList(v any) ([]map[string]any, error) {
	raw, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("expected a list in JMAP response")
	}

	out := make([]map[string]any, 0, len(raw))
	for _, item := range raw {
		object, ok := item.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("expected an object in JMAP list")
		}
		out = append(out, object)
	}

	return out, nil
}

func encode(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%v", v)
	}

	return string(b)
}

func truncate(b []byte) string {
	const limit = 512
	if len(b) > limit {
		return string(b[:limit]) + "..."
	}

	return string(b)
}
