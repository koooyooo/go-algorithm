package sort

var InsertSort = func(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		insIdx := i
		insVal := array[i]
		// 挿入位置を確認
		for j := 0; j < i; j++ {
			if insVal < array[j] {
				insIdx = j
				break
			}
		}
		// 挿入位置 - 処理中の位置
		for k := i; k >= insIdx; k-- {
			if k == 0 {
				continue
			}
			array[k] = array[k-1]
		}
		array[insIdx] = insVal
	}
	return array
}
