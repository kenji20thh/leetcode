/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
const nums = [5, 0, 1, 4]


var firstStableIndex = function(nums, k) {
    let instabScore = Math.max(...nums) - Math.min(...nums)
    for (let i = 1; i < nums.length; i++) {
        let left = nums.slice(0, i)
        let right = nums.slice(i, nums.length)
        instabScore = Math.max(...left) - Math.min(...right)
        if (instabScore <= k) {
            return i
        }    
    }
    return 0  
};

console.log(firstStableIndex(nums, 2))