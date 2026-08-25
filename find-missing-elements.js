/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findMissingElements = function(nums) {
  let result = []
  nums.sort((a,b) => a-b)
  const start = nums[0]
  const end = nums[nums.length - 1]
  for (let i = start ; i < end ; i ++) {
    if (!nums.includes(i)) result.push(i)
    continue 
  }
  return result
};
