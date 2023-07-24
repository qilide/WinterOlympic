package country

type CountryIdBinder struct {
	CountryId string `json:"country_id"`
}

type CountryBinder struct {
	Year      string `json:"year"`
	CountryId string `json:"country_id"`
}
