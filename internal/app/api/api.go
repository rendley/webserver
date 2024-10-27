package api

type API struct {
	// unexported field
	config *Config
}

func New(config *Config) *API {
	return &API{
		config: config,
	}
}

func (api *API) Start() error {
	return nil
}
