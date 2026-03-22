func removeElement(nums []int, val int) int {
    k := 0
    // numsをループしてvをとりだす
    for _, v := range nums {
    // vがvalと違う値なら、
        if v != val {
            // nums[k]にvを格納
            nums[k] = v
            // k += 1
            k++
        }
    }
    // kを返す
    return k
}