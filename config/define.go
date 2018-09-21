package config

type TomlConfig struct {
	Http 	HttpConfig
	Seelog	SeelogConfig
}

type HttpConfig struct {
	Server_ip_port string
}

type SeelogConfig struct {
	Config_path string
}