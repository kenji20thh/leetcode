/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var missingMultiple = function(nums, k) {
    let found = false
    let i = k
    while (!found) {
        if (i % k == 0 && !nums.includes(i)) return i 
        i += k
    }
    
};