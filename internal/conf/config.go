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
	CONNECTALL_RALLY_URL      string `koanf:"CONNECTALL_RALLY_URL"`
	CONNECTALL_RALLY_API_KEY  string `koanf:"CONNECTALL_RALLY_API_KEY"`
	CONNECTALL_WORKSPACE_ID   string `koanf:"CONNECTALL_WORKSPACE_ID"`
	CONNECTALL_CUSTOM_ATTR_ID string `koanf:"CONNECTALL_CUSTOM_ATTR_ID"`
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
		"CONNECTALL_RALLY_URL":      "https://rally1.rallydev.com/slm/webservice/v2.0",
		"CONNECTALL_WORKSPACE_ID":   "836073279859",
		"CONNECTALL_CUSTOM_ATTR_ID": "836082869309",
	}, "."), nil)
}

func loadEnvironment(k *koanf.Koanf) {
	k.Load(env.Provider("", ".", nil), nil)
}

func loadLocalFile(k *koanf.Koanf) {
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
	}
}
