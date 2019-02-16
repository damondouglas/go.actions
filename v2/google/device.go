package google

// Device represents information about the device the user is using to interact with the Action.
type Device struct {
	Location *Location
}

// Location represents a location.
type Location struct {
	Coordinates struct {
		Latitude  float64
		Longitude float64
	}
	FormattedAddress string
	ZipCode          string
	City             string
	PostalAddress    struct {
		Revision           int
		RegionCode         string
		LanguageCode       string
		PostalCode         string
		SortingCode        string
		AdministrativeArea string
		Locality           string
		Sublocality        string
		AddressLines       []string
		Recipients         []string
		Organization       string
	}
	Name        string
	PhoneNumber string
	Notes       string
	PlaceID     string
}
