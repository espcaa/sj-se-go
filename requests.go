package sj

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) DoRequest(ctx context.Context, method, path string, body any, response any) error {
	url := fmt.Sprintf("%s%s", c.baseUrl, path)

	fmt.Println("Request URL:", url)

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return fmt.Errorf("failed to encode request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, &buf)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	if response != nil {
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

func (c Client) GetConfig(ctx context.Context) (*ConfigResponse, error) {

	var configResponse ConfigResponse
	err := c.DoRequest(ctx, http.MethodGet, "/config", nil, &configResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	return &configResponse, nil
}

func (c Client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	var searchResponse SearchResponse
	err := c.DoRequest(ctx, http.MethodPost, "/search", request, &searchResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to search journeys: %w", err)
	}

	return &searchResponse, nil
}

func (c Client) GetJourneyOffer(ctx context.Context, departureID string, passengerListID string) (*OfferResponse, error) {
	var journeyOffersResponse OfferResponse
	err := c.DoRequest(ctx, http.MethodGet, fmt.Sprintf("/departures/%s/offers?passengerListId=%s", departureID, passengerListID), nil, &journeyOffersResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to get journey offers: %w", err)
	}

	return &journeyOffersResponse, nil
}

func (c Client) GetOptionsFromDepartureID(ctx context.Context, departureId string) (*TimeTableResponse, error) {
	var timetableResponse TimeTableResponse
	err := c.DoRequest(ctx, http.MethodGet, fmt.Sprintf("/departures/search/%s", departureId), nil, &timetableResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to get options from departure ID: %w", err)
	}

	return &timetableResponse, nil
}
