package sort

var BubbleSort = func(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
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
