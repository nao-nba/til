func removeDuplicates(nums []int) int {
    // kをnums[i]との判定用に用意。初期値1で。
    // numsをループ
    // nums[i]がnums[k-1]と違う値なら、nums[k]=nums[i]
    // k+1
    // kを返す
		
    if len(nums) == 0 {
        return 0
    }

    k := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[k-1] {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}