package conf

import (
	"log"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	RALLY_URL     string `koanf:"RALLY_URL"`
	RALLY_API_KEY string `koanf:"RALLY_API_KEY"`
	WORKSPACE_ID  string `koanf:"WORKSPACE_ID"`
}

func Load() *Config {
	k := koanf.New(".")

	loadDefaults(k)
	loadEnvironment(k)
	loadLocalFile(k)

	var cfg Config
	err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "koanf"})
	if err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	return &cfg
}

func loadDefaults(k *koanf.Koanf) {
	k.Load(confmap.Provider(map[string]any{
		"RALLY_URL": "https://rally1.rallydev.com/slm/webservice/v2.0",
		"WORKSPACE_ID": "41529001",
	}, "."), nil)
}

func loadEnvironment(k *koanf.Koanf) {
	k.Load(env.Provider("", ".", nil), nil)
}

func loadLocalFile(k *koanf.Koanf) {
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
	}
}
