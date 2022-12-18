package tools

// ServerInfo 单条服务器数据
type ServerInfo struct {
	CpuInfo  CpuInfo
	MemInfo  MemInfo
	DataList []Datainfo
	Test     string
}

// CpuInfo CPU
type CpuInfo struct {
	CpuName   string `json:"cpu_name"`
	CpuMHz    string `json:"cpu_MHz"`
	CacheSize string `json:"cache size"`
	CpuCores  string `json:"cpu cores"`
}

// MemInfo 内存
type MemInfo struct {
	MemTotal     string `json:"MemTotal"`
	MemFree      string `json:"MemFree"`
	MemAvailable string `json:"MemAvailable"`
	Buffers      string `json:"Buffers"`
	Cached       string `json:"Cached"`
	SwapCached   string `json:"SwapCached"`
	Active       string `json:"Active"`
	Inactive     string `json:"Inactive"`
	SwapTotal    string `json:"SwapTotal"`
}

// Datainfo 存储
type Datainfo struct {
	Filesystem string `json:"Filesystem"`
	Size       string `json:"Size"`
	Used       string `json:"Used"`
	Avail      string `json:"Avail"`
	Use        string `json:"Use%"`
	Mounted    string `json:"Mounted on"`
}
