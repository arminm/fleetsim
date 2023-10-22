package sim

type Config struct {
}
type sim struct {
	config *Config
}

func CreateSim() *sim {
	return &sim{config: &Config{}}
}

func (s *sim) Start() {
	println("Start sim!")
}
