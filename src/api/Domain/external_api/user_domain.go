package external_api

type User struct {
	ID               int    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	UserType         string      `json:"user_type"`
	Tags             []string    `json:"tags"`
	SiteID           string      `json:"site_id"`

}


