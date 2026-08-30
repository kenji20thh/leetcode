/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumDeletions = function(nums) {
    const min = Math.min(...nums);
    const max = Math.max(...nums);
    const minIndex = nums.indexOf(min);
    const maxIndex = nums.indexOf(max);
    console.log(minIndex, maxIndex);
   if (minIndex < nums.length / 2&& maxIndex < nums.length / 2) {
        return Math.max(minIndex, maxIndex) + 1;
    } else if (minIndex >= nums.length / 2 && maxIndex >= nums.length / 2) {
        return Math.max(nums.length - minIndex, nums.length - maxIndex);
    } else {
        return Math.min(minIndex + 1 + nums.length - maxIndex, maxIndex + 1 + nums.length - minIndex);
    }

};

const nums = [2,10,7,5,4,1,8,6]

console.log(minimumDeletions(nums));