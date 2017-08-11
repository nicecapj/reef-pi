package controller

import (
	"github.com/reef-pi/reef-pi/controller/ato"
	"github.com/reef-pi/reef-pi/controller/equipments"
	"github.com/reef-pi/reef-pi/controller/lighting"
	"github.com/reef-pi/reef-pi/controller/temperature"
	"github.com/reef-pi/reef-pi/controller/utils"
)

type Config struct {
	EnableGPIO  bool               `yaml:"enable_gpio"`
	EnablePWM   bool               `yaml:"enable_pwm"`
	HighRelay   bool               `yaml:"high_relay"`
	Database    string             `yaml:"database"`
	Equipments  equipments.Config  `yaml:"equipments"`
	Lighting    LightingConfig     `yaml:"lighting"`
	AdafruitIO  utils.AdafruitIO   `yaml:"adafruitio"`
	DevMode     bool               `yaml:"dev_mode"`
	ATO         ato.Config         `yaml:"ato"`
	Temperature temperature.Config `yaml:"temperature"`
	Interface   string             `yaml:"interface"`
	Display     bool               `yaml:"display"`
	Admin       AdminConfig        `yaml:"admin"`
	Name        string             `yaml:"name"`
	Dashboard   DashboardConfig    `yaml:"dashboard"`
	Timers      TimersConfig       `yaml:"timers"`
}

type LightingConfig struct {
	Enable   bool                           `yaml:"enable"`
	Channels map[string]lighting.LEDChannel `yaml:"channels"`
}

type DashboardConfig struct {
	Enable bool `yaml:"enable"`
}
type AdminConfig struct {
	Enable bool `yaml:"enable"`
}
type TimersConfig struct {
	Enable bool `yaml:"enable"`
}

var DefaultConfig = Config{
	Database:   "reef-pi.db",
	EnableGPIO: true,
	Equipments: equipments.Config{
		Outlets: make(map[string]equipments.Outlet),
	},
	Lighting: LightingConfig{
		Channels: make(map[string]lighting.LEDChannel),
	},
}
