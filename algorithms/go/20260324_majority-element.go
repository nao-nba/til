func majorityElement(nums []int) int {
    // 出力candidateにnums[0]をセット
    // count = 0
    // numsをループして値を取り出す
    // nums[0]が暫定の出力kとする
    // 値がnums[0]と同一ならcount+1,不一致ならcount-1
    // count<= 0になったら、その時の値を暫定kに更新
    // 最後にcandidateを返す

    candidate := nums[0]
    count := 0

    for _, n := range nums {
        if count == 0 {
            candidate = n
        }

        if n == candidate {
            count++
        } else {
            count--
        }

    }
    return candidate
}