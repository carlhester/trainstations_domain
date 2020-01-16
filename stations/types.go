package stations

type Station struct {
	Abbr string `json:"Abbr"`
	Name string `json:"Name"`
}

type Repository interface {
	Get(string) (Station, error)
	GetAll() ([]Station, error)
}
