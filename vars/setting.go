package vars

type AppSetting struct {
	// 服务监听地址
	Address string
	// 缓存默认过期时间(秒)
	DurationTime int64
	// 提取码尝试生成次数
	GenerateCount int
}

var AppSettings = &AppSetting{}

type ClientServerSetting struct {
	Address  string
	RootPath string
}

var ServerSettings = &ClientServerSetting{}
