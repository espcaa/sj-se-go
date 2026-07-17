package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/espcaa/sj-se-go"
	"github.com/sosodev/duration"
)

var baseUrl = "https://prod-api.adp.sj.se/public/sales/booking/v3"
var subscriptionKey = "d6625619def348d38be070027fd24ff6"

func main() {
	client := sj.NewClient(baseUrl, subscriptionKey)

	// context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get the config to get train stations ids
	config, err := client.GetConfig(ctx)
	if err != nil {
		fmt.Printf("Error getting config: %v\n", err)
		return
	}

	// build a station map for easy lookup
	stationMap := make(map[string]string)
	for _, station := range config.Stations {
		stationMap[station.Name] = station.UicStationCode
	}

	// create the passenger object

	age := 16
	passenger := sj.SearchPassenger{
		PassengerCategory: sj.SearchPassengerCategory{
			Type: "CHILD_AND_YOUTH",
			Age:  &age,
		},
		// you need a valid interail code to use this, but that's how you apply discounts!
		//
		//},
	}

	// create the search request
	searchRequest := sj.SearchRequest{
		Origin:        stationMap["Berlin"],
		Destination:   stationMap["Stockholm Central"],
		DepartureDate: "2026-07-26",
		Passengers:    []sj.SearchPassenger{passenger},
	}

	fmt.Printf("Searching for journeys from %s to %s on %s for %d passengers...\n",
		searchRequest.Origin, searchRequest.Destination, searchRequest.DepartureDate, len(searchRequest.Passengers))

	// search for journeys
	searchResponse, err := client.Search(ctx, searchRequest)
	if err != nil {
		fmt.Printf("Error searching journeys: %v\n", err)
		return
	}

	// get the different journey options

	departureID := searchResponse.DepartureSearchId

	optionsResponse, err := client.GetOptionsFromDepartureID(ctx, *departureID)
	if err != nil {
		fmt.Printf("Error getting options: %v\n", err)
		return
	}

	// print the options
	type Option struct {
		DepartureID string
		Departure   string
		Arrival     string
		Duration    duration.Duration
		Price       *string
		SoldOut     bool
	}

	var options []Option

	for _, option := range optionsResponse.Travels {
		for _, departure := range option.Departures {

			duration, err := duration.Parse(departure.Duration)
			if err != nil {
				fmt.Printf("Error parsing duration for departure %s: %v\n", departure.DepartureID, err)
				continue
			}

			option := Option{
				DepartureID: departure.DepartureID,
				Departure:   departure.Origin.Name,
				Arrival:     departure.Destination.Name,
				Duration:    *duration,
				Price:       nil,
				SoldOut:     false,
			}
			options = append(options, option)
		}
	}

	// now get the prices for each option
	for i, option := range options {
		offerResponse, err := client.GetJourneyOffer(ctx, option.DepartureID, searchResponse.PassengerListID)
		if err != nil {
			fmt.Printf("Error getting offers for departure %s: %v\n", option.DepartureID, err)
			continue
		}

		priceFloat, err := strconv.ParseFloat(offerResponse.PriceFrom.Price, 64)
		if err != nil {
			priceFloat = 0.0
		}
		priceFormatted := fmt.Sprintf("%.2f %s", priceFloat, offerResponse.PriceFrom.Currency)

		options[i].Price = &priceFormatted
		options[i].SoldOut = !offerResponse.Available

	}

	// print the options with prices
	for _, option := range options {

		priceStr := "N/A"
		if option.Price != nil {
			priceStr = *option.Price
		}

		fmt.Printf("Departure: %s, Arrival: %s, Duration: %s, Price: %s, Sold Out: %t\n",
			option.Departure, option.Arrival, option.Duration.String(), priceStr, option.SoldOut)
	}
}
