func twoSum(nums []int, target int) []int {
    newMap:= make(map[int]int)
    
    for i,n :=range nums{
        diff:=target - n
        if v,ok:=newMap[diff];ok{
            return []int{v,i}
        }
        newMap[n] = i
    }
    return []int{}
}
