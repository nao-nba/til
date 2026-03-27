package algorithms
func romanToInt(s string) int {
    m := map[byte]int{
        'I': 1, 'V': 5, 'X': 10, 'L': 50,
        'C': 100, 'D': 500, 'M': 1000,
    }
    ans := 0
    n := len(s)
    for i := 0; i < n; i++ {
        val := m[s[i]]
        // 次の文字が存在し、かつ現在の値より大きい場合は引き算
        if i < n-1 && val < m[s[i+1]] {
            ans -= val
        } else {
            ans += val
        }
    }
    return ans
}

// mapはkey:valueの仮想配列
// key: 英数字1文字の場合はbyte, 日本語1文字の場合はrune, 2文字以上の場合はstring