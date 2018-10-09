package config

import "time"

type TomlConfig struct {
	Http 		HttpConfig
	Flashlog 	FlashlogCoreConfig
}

type HttpConfig struct {
	Server_ip_port string
}

type FilemonitorConfig struct {
	Active		bool
}

type FlashlogSubConfig struct {
	Timestep		int
	Rollsize		int64
}

type FlashlogTimertick struct {
	Timeout_logfile		time.Duration
	Empty_directory		time.Duration
}

type FlashlogCoreConfig struct {
	Logpath 	string
	Threshold	int64
	Standard	FlashlogSubConfig
	Nonstandard	FlashlogSubConfig
	Timertick	FlashlogTimertick
	Monitor		FilemonitorConfig
}