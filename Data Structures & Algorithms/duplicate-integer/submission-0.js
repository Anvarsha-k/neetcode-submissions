class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    nums=[1,2,3,3]
    hasDuplicate(nums) {
    for(let i=0;i<nums.length;i++){
        for(let j=i+1;j<nums.length;j++)
        {
            if(nums[i]==nums[j])
            {
                console.log(nums[i])
                return true
            }
            
        }
       
       
    
    }
    return false
}
}