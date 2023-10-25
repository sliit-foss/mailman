package dto

type MemStats struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Free       uint64 `json:"free"`
	Sys        uint64 `json:"sys"`
}
