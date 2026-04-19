func topKFrequent(nums []int, k int) []int {
	var count = make(map[int]int)
	newarr:=[]int{}
	var bucket = make([][]int,len(nums)+1)
	for _,ch:=range nums{
		count[ch] = count[ch] + 1
	}
	
	for key,value:=range count{
		bucket[value] = append(bucket[value],key)
	}
	for i:=len(bucket)-1;i>=0 && len(newarr)<k;i--{
	   for _,ve:=range bucket[i]{
	       newarr = append(newarr,ve)
	       if (len(newarr) == k){
	           break
	       }
	   }
	}
	fmt.Println(bucket)
    return newarr
}
