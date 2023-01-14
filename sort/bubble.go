package sort

var BubbleSort = func(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		// 後方から検索し現在のIndex(i)までを範囲とする、つまり探索済の前方を走査しないことで、処理効率化を図っている
		for j := length - 1; i < j; j-- {
			if array[j-1] > array[j] {
				tmp := array[j]
				array[j] = array[j-1]
				array[j-1] = tmp
			}
		}
	}
	return array
}
