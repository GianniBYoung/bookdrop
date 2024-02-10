package configa

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	defaultSender   string `yaml:"defaultSender"`
	defaultReciever string `yaml:"defaultReciever"`
	apikey          string `yaml:"apikey"`
	DebugMode       bool   `yaml:"debugMode"`
}

var config Config
var ConfigPath string
var ConfigDirPath string

func SurveyUser() {
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[bool]().Title("Enable debug mode?").Options(
			huh.NewOption("Yes", true),
			huh.NewOption("No", false),
		).Value(&config.DebugMode),
	),
		huh.NewGroup(
			huh.NewSelect[string]().Title("Who is the sender?").Options(
				huh.NewOption("us.nelnet.biz", ".us.nelnet.biz"),
				huh.NewOption("glhec.org", ".glhec.org"),
				huh.NewOption("nulsc.biz", ".nulsc.biz"),
			).Value(&config.defaultSender),
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
func GenerateConfig() error {
	conf, err := yaml.Marshal(&config)

	if err != nil {
		log.Error("", err, "YAML Marshal Err")
	}

	err = os.WriteFile(ConfigPath, conf, 0644)

	if err != nil {
		log.Fatal("", err, "Error Writing Config File to "+ConfigPath)
		return err
	}

	return nil
}

func ReadConfig() Config {
	f, err := os.ReadFile(ConfigPath)

	if err != nil {
		log.Fatal(err)
	}

	var c Config

	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", c)

	return c
}

func Configure() {
	log.SetReportTimestamp(false)
	log.SetReportCaller(false)

	GetConfigPath()
	SurveyUser()
	e := GenerateConfig()

	if e != nil {
		log.Error("", e, "config file error")
	}

	log.Print(config)
	log.Print("Config Generated!")
}
