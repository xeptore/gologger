package main

type configType struct {
	Colors     map[string]colorKind `yaml:"colors"`
	Styles     map[string]int       `yaml:"styles"`
	ColorReset int                  `yaml:"color_reset"`
}

type colorKind struct {
	Dark  int `yaml:"dark"`
	Light int `yaml:"light"`
}
