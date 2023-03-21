package startup

import "github.com/FarmerChillax/AirRocket/vars"

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	vars.AppSettings = &vars.AppSetting{
		Address:       "127.0.0.1:5000",
		DurationTime:  3600,
		GenerateCount: 10,
	}

	return nil
}

func LoadNodeServerConfig() error {
	settings := vars.ClientServerSetting{
		Address:  "127.0.0.1:6000",
		RootPath: "./",
	}
	vars.ServerSettings = &settings
	return nil
}
