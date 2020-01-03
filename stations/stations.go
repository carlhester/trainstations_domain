package stations

type Station struct {
	Abbr string
	Name string
}

type Repository interface {
	GetAllStations() ([]Station, error)
	Add(Station) error
}
