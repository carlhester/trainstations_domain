package stations

type Station struct {
	Abbr string
	Name string
}

type Repository interface {
	Add(Station) error
	Get(string) (Station, error)
	GetAll() ([]Station, error)
}
