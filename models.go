package sj

import "time"

type SearchRequest struct {
	Origin        string            `json:"origin"`
	Destination   string            `json:"destination"`
	DepartureDate string            `json:"departureDate"` // YYYY-MM-DD
	Passengers    []SearchPassenger `json:"passengers"`
}

type SearchPassengerCategory struct {
	Type string `json:"type"`          // "CHILD_AND_YOUTH", "ADULT"
	Age  *int   `json:"age,omitempty"` // only for children
}

type SearchPassenger struct {
	PassengerCategory SearchPassengerCategory `json:"passengerCategory"`
	DiscountCards     []*DiscountCard         `json:"discountCards,omitempty"`
}

type Passenger struct {
	Id                          string            `json:"id"`
	FirstName                   string            `json:"firstName"`
	LastName                    string            `json:"lastName"`
	PassengerCategory           PassengerCategory `json:"passengerCategory"`
	DiscountCards               []*DiscountCard   `json:"discountCards,omitempty"`
	PersonalInformationIsMasked bool              `json:"personalInformationIsMasked"`
}

type DiscountCard struct {
	Code       string `json:"code"`
	Identifier string `json:"identifier"`
}

type TrainStation struct {
	UICStationCode string  `json:"uicStationCode"`
	Name           string  `json:"name"`
	ShortName      *string `json:"shortName,omitempty"`
}

type SearchResponse struct {
	DepartureSearchId      *string `json:"departureSearchId,omitempty"`
	ReturnSearchId         *string `json:"returnDepartureSearchId,omitempty"`
	DepartureSearchExpires *string `json:"departureSearchExpires,omitempty"`
	ReturnSearchExpires    *string `json:"returnDepartureSearchExpires,omitempty"`

	Origin      TrainStation `json:"origin"`
	Destination TrainStation `json:"destination"`

	PassengerListID      string  `json:"passengerListId"`
	PassengerListExpires *string `json:"passengerListExpires,omitempty"`

	Passengers []Passenger `json:"passengers"`
}

type SearchResult struct {
	Travels []Travel `json:"travels"`
}

type Travel struct {
	Origin        TrainStation `json:"origin"`
	Destination   TrainStation `json:"destination"`
	Direction     string       `json:"direction"`     // "OUTBOUND" or "RETURN"
	DepartureDate string       `json:"departureDate"` // YYYY-MM-DD
	Departures    []Departure  `json:"departures"`
}

type Departure struct {
	DepartureID       string       `json:"departureId"`
	Origin            TrainStation `json:"origin"`
	Destination       TrainStation `json:"destination"`
	Direction         string       `json:"direction"`     // "OUTBOUND" or "RETURN"
	DepartureDate     string       `json:"departureDate"` // YYYY-MM-DD
	NumberOfChanges   int          `json:"numberOfChanges"`
	Producer          string       `json:"producer"`
	NumberOfOperators int          `json:"numberOfOperators"`
	ArrivalDateTime   string       `json:"arrivalDateTime"`   // ISO 8601
	DepartureDateTime string       `json:"departureDateTime"` // ISO 8601
	Duration          string       `json:"duration"`          // ISO 8601 duration format PT#H#M
	Resplus           bool         `json:"resplus"`
	Legs              []Leg        `json:"legs"`
}

