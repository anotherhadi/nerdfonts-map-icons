package icons

import (
	"embed"
	"log"

	"gopkg.in/yaml.v3"
)

//go:embed icons.yaml
var iconsFile embed.FS

type Icons struct {
	Icons []struct {
		Name            string   `yaml:"name"`
		DesktopIconName string   `yaml:"desktop_icon_name"`
		Icon            string   `yaml:"icon"`
		Tags            []string `yaml:"tags"`
	}
}

func GetIcons() Icons {
	icons := Icons{}

	content, err := iconsFile.ReadFile("icons.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &icons)
	if err != nil {
		log.Fatal(err)
	}

	return icons
}
