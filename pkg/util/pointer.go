package util

func PtrInt32(n int) *int32 {
	tmp := int32(n)
	return &tmp
}
