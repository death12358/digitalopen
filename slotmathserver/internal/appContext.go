package internal

import (
	"digitalopen/games/scriptor"
	"digitalopen/slotmathserver/resource"
	"time"

	fasthttp "github.com/Bofry/host-fasthttp"
)

type (
	AppContext struct {
		Host            *Host
		Config          *Config
		ServiceProvider *ServiceProvider
	}

	Host fasthttp.Host

	Config struct {
		// host-fasthttp server configuration
		ListenAddress  string `yaml:"address"        arg:"address;the combination of IP address and listen port"`
		EnableCompress bool   `yaml:"useCompress"    arg:"use-compress;indicates the response enable compress or not"`
		ServerName     string `yaml:"serverName"`
		Version        string `resource:".VERSION"`

		// put your configuration below
		MathRedis_Host     string `yaml:"Math_Redis_Host"        env:"Math_Redis_Host"`
		MathRedis_Port     int    `yaml:"Math_Redis_Port"        env:"Math_Redis_Port"`
		MathRedis_Password string `yaml:"Math_Redis_Password"    env:"Math_Redis_Password"`
	}

	ServiceProvider struct{}
)

func (h *Host) Init(conf *Config) {
	h.Server = &fasthttp.Server{
		Name: conf.ServerName,
		// 考慮啟用 Keep-alive，這有助於減少連接的建立和拆除，特別是在頻繁請求的情況下。
		DisableKeepalive: false,

		// 保持 Header 名稱正常化，這樣可以確保 Header 名稱符合 HTTP 標準。
		DisableHeaderNamesNormalizing: false,

		// 設置最大連接數。這有助於控制同時處理的請求數量。
		MaxConnsPerIP: 100,

		// 設置讀取和寫入緩衝區大小。根據您的應用需求和可用內存量，可以增大或減小這些值。
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,

		// 設置請求和響應緩衝區大小。增大這些值可以提高性能，但會增加內存使用量。
		MaxRequestBodySize: 4 * 1024 * 1024, // 4 MB

		// 設置連接閒置超時。這有助於釋放閒置連接的資源，讓伺服器能處理更多的活動連接。
		IdleTimeout: 60 * time.Second,

		// 設置每個連接的讀取超時。
		ReadTimeout: 5 * time.Second,

		// 設置每個連接的寫入超時。
		WriteTimeout: 360 * time.Second,
	}
	h.ListenAddress = conf.ListenAddress
	h.EnableCompress = conf.EnableCompress
	h.Version = conf.Version
}

func (p *ServiceProvider) Init(conf *Config) {
	// initialize service provider components
	resource.NewSlotScriptor(&scriptor.Option{
		Host:     conf.MathRedis_Host,
		Port:     conf.MathRedis_Port,
		Password: conf.MathRedis_Password,
	})
}
