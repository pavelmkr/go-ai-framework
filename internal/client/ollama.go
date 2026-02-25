package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type OllamaClient struct {
	Endpoint string
	Model    string
}

type request struct {
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	Stream  bool   `json:"stream"`
}

type response struct {
	Response string `json:"response"`
}

func New(endpoint, model string) *OllamaClient {
	return &OllamaClient{
		Endpoint: endpoint,
		Model:    model,
	}
}

func (o *OllamaClient) Generate(ctx context.Context, prompt string) (string, error) {
	reqBody, _ := json.Marshal(request{
		Model: o.Model,
		Prompt: prompt,
		Stream: false,
	})

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		o.Endpoint+"api/generate",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var r response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}

	return r.Response, nil
}
