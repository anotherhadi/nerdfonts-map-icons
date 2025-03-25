package icons

import (
	"embed"
	"errors"
	"log"

	"gopkg.in/yaml.v3"
)

//go:embed icons.yaml
var iconsFile embed.FS

type Icon struct {
	Name            string   `yaml:"name"`
	DesktopIconName string   `yaml:"desktop_icon_name"`
	Icon            string   `yaml:"icon"`
	Tags            []string `yaml:"tags"`
}

type Icons []Icon

type out struct {
	Icons Icons
}

// AllIcons contains all the icons from the icons.yaml file
var AllIcons = GetAllIcons()

// GetAllIcons returns all the icons from the icons.yaml file
func GetAllIcons() Icons {
	icons := out{}

	content, err := iconsFile.ReadFile("icons.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &icons)
	if err != nil {
		log.Fatal(err)
	}

	return icons.Icons
}

// GetIconByName returns the icon with the given name
func GetIconByName(name string) (Icon, error) {
	for _, icon := range AllIcons {
		if icon.Name == name {
			return icon, nil
		}
	}
	return Icon{}, errors.New("Icon not found")
}

// GetIconByDesktopIconName returns the icon with the given desktop icon name
func GetIconByDesktopIconName(name string) (Icon, error) {
	for _, icon := range AllIcons {
		if icon.DesktopIconName == name {
			return icon, nil
		}
	}
	return Icon{}, errors.New("Icon not found")
}

// GetIconByTag returns the icons with the given tag
func GetIconByTag(tag string) (Icons, error) {
	var icons Icons
	for _, icon := range AllIcons {
		for _, t := range icon.Tags {
			if t == tag {
				icons = append(icons, icon)
			}
		}
	}
	if len(icons) == 0 {
		return icons, errors.New("Icon not found")
	}
	return icons, nil
}
