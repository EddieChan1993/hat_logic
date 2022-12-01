package util

//SliceIsMember slice中是否包含该元素
func SliceIsMember(s []int32, e int32) bool {
	for _, i := range s {
		if e == i {
			return true
		}
	}
	return false
}

//SliceAppend slice追加
func SliceAppend(base, add []int32) []int32 {
	res := make([]int32, 0, len(base)+len(add))
	res = append(res, base...)
	res = append(res, add...)
	return res
}
