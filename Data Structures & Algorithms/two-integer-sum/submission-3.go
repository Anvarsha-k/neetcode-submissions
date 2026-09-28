func twoSum(nums []int, target int) []int {
    newMap:= make(map[int]int)
    list:=[]int{}
    
    for i:=0;i<len(nums);i++{
        value:=target - nums[i]
        if _,ok:=newMap[value];ok{
            list = append(list,newMap[value])
            list = append(list,i)
            return list
        }else{
            newMap[nums[i]] = i
        }
    }
    fmt.Println(newMap)
    return list
}
