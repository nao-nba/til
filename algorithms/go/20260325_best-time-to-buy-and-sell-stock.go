package go
func maxProfit(prices []int) int {
    // prices length=1の場合は0を返す
    // 暫定最安値を定義（初期値はprices[0]）
    // 最大利益を定義（初期値0）
    // ループを回してindexを取り出す
    // 暫定最安値より小さい値があれば、値を更新
    // それ以外は暫定最大利益として保持
    // もし暫定最大利益の方が最大利益よりも大きければ更新
    // 最大利益を返す
    
    if len(prices) == 0 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0

    for i := 1; i < len(prices); i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else {
            profit := prices[i] - minPrice
            if profit > maxProfit {
                maxProfit = profit
            }
        }
    }
    return maxProfit

}