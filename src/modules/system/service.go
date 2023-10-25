package system

import (
	"mailman/src/modules/system/dto"
	"runtime"
)

func GetMemoryUsage() *dto.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &dto.MemStats{
		Alloc:      m.Alloc,
		TotalAlloc: m.TotalAlloc,
		Free:       m.Frees,
		Sys:        m.Sys,
	}
}
