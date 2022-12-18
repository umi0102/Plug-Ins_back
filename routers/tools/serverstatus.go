package tools

import (
	"regexp"
	"strings"
)

func GetDataList(info string) []Datainfo {
	reg := regexp.MustCompile(".*?\n")
	di := Datainfo{}
	dis := make([]Datainfo, 0)
	replace := strings.Replace(info, " on", "", 1)

	for k, v := range reg.FindAllString(replace, -1) {
		if k == 0 {
			continue
		}
		i := make([]string, 0)
		for _, k := range strings.Split(v, " ") {

			if len(k) == 0 {
				continue
			}
			i = append(i, k)
		}

		di.Filesystem = i[0]
		di.Size = i[1]
		di.Used = i[2]
		di.Avail = i[3]
		di.Use = i[4]
		di.Mounted = strings.Replace(i[5], "\n", "", -1)
		dis = append(dis, di)
	}
	return dis
}

// GetCpuInfo 获取服务器CPU信息
func GetCpuInfo(info string) CpuInfo {
	modelnameIndex := strings.Index(info, "model name")
	modelnameIndexLast := strings.Index(info, "stepping")
	MHzIndex := strings.Index(info, "cpu MHz")
	MHzIndexLast := strings.Index(info, "cache size")
	CacheSizeIndex := strings.Index(info, "cache size")
	CacheSizeIndexLast := strings.Index(info, "physical id")
	CpuCoresIndex := strings.Index(info, "cpu cores")
	CpuCoresIndexLast := strings.Index(info, "apicid")

	return CpuInfo{
		CpuName:   info[modelnameIndex+13 : modelnameIndexLast-1],
		CpuMHz:    info[MHzIndex+11 : MHzIndexLast-1],
		CacheSize: info[CacheSizeIndex+13 : CacheSizeIndexLast-1],
		CpuCores:  info[CpuCoresIndex+12 : CpuCoresIndexLast-1]}
}

// GetMemInfo 获取服务器内存信息
func GetMemInfo(info string) MemInfo {
	MemTotalIndex := strings.Index(info, "MemTotal")
	MemTotalIndexLast := strings.Index(info, "MemFree")
	MemFreeIndex := strings.Index(info, "MemFree")
	MemFreeLast := strings.Index(info, "MemAvailable")
	MemAvailableIndex := strings.Index(info, "MemAvailable")
	MemAvailableIndexLast := strings.Index(info, "Buffers")
	BuffersIndex := strings.Index(info, "Buffers")
	BuffersIndexLast := strings.Index(info, "Cached")
	CachedIndex := strings.Index(info, "Cached")
	CachedIndexLast := strings.Index(info, "SwapCached")
	SwapCachedIndex := strings.Index(info, "SwapCached")
	SwapCachedIndexLast := strings.Index(info, "Active")
	ActiveIndex := strings.Index(info, "Active")
	ActiveIndexLast := strings.Index(info, "Inactive")
	InactiveIndex := strings.Index(info, "Inactive")
	InactiveIndexLast := strings.Index(info, "Active(anon)")
	SwapTotalIndex := strings.Index(info, "SwapTotal")
	SwapTotalIndexLast := strings.Index(info, "SwapFree")

	return MemInfo{
		MemTotal:     info[MemTotalIndex+17 : MemTotalIndexLast-1],
		MemFree:      info[MemFreeIndex+18 : MemFreeLast-1],
		MemAvailable: info[MemAvailableIndex+17 : MemAvailableIndexLast-1],
		Buffers:      info[BuffersIndex+18 : BuffersIndexLast-1],
		Cached:       info[CachedIndex+17 : CachedIndexLast-1],
		SwapCached:   info[SwapCachedIndex+20 : SwapCachedIndexLast-1],
		Active:       info[ActiveIndex+18 : ActiveIndexLast-1],
		Inactive:     info[InactiveIndex+18 : InactiveIndexLast-1],
		SwapTotal:    info[SwapTotalIndex+17 : SwapTotalIndexLast-1],
	}
}
