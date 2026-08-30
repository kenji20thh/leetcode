/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumDeletions = function(nums) {
    const min = Math.min(...nums);
    const max = Math.max(...nums);
    const minIndex = nums.indexOf(min);
    const maxIndex = nums.indexOf(max);
   if (minIndex > maxIndex) {
       return nums.length - minIndex + maxIndex + 1;  
   }
   return nums.length - maxIndex + minIndex + 1;

};