package configa

import (
	"errors"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type Configuration struct {
	DefaultSender   string `yaml:"defaultSender"`
	DefaultReciever string `yaml:"defaultReciever"`
	Apikey          string `yaml:"apikey"`
	DebugMode       bool   `yaml:"debugMode"`
}

var Config Configuration
var ConfigPath string
var ConfigDirPath string

func SurveyUser() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Enable debug mode?").Value(&Config.DebugMode),
			huh.NewInput().Title("Who is the sender?").Value(&Config.DefaultSender).
				Validate(func(str string) error {
					// match regex for email
					if str == "Frank" {
						return errors.New("Sorry, we don’t serve customers named Frank.")
					}
					return nil
				}),
			huh.NewInput().Title("What is the device's email?").Value(&Config.DefaultReciever).
				Validate(func(str string) error {
					// match regex for email
					if str == "Frank" {
						return errors.New("Sorry, we don’t serve customers named Frank.")
					}
					return nil
				}),
		),
	)
	form.Run()
}

func GetConfigPath() {

	ConfigDirPath = os.Getenv("XDG_CONFIG_HOME")

	if ConfigDirPath == "" {
		homeDir, _ := os.UserHomeDir()
		ConfigDirPath = homeDir + "/.config"
		err := os.MkdirAll(homeDir+".config", 0755)
		if err != nil {
			log.Error("Unable to create config path", err, "Path Error:")
		}
	}

	ConfigPath = ConfigDirPath + "/bookdrop.yaml"

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

func ReadConfig() {
	f, err := os.ReadFile(ConfigPath)

	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(f, &Config); err != nil {
		log.Fatal(err)
	}

	log.Debug("%+v\n", Config)
}

func Configure() {
	log.SetReportTimestamp(false)
	log.SetReportCaller(false)

	GetConfigPath()
	SurveyUser()
	e := generateConfig()

	if e != nil {
		log.Error("", e, "config file error")
	}

	log.Debug("Config Generated!")
}
