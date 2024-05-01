package configa

import (
	"os"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type Configuration struct {
	DefaultSender   string `yaml:"defaultSender"`
	DefaultReciever string `yaml:"defaultReciever"`
	ApiKey          string `yaml:"apiKey,omitempty"`
	DebugMode       bool   `yaml:"debugMode"`
}

var Config Configuration
var envApiKey string

func SurveyUser() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Enable debug mode?").Value(&Config.DebugMode),
			huh.NewInput().Title("Who is the sender?").Value(&Config.DefaultSender),
			huh.NewInput().Title("What is the device's email?").Value(&Config.DefaultReciever),
		),
	)
	form.Run()
}

// This function should always generate and overwrite a config.
func generateConfig(path string) error {
	Config, err := yaml.Marshal(&Config)

	if err != nil {
		log.Error("", err, "YAML Marshal Err")
	}

	err = os.WriteFile(path, Config, 0644)
	log.Error(path)

	if err != nil {
		log.Error("", err, "Error Writing Config File to "+path)
		return err
	}

	return nil
}

func readConfig() {

	configPath, err := xdg.SearchConfigFile("bookdrop/config.yaml")
	if err != nil {
		log.Errorf("No config file found: %s", err)
	}
	log.Debugf("Config file was found at: %s", configPath)

	f, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatal("", err, "Error retrieving config from path")
	}

	if err := yaml.Unmarshal(f, &Config); err != nil {
		log.Fatal("", err, "Error unmarshalling config")
	}

	if envApiKey != "" {
		Config.ApiKey = envApiKey
	}
}

func Configure() {
	configPath, err := xdg.ConfigFile("bookdrop/config.yaml")
	log.Debug(configPath)

	if err != nil {
		log.Fatal("", err, "Error Getting Config File Path: "+configPath)
	}

	envApiKey = os.Getenv("RESEND_API_KEY")
	Config.ApiKey = envApiKey
	if _, err := os.Stat(configPath); err != nil {
		log.Info("Config file not found, generating!")
		SurveyUser()
		e := generateConfig(configPath)
		if e != nil {
			log.Error("", e, "config file error")
		}

		if envApiKey != "" {
			Config.ApiKey = envApiKey
		}

		log.Debug("Config Generated!")
		log.Debug(Config)
		return
	}

	readConfig()

}
