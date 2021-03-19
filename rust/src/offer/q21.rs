use std::ptr::null;

pub struct Solution{

}

pub struct Solution2{

}

impl Solution2 {
    pub fn exchange(nums: Vec<i32>) -> Vec<i32> {
        let mut result = Vec::new();
        for (i, num) in nums.iter().enumerate(){
            if num%2==1{
                result.insert(0, *num);
            }else{
                result.insert(i, *num);
            }
        }

        return result
    }
}

impl Solution {
    pub fn exchange(mut nums: Vec<i32>) -> Vec<i32> {
        let mut l:usize = 0;
        let mut r:usize = nums.len();
        while l < r{
            if nums[l]%2 == 0{
                r -= 1;
                nums.swap(l, r);
            }else{
                l += 1;
            }
        }

        nums
    }
}