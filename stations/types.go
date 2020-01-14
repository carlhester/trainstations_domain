package stations

type Station struct {
	Abbr string `json:"Abbr"`
	Name string `json:"Name"`
}

type Repository interface {
	Add(Station) error
	Get(string) (Station, error)
	GetAll() ([]Station, error)
}
