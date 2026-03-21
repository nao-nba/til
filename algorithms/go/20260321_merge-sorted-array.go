package main

func merge(nums1 []int, m int, nums2 []int, n int) { // 全てint型と定義
	// 有効な要素数-1
	p1 := m - 1
	p2 := n - 1
	// outputの要素数-1
	p := m + n - 1

	// nums2の有効な要素が存在する間
	for p2 >= 0 {
		// nums1も存在、かつ指定された要素の値がnums1の方が大きい場合
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			// nums1の末尾にnums1の値を格納
			nums1[p] = nums1[p1]
			// p1の末尾は比較終了したので-1
			p1--
		} else {
			// nums2が大きかった場合
			nums1[p] = nums2[p2]
			p2--
		}
		// nums1の末尾-1
		p--
	}
}
