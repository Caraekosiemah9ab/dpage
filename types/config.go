package types

type ServerConfig struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type RouteConfig struct {
	URL      string `json:"url"`
	Plugin   string `json:"plugin"`
	Handler  string `json:"handler"`
	Template string `json:"template"`
}

type NavigationConfig struct {
	URL    string `json:"url"`
	Target string `json:"target,omitempty"`
	Name   string `json:"name"`
}