type Leg struct {
	Origin                     TrainStation `json:"origin"`
	Destination                TrainStation `json:"destination"`
	DepartureDateTime          string       `json:"departureDateTime"`    // ISO 8601
	ArrivalDateTime            string       `json:"arrivalDateTime"`      // ISO 8601
	Duration                   string       `json:"duration"`             // ISO 8601 duration format PT#H#M
	ChangeTime                 string       `json:"changeTime,omitempty"` // ISO 8601 duration format PT#H#M
	ServiceName                string       `json:"serviceName"`
	SecondaryServiceName       *string      `json:"secondaryServiceName,omitempty"`
	PublicServiceName          *string      `json:"publicServiceName,omitempty"`
	ServiceType                ServiceType  `json:"serviceType"`
	ServiceIdentifier          string       `json:"serviceIdentifier"`
	Producer                   string       `json:"producer"`
	ServiceScheduleDate        string       `json:"serviceScheduleDate"` // YYYY-MM-DD
	TransportMethod            string       `json:"transportMethod"`     // "TRAIN", other things prob
	TransportMethodDescription *string      `json:"transportMethodDescription,omitempty"`
	Vehicle                    string       `json:"vehicle"`
	SeatMapAvailable           bool         `json:"seatMapAvailable"`
	EarnSJPrioPoints           bool         `json:"earnSjPrioPoints"`
	NightTrain                 bool         `json:"nightTrain"`
	ServiceBrandName           struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"serviceBrandName"`
	ServiceBrandNameDescription *string           `json:"serviceBrandNameDescription,omitempty"`
	ServiceProperties           []ServiceProperty `json:"serviceProperties"`
	// travelInformationMessages is [] so idk
	International bool `json:"international"`
	PathWay       bool `json:"pathWay"`
	// unavailableReasons is [] so idk
	VehicleImage *string `json:"vehicleImage,omitempty"`
}

type ServiceProperty struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ServiceType struct {
	Code               string `json:"code"`
	Name               string `json:"name"`
	Modality           string `json:"modality"`
	BrandCode          string `json:"brandCode"`
	Description        string `json:"description"`
	OperatorName       string `json:"operatorName"`
	RicsCode           string `json:"ricsCode"`
	ExternalReferences struct {
		ResplusOperator          string `json:"RESPLUS_OPERATOR"`
		TransferRuleOperator     string `json:"TRANSFER_RULE_OPERATOR"`
		TTFilterGroup            string `json:"TTFILTERGROUP"`
		TransferRuleServiceGroup string `json:"TRANSFER_RULE_SERVICEGROUP"`
	} `json:"externalReferences"`
}

type OfferResponse struct {
	DepartureID                      string           `json:"departureId"`
	SeatOffers                       SeatOffers       `json:"seatOffers"`
	BedOffers                        BedOffers        `json:"bedOffers"`
	HasThroughfare                   bool             `json:"hasThroughfare"`
	DepartureStatus                  []string         `json:"departureStatus"`
	Availabilities                   []Availability   `json:"availabilities"`
	PromoCodesSentInRequest          []string         `json:"promoCodesSentInRequest"`
	PriceFrom                        PriceDetail      `json:"priceFrom"`
	PointPriceFrom                   PointPriceDetail `json:"pointPriceFrom"`
	Available                        bool             `json:"available"`
	HasContractPrice                 bool             `json:"hasContractPrice"`
	HasPromotionPrice                bool             `json:"hasPromotionPrice"`
	HasPromotionNonDiscount          bool             `json:"hasPromotionNonDiscount"`
	HasCampaignPrice                 bool             `json:"hasCampaignPrice"`
	HasPossiblePromotionPriceOnAddon bool             `json:"hasPossiblePromotionPriceOnAddon"`
	HasFamilyDiscountPrice           bool             `json:"hasFamilyDiscountPrice"`
	HasCompartmentPrice              bool             `json:"hasCompartmentPrice"`
	HasLastMinutePrice               bool             `json:"hasLastMinutePrice"`
	HasInterRailPrice                bool             `json:"hasInterRailPrice"`
	HasPersonalAssistantPrice        bool             `json:"hasPersonalAssistantPrice"`
	HasAccompanyingPersonPrice       bool             `json:"hasAccompanyingPersonPrice"`
	HasServiceDogPrice               bool             `json:"hasServiceDogPrice"`
	BookingsInConflict               []any            `json:"bookingsInConflictForDoubleBooking"`
}

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type JourneyPrices struct {
	Price                            Money  `json:"price"`
	OriginalPrice                    Money  `json:"originalPrice"`
	TotalPriceToBePaid               Money  `json:"totalPriceToBePaid"`
	RefundAmount                     *Money `json:"refundAmount"`
	FeeAmount                        *Money `json:"feeAmount"`
	RepriceAmount                    *Money `json:"repriceAmount"`
	Corporate                        bool   `json:"corporate"`
	HasTravelPassDiscount            bool   `json:"hasTravelPassDiscount"`
	HasEmployeeTravelPassDiscount    bool   `json:"hasEmployeeTravelPassDiscount"`
	HasPromotionDiscount             bool   `json:"hasPromotionDiscount"`
	HasPromotionNonDiscount          bool   `json:"hasPromotionNonDiscount"`
	HasCampaignDiscount              bool   `json:"hasCampaignDiscount"`
	HasPossiblePromotionPriceOnAddon bool   `json:"hasPossiblePromotionPriceOnAddon"`
	HasPersonalAssistantDiscount     bool   `json:"hasPersonalAssistantDiscount"`
	HasAccompanyingPersonDiscount    bool   `json:"hasAccompanyingPersonDiscount"`
	HasServiceDogDiscount            bool   `json:"hasServiceDogDiscount"`
	HasFamilyDiscount                bool   `json:"hasFamilyDiscount"`
	HasLastMinute                    bool   `json:"hasLastMinute"`
	HasInterRail                     bool   `json:"hasInterRail"`
	HasRepricedFreePrice             bool   `json:"hasRepricedFreePrice"`
}

type PointAmount struct {
	Amount string `json:"amount"`
}

type JourneyPointPrices struct {
	PointPrice                       PointAmount `json:"pointPrice"`
	OriginalPointPrice               PointAmount `json:"originalPointPrice"`
	Price                            *Money      `json:"price"`
	OriginalPrice                    *Money      `json:"originalPrice"`
	Corporate                        bool        `json:"corporate"`
	HasTravelPassDiscount            bool        `json:"hasTravelPassDiscount"`
	HasEmployeeTravelPassDiscount    bool        `json:"hasEmployeeTravelPassDiscount"`
	HasPromotionDiscount             bool        `json:"hasPromotionDiscount"`
	HasPromotionNonDiscount          bool        `json:"hasPromotionNonDiscount"`
	HasCampaignDiscount              bool        `json:"hasCampaignDiscount"`
	HasPossiblePromotionPriceOnAddon bool        `json:"hasPossiblePromotionPriceOnAddon"`
	HasFamilyDiscount                bool        `json:"hasFamilyDiscount"`
	HasPersonalAssistantDiscount     bool        `json:"hasPersonalAssistantDiscount"`
	HasAccompanyingPersonDiscount    bool        `json:"hasAccompanyingPersonDiscount"`
	HasServiceDogDiscount            bool        `json:"hasServiceDogDiscount"`
	HasLastMinute                    bool        `json:"hasLastMinute"`
	HasInterRail                     bool        `json:"hasInterRail"`
}

type PriceDetail struct {
	Price                            string  `json:"price"`
	Currency                         string  `json:"currency"`
	FareClass                        string  `json:"fareClass"`
	OriginalPrice                    string  `json:"originalPrice"`
	TotalPriceToBePaid               string  `json:"totalPriceToBePaid"`
	RefundAmount                     *string `json:"refundAmount"`
	FeeAmount                        *string `json:"feeAmount"`
	RepriceAmount                    *string `json:"repriceAmount"`
	Flexibility                      string  `json:"flexibility"`
	Contract                         bool    `json:"contract"`
	HasTravelPassDiscount            bool    `json:"hasTravelPassDiscount"`
	HasEmployeeTravelPassDiscount    bool    `json:"hasEmployeeTravelPassDiscount"`
	HasPromotionDiscount             bool    `json:"hasPromotionDiscount"`
	HasPromotionNonDiscount          bool    `json:"hasPromotionNonDiscount"`
	HasCampaignDiscount              bool    `json:"hasCampaignDiscount"`
	HasPossiblePromotionPriceOnAddon bool    `json:"hasPossiblePromotionPriceOnAddon"`
	HasFamilyDiscount                bool    `json:"hasFamilyDiscount"`
	HasPersonalAssistantDiscount     bool    `json:"hasPersonalAssistantDiscount"`
	HasAccompanyingPersonDiscount    bool    `json:"hasAccompanyingPersonDiscount"`
	HasServiceDogDiscount            bool    `json:"hasServiceDogDiscount"`
	DisruptionRebook                 bool    `json:"disruptionRebook"`
	HasLastMinute                    bool    `json:"hasLastMinute"`
	HasInterRail                     bool    `json:"hasInterRail"`
	HasRepricedFreePrice             bool    `json:"hasRepricedFreePrice"`
}

type PointPriceDetail struct {
	Price                            *string `json:"price"`
	Currency                         string  `json:"currency"`
	PointPrice                       string  `json:"pointPrice"`
	FareClass                        string  `json:"fareClass"`
	Flexibility                      string  `json:"flexibility"`
	Contract                         bool    `json:"contract"`
	HasTravelPassDiscount            bool    `json:"hasTravelPassDiscount"`
	HasEmployeeTravelPassDiscount    bool    `json:"hasEmployeeTravelPassDiscount"`
	HasPromotionDiscount             bool    `json:"hasPromotionDiscount"`
	HasPromotionNonDiscount          bool    `json:"hasPromotionNonDiscount"`
	HasCampaignDiscount              bool    `json:"hasCampaignDiscount"`
	HasPossiblePromotionPriceOnAddon bool    `json:"hasPossiblePromotionPriceOnAddon"`
	DisruptionRebook                 bool    `json:"disruptionRebook"`
	HasFamilyDiscount                bool    `json:"hasFamilyDiscount"`
	HasLastMinute                    bool    `json:"hasLastMinute"`
	HasInterRail                     bool    `json:"hasInterRail"`
	HasPersonalAssistantDiscount     bool    `json:"hasPersonalAssistantDiscount"`
	HasAccompanyingPersonDiscount    bool    `json:"hasAccompanyingPersonDiscount"`
	HasServiceDogDiscount            bool    `json:"hasServiceDogDiscount"`
}

type SeatOffers struct {
	Offers                           SeatCategories    `json:"offers"`
	PriceFrom                        *PriceDetail      `json:"priceFrom"`
	PointPriceFrom                   *PointPriceDetail `json:"pointPriceFrom"`
	Available                        bool              `json:"available"`
	HasChangedDepartureForFree       bool              `json:"hasChangedDepartureForFree"`
	HasContractPrice                 bool              `json:"hasContractPrice"`
	HasPromotionPrice                bool              `json:"hasPromotionPrice"`
	HasPromotionNonDiscount          bool              `json:"hasPromotionNonDiscount"`
	HasCampaignPrice                 bool              `json:"hasCampaignPrice"`
	HasPossiblePromotionPriceOnAddon bool              `json:"hasPossiblePromotionPriceOnAddon"`
	HasFamilyDiscountPrice           bool              `json:"hasFamilyDiscountPrice"`
	HasInterRailPrice                bool              `json:"hasInterRailPrice"`
	HasLastMinutePrice               bool              `json:"hasLastMinutePrice"`
	HasDisruptionRebookPrice         bool              `json:"hasDisruptionRebookPrice"`
	HasPersonalAssistantPrice        bool              `json:"hasPersonalAssistantPrice"`
	HasAccompanyingPersonPrice       bool              `json:"hasAccompanyingPersonPrice"`
	HasServiceDogPrice               bool              `json:"hasServiceDogPrice"`
}

type SeatCategories struct {
	Second     SeatOffer `json:"SECOND"`
	SecondCalm SeatOffer `json:"SECOND_CALM"`
	First      SeatOffer `json:"FIRST"`
}

type SeatOffer struct {
	GenderTypeRequired               bool                   `json:"genderTypeRequired"`
	Flexibilities                    FlexibilitiesContainer `json:"flexibilities"`
	PriceFrom                        *PriceDetail           `json:"priceFrom"`
	PointPriceFrom                   *PointPriceDetail      `json:"pointPriceFrom"`
	SalesCategoryUsps                []USPCard              `json:"salesCategoryUsps"`
	CompartmentInformationType       *string                `json:"compartmentInformationType"`
	DisruptionRebook                 bool                   `json:"disruptionRebook"`
	Available                        bool                   `json:"available"`
	HasContractPrice                 bool                   `json:"hasContractPrice"`
	HasPromotionPrice                bool                   `json:"hasPromotionPrice"`
	HasPromotionNonDiscount          bool                   `json:"hasPromotionNonDiscount"`
	HasCampaignPrice                 bool                   `json:"hasCampaignPrice"`
	HasPossiblePromotionPriceOnAddon bool                   `json:"hasPossiblePromotionPriceOnAddon"`
	HasFamilyDiscountPrice           bool                   `json:"hasFamilyDiscountPrice"`
	HasInterRailPrice                bool                   `json:"hasInterRailPrice"`
	HasLastMinutePrice               bool                   `json:"hasLastMinutePrice"`
	HasPersonalAssistantPrice        bool                   `json:"hasPersonalAssistantPrice"`
	HasAccompanyingPersonPrice       bool                   `json:"hasAccompanyingPersonPrice"`
	HasServiceDogPrice               bool                   `json:"hasServiceDogPrice"`
	HasChangedDepartureForFree       bool                   `json:"hasChangedDepartureForFree"`
}

type BedOffers struct {
	Offers                           BedCategories    `json:"offers"`
	PriceFrom                        PriceDetail      `json:"priceFrom"`
	PointPriceFrom                   PointPriceDetail `json:"pointPriceFrom"`
	Available                        bool             `json:"available"`
	HasChangedDepartureForFree       bool             `json:"hasChangedDepartureForFree"`
	HasDisruptionRebookPrice         bool             `json:"hasDisruptionRebookPrice"`
	HasContractPrice                 bool             `json:"hasContractPrice"`
	HasPromotionPrice                bool             `json:"hasPromotionPrice"`
	HasPromotionNonDiscount          bool             `json:"hasPromotionNonDiscount"`
	HasCampaignPrice                 bool             `json:"hasCampaignPrice"`
	HasPossiblePromotionPriceOnAddon bool             `json:"hasPossiblePromotionPriceOnAddon"`
	HasFamilyDiscountPrice           bool             `json:"hasFamilyDiscountPrice"`
	HasInterRailPrice                bool             `json:"hasInterRailPrice"`
	HasLastMinutePrice               bool             `json:"hasLastMinutePrice"`
	HasPersonalAssistantPrice        bool             `json:"hasPersonalAssistantPrice"`
	HasAccompanyingPersonPrice       bool             `json:"hasAccompanyingPersonPrice"`
	HasServiceDogPrice               bool             `json:"hasServiceDogPrice"`
}

type BedCategories struct {
	Couchette CouchetteComforts `json:"COUCHETTE"`
	Sleeper   SleeperComforts   `json:"SLEEPER"`
}

type CouchetteComforts struct {
	ComfortTypes                     CouchetteComfortTypes `json:"comfortTypes"`
	PriceFrom                        PriceDetail           `json:"priceFrom"`
	PointPriceFrom                   PointPriceDetail      `json:"pointPriceFrom"`
	Available                        bool                  `json:"available"`
	HasContractPrice                 bool                  `json:"hasContractPrice"`
	HasPromotionPrice                bool                  `json:"hasPromotionPrice"`
	HasPromotionNonDiscount          bool                  `json:"hasPromotionNonDiscount"`
	HasCampaignPrice                 bool                  `json:"hasCampaignPrice"`
	HasPossiblePromotionPriceOnAddon bool                  `json:"hasPossiblePromotionPriceOnAddon"`
	HasDisruptionRebookPrice         bool                  `json:"hasDisruptionRebookPrice"`
	HasFamilyDiscountPrice           bool                  `json:"hasFamilyDiscountPrice"`
	HasInterRailPrice                bool                  `json:"hasInterRailPrice"`
	HasLastMinutePrice               bool                  `json:"hasLastMinutePrice"`
	HasPersonalAssistantPrice        bool                  `json:"hasPersonalAssistantPrice"`
	HasAccompanyingPersonPrice       bool                  `json:"hasAccompanyingPersonPrice"`
	HasServiceDogPrice               bool                  `json:"hasServiceDogPrice"`
	HasChangedDepartureForFree       bool                  `json:"hasChangedDepartureForFree"`
}

type CouchetteComfortTypes struct {
	CouchettePrivate BedComfortDetail `json:"COUCHETTE_PRIVATE"`
	CouchetteShared  BedComfortDetail `json:"COUCHETTE_SHARED"`
}

type SleeperComforts struct {
	ComfortTypes                     SleeperComfortTypes `json:"comfortTypes"`
	PriceFrom                        *PriceDetail        `json:"priceFrom"`
	PointPriceFrom                   *PointPriceDetail   `json:"pointPriceFrom"`
	Available                        bool                `json:"available"`
	HasContractPrice                 bool                `json:"hasContractPrice"`
	HasPromotionPrice                bool                `json:"hasPromotionPrice"`
	HasPromotionNonDiscount          bool                `json:"hasPromotionNonDiscount"`
	HasCampaignPrice                 bool                `json:"hasCampaignPrice"`
	HasPossiblePromotionPriceOnAddon bool                `json:"hasPossiblePromotionPriceOnAddon"`
	HasDisruptionRebookPrice         bool                `json:"hasDisruptionRebookPrice"`
	HasFamilyDiscountPrice           bool                `json:"hasFamilyDiscountPrice"`
	HasInterRailPrice                bool                `json:"hasInterRailPrice"`
	HasLastMinutePrice               bool                `json:"hasLastMinutePrice"`
	HasPersonalAssistantPrice        bool                `json:"hasPersonalAssistantPrice"`
	HasAccompanyingPersonPrice       bool                `json:"hasAccompanyingPersonPrice"`
	HasServiceDogPrice               bool                `json:"hasServiceDogPrice"`
	HasChangedDepartureForFree       bool                `json:"hasChangedDepartureForFree"`
}

type SleeperComfortTypes struct {
	SleeperSecondShared     BedComfortDetail `json:"SLEEPER_SECOND_SHARED"`
	SleeperSecondPrivate    BedComfortDetail `json:"SLEEPER_SECOND_PRIVATE"`
	SleeperFirstPrivate     BedComfortDetail `json:"SLEEPER_FIRST_PRIVATE"`
	SleeperFirstPrivateSolo BedComfortDetail `json:"SLEEPER_FIRST_PRIVATE_SOLO"`
}

type BedComfortDetail struct {
	GenderTypeRequired               bool                   `json:"genderTypeRequired"`
	Flexibilities                    FlexibilitiesContainer `json:"flexibilities"`
	PriceFrom                        *PriceDetail           `json:"priceFrom"`
	PointPriceFrom                   *PointPriceDetail      `json:"pointPriceFrom"`
	SalesCategoryUsps                []USPCard              `json:"salesCategoryUsps"`
	CompartmentInformationType       string                 `json:"compartmentInformationType"`
	DisruptionRebook                 bool                   `json:"disruptionRebook"`
	Available                        bool                   `json:"available"`
	HasContractPrice                 bool                   `json:"hasContractPrice"`
	HasPromotionPrice                bool                   `json:"hasPromotionPrice"`
	HasPromotionNonDiscount          bool                   `json:"hasPromotionNonDiscount"`
	HasCampaignPrice                 bool                   `json:"hasCampaignPrice"`
	HasPossiblePromotionPriceOnAddon bool                   `json:"hasPossiblePromotionPriceOnAddon"`
	HasFamilyDiscountPrice           bool                   `json:"hasFamilyDiscountPrice"`
	HasInterRailPrice                bool                   `json:"hasInterRailPrice"`
	HasLastMinutePrice               bool                   `json:"hasLastMinutePrice"`
	HasPersonalAssistantPrice        bool                   `json:"hasPersonalAssistantPrice"`
	HasAccompanyingPersonPrice       bool                   `json:"hasAccompanyingPersonPrice"`
	HasServiceDogPrice               bool                   `json:"hasServiceDogPrice"`
	HasChangedDepartureForFree       bool                   `json:"hasChangedDepartureForFree"`
}

type FlexibilitiesContainer struct {
	FullFlex FlexibilityDetail `json:"FULLFLEX"`
	NoFlex   FlexibilityDetail `json:"NOFLEX"`
	SemiFlex FlexibilityDetail `json:"SEMIFLEX"`
}

type FlexibilityDetail struct {
	OfferID            *string             `json:"offerId"`
	Expires            *time.Time          `json:"expires"`
	Available          bool                `json:"available"`
	JourneyPrices      *JourneyPrices      `json:"journeyPrices"`
	JourneyPointPrices *JourneyPointPrices `json:"journeyPointPrices"`
	ContainsBedProduct bool                `json:"containsBedProduct"`
	OfferInformation   []OfferInfo         `json:"offerInformation"`
	SalesCategoryUsps  []USPCard           `json:"salesCategoryUsps"`
	CompartmentSize    *string             `json:"compartmentSize"`
	DisruptionRebook   bool                `json:"disruptionRebook"`
	MixedFlex          bool                `json:"mixedFlex"`
	MixedComfort       bool                `json:"mixedComfort"`
	MixedComfortType   bool                `json:"mixedComfortType"`
	MixedOperator      bool                `json:"mixedOperator"`
}

type OfferInfo struct {
	DepartureStation         Station          `json:"departureStation"`
	ArrivalStation           Station          `json:"arrivalStation"`
	DepartureDateTime        time.Time        `json:"departureDateTime"`
	ArrivalDateTime          time.Time        `json:"arrivalDateTime"`
	ServiceIdentifier        string           `json:"serviceIdentifier"`
	ServiceName              string           `json:"serviceName"`
	SecondaryServiceName     *string          `json:"secondaryServiceName"`
	PublicServiceName        string           `json:"publicServiceName"`
	RicsCode                 string           `json:"ricsCode"`
	SalesCategoryComfortType string           `json:"salesCategoryComfortType"`
	SalesCategoryComfort     string           `json:"salesCategoryComfort"`
	SalesCategoryFlexibility string           `json:"salesCategoryFlexibility"`
	GenderTypeRequired       bool             `json:"genderTypeRequired"`
	PassengerOffers          []PassengerOffer `json:"passengerOffers"`
	ProductName              string           `json:"productName"`
	ProductCode              string           `json:"productCode"`
	NightTrain               bool             `json:"nightTrain"`
}

type PassengerOffer struct {
	PassengerID        string       `json:"passengerId"`
	Discounts          []Discount   `json:"discounts"`
	TravelPassDiscount any          `json:"travelPassDiscount"`
	Price              Money        `json:"price"`
	PointPrice         *PointAmount `json:"pointPrice"`
}

type Discount struct {
	Code                    string   `json:"code"`
	Type                    string   `json:"type"`
	Description             string   `json:"description"`
	Amount                  string   `json:"amount"`
	UnlockedByPromotionCode bool     `json:"unlockedByPromotionCode"`
	PromoCodes              []string `json:"promoCodes"`
}

type USPCard struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Availability struct {
	ServiceIdentifier string           `json:"serviceIdentifier"`
	ThroughCarriage   []string         `json:"throughCarriageServiceIdentifiers"`
	AvailableSeats    SeatAvailability `json:"availableSeats"`
	PublicServiceName string           `json:"publicServiceName"`
}

type SeatAvailability struct {
	AvailableForComforts map[string]ComfortStatus `json:"availableForComforts"`
}

type ComfortStatus struct {
	AvailableInOffer bool `json:"availableInOffer"`
	AvailableSeats   bool `json:"availableSeats"`
}

type ConfigResponse struct {
	SalesChannel                  *SalesChannel            `json:"salesChannel,omitempty"`
	Stations                      []Station                `json:"stations"`
	Cards                         []Card                   `json:"cards,omitempty"`
	PassengerCategories           []PassengerCategory      `json:"passengerCategories,omitempty"`
	TravelPassPassengerCategories []PassengerCategory      `json:"travelPassPassengerCategories,omitempty"`
	PrkPassengerCategories        []PassengerCategory      `json:"prkPassengerCategories,omitempty"`
	SpecialNeeds                  []SpecialNeed            `json:"specialNeeds,omitempty"`
	SalesPeriodRestrictions       []SalesPeriodRestriction `json:"salesPeriodRestrictions,omitempty"`
	DiscountCardTypes             []DiscountCardType       `json:"discountCardTypes,omitempty"`
	AdditionalSearchFilters       []AdditionalSearchFilter `json:"additionalSearchFilters,omitempty"`
	Restriction                   any                      `json:"restriction"`
	CustomerPassengerRelations    []string                 `json:"customerPassengerRelations,omitempty"`
}

type SalesChannel struct {
	SalesChannelCode string        `json:"salesChannelCode"`
	GroupSettings    GroupSettings `json:"groupSettings"`
}

type GroupSettings struct {
	BookingGroupStart  int `json:"bookingGroupStart"`
	BookingMaxAdults   int `json:"bookingMaxAdults"`
	BookingMaxChildren int `json:"bookingMaxChildren"`
	BookingMaxPersons  int `json:"bookingMaxPersons"`
}

type Station struct {
	Name           string   `json:"name"`
	UicStationCode string   `json:"uicStationCode"`
	ShortName      string   `json:"shortName"`
	Synonyms       []string `json:"synonyms"`
	SequenceNumber int      `json:"sequenceNumber"`
	Latitude       string   `json:"latitude"`
	Longitude      string   `json:"longitude"`
	Agglomeration  bool     `json:"agglomeration"`
}

type Card struct {
	Type     string `json:"type"`
	Prefix   string `json:"prefix"`
	Length   int    `json:"length"`
	Category string `json:"category"`
	Personal bool   `json:"personal"`
}

type PassengerCategory struct {
	Type    string   `json:"type"`
	AgeSpan *AgeSpan `json:"ageSpan"`
}

type AgeSpan struct {
	AgeFrom int `json:"ageFrom"`
	AgeTo   int `json:"ageTo"`
}

type SpecialNeed struct {
	Code        string `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type SalesPeriodRestriction struct {
	Type                      string `json:"type"`
	NumberOfDays              int    `json:"numberOfDays"`
	SalesRestrictionStartDate string `json:"salesRestrictionStartDate"`
	SalesRestrictionEndDate   string `json:"salesRestrictionEndDate"`
}

type DiscountCardType struct {
	Code string `json:"code"`
}

type AdditionalSearchFilter struct {
	FilterName         string   `json:"filterName"`
	FilterType         string   `json:"filterType"`
	Description        string   `json:"description"`
	PossibleEnumValues []string `json:"possibleEnumValues"`
	DefaultValues      []string `json:"defaultValues"`
	AllowedUsages      []string `json:"allowedUsages"`
}

type TimeTableResponse struct {
	Travels                          []Travel `json:"travels"`
	OnlyShowDeparturesValidForRebook bool     `json:"onlyShowDeparturesValidForRebook"`
}
