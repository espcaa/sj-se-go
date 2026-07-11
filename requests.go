package sj

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetConfig(ctx context.Context) (*ConfigResponse, error) {
	url := fmt.Sprintf("%s/config", c.baseUrl)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	// subscription key
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var configResponse ConfigResponse
	if err := json.NewDecoder(resp.Body).Decode(&configResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &configResponse, nil
}

func (c Client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	url := fmt.Sprintf("%s/search", c.baseUrl)

	// encode the request body as json
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)
	req.Header.Set("ocp-apim-trace", "true")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var searchResponse SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &searchResponse, nil
}

func (c Client) GetJourneyOffer(ctx context.Context, departureID string, passengerListID string) (*OfferResponse, error) {
	url := fmt.Sprintf("%s/departures/%s/offers?passengerListId=%s", c.baseUrl, departureID, passengerListID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var journeyOffersResponse OfferResponse
	if err := json.NewDecoder(resp.Body).Decode(&journeyOffersResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &journeyOffersResponse, nil
}

func (c Client) GetOptionsFromDepartureID(ctx context.Context, departureId string) (*TimeTableResponse, error) {
	url := fmt.Sprintf("%s/departures/search/%s", c.baseUrl, departureId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var timetableResponse TimeTableResponse
	if err := json.NewDecoder(resp.Body).Decode(&timetableResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &timetableResponse, nil
}
