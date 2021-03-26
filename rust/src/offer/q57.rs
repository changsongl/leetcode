pub struct Solution {

}

impl Solution {
    pub fn find_continuous_sequence(target: i32) -> Vec<Vec<i32>> {
        let mut sum = 0;
        let (mut ans, mut v) = (vec![], vec![]);
        let first:usize = 0;
        for i in 1..target/2+2 {
            v.push(i);
            sum += i;

            while sum > target {
                let num = v.remove(first);
                sum -= num;
            }

            if sum == target{
                ans.push(v.clone())
            }
        }

        ans
    }
}