use std::thread::current;
use std::cmp::max;

pub struct Solution {

}

impl Solution {
    pub fn max_sub_array(mut nums: Vec<i32>) -> i32 {
        let mut ans = nums[0];

        for i in 1..nums.len() {
            nums[i] = max(nums[i], nums[i-1] + nums[i]);
            ans = max(ans, nums[i]);
        }

        ans
    }
}

// impl Solution {
//     pub fn max_sub_array(nums: Vec<i32>) -> i32 {
//         let mut last = nums[0];
//         let mut ans = nums[0];
//
//         let mut i = 1;
//         while i < nums.len() {
//             last = max(nums[i], last + nums[i]);
//             ans = max(ans, last);
//             i += 1;
//         }
//
//         ans
//     }
// }

// impl Solution {
//     pub fn max_sub_array(nums: Vec<i32>) -> i32 {
//         let n = nums.len();
//         let mut i = 0;
//         let mut max_num = nums[0];
//
//         while i < n{
//             let mut j = i;
//             let mut cur = 0;
//             while j < n{
//                 cur += nums[j];
//                 max_num = max(max_num, cur);
//                 j += 1;
//             }
//
//             i += 1;
//         }
//
//         max_num
//     }
// }