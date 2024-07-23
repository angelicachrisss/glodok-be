package config

import (
	"io/ioutil"
	"os"
	"strings"

	// "time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	config      *Config
	watchConfig WatchCfg
)

const (
	envDevelopment   = "development"
	// envStaging       = "staging"
	// envProduction    = "production"
	// envProductionCHC = "chc-production"
)

type (
	option struct {
		configFile string
	}

	WatchCfg struct {
		Path string
		Name string
	}
)

// Init ...
func Init(opts ...Option) error {
	opt := &option{
		configFile: getDefaultConfigFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	configSplit := strings.Split(opt.configFile, "/")
	lengthSplit := len(configSplit)
	watchConfig.Path = strings.Join(configSplit[:lengthSplit-1], "/")
	watchConfig.Name = configSplit[lengthSplit-1]

	return yaml.Unmarshal(out, &config)
}

func PrepareWatchPath() {
	viper.SetConfigName(watchConfig.Name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(watchConfig.Path)
}

// Option ...

type Option func(*option)

// WithConfigFile ...
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func getDefaultConfigFile() string {
	configPath := "./files/etc/glodok-be/glodok-be.development.yaml"
	namespace, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	env := string(namespace)
	if os.Getenv("GOPATH") == "" {
		configPath = "files/etc/glodok-be/glodok-be.development.yaml"
	}

	if env != "" {
		switch env {
		// case envStaging:
		// 	time.Sleep(30 * time.Second)
		// 	configPath = "/vault/secrets/database.yaml"
		// case envProduction:
		// 	time.Sleep(30 * time.Second)
		// 	configPath = "/vault/secrets/database.yaml"
		default:
			configPath = "./glodok-be/glodok-be.development.yaml"
		}
	}

	return configPath
}

// Get ...
func Get() *Config {
	return config
}
