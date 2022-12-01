package util

//MapMergeSum map合并，值累加
func MapMergeSum(base, add map[int32]int64) map[int32]int64 {
	if len(base) == 0 {
		return add
	}
	for k, v := range add {
		if old, had := base[k]; had {
			base[k] = old + v
		} else {
			base[k] = v
		}
	}
	return base
}
