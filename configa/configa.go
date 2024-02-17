package configa

import (
	"os"

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
var ConfigPath string
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

func getConfigPath() {

	configDirPath := os.Getenv("XDG_CONFIG_HOME")

	if configDirPath == "" {
		homeDir, _ := os.UserHomeDir()
		configDirPath = homeDir + "/.config"
		err := os.MkdirAll(homeDir+".config", 0755)
		if err != nil {
			log.Error("Unable to create config path", err, "Path Error:")
		}
	}

	ConfigPath = configDirPath + "/bookdrop.yml"

}

// This function should always generate and overwrite a config.
func generateConfig() error {
	Config, err := yaml.Marshal(&Config)

	if err != nil {
		log.Error("", err, "YAML Marshal Err")
	}

	err = os.WriteFile(ConfigPath, Config, 0644)

	if err != nil {
		log.Fatal("", err, "Error Writing Config File to "+ConfigPath)
		return err
	}

	return nil
}

func readConfig() {
	f, err := os.ReadFile(ConfigPath)

	if err != nil {
		log.Fatal("", err, "Error retrieving config from path")
	}

	if err := yaml.Unmarshal(f, &Config); err != nil {
		log.Fatal("", err, "Error unmarshalling config")
	}

	if envApiKey != "" {
		Config.ApiKey = envApiKey
	}

	log.Debug("%+v\n", Config)
}

func Configure() {
	log.SetReportTimestamp(false)
	log.SetReportCaller(false)

	getConfigPath()

	envApiKey = os.Getenv("RESEND_API_KEY")
	Config.ApiKey = envApiKey
	if _, err := os.Stat(ConfigPath); err != nil {
		log.Info("Config file not found, generating!")
		SurveyUser()
		e := generateConfig()
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
