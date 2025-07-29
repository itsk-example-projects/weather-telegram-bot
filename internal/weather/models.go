package weather

type NominatimPlace struct {
	PlaceID     int          `json:"place_id"`
	Licence     string       `json:"licence"`
	OsmType     string       `json:"osm_type"`
	OsmID       int          `json:"osm_id"`
	Lat         string       `json:"lat"`
	Lon         string       `json:"lon"`
	Class       string       `json:"class"`
	Type        string       `json:"type"`
	PlaceRank   int          `json:"place_rank"`
	Importance  float64      `json:"importance"`
	AddressType string       `json:"addresstype"`
	Name        string       `json:"name"`
	DisplayName string       `json:"display_name"`
	Address     PlaceAddress `json:"address"`
	BoundingBox []string     `json:"boundingbox"`
}

type PlaceAddress struct {
	HouseNumber   string `json:"house_number,omitempty"`
	Road          string `json:"road,omitempty"`
	Suburb        string `json:"suburb,omitempty"`
	City          string `json:"city,omitempty"`
	Municipality  string `json:"municipality,omitempty"`
	State         string `json:"state,omitempty"`
	ISO3166v2lvl4 string `json:"ISO3166-2-lvl4,omitempty"`
	Region        string `json:"region,omitempty"`
	Postcode      string `json:"postcode,omitempty"`
	Country       string `json:"country,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
}
