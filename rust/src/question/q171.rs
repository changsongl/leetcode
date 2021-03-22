pub struct Solution{

}

// impl Solution {
//     pub fn hammingWeight (mut n:u32) -> i32 {
//         let mut v:Vec<u32> = vec![0; 32];
//         if v[0] == 0{
//             v[0] = 1;
//             for i in 1..v.len(){
//                 v[i] = v[i-1] * 2;
//             }
//         }
//
//         let mut i = 0;
//         let mut num = 0;
//         while i < v.len() {
//             if v[i] > n{
//                 break;
//             }
//             i += 1;
//         }
//
//         while i > 0{
//             if n >= v[i-1]{
//                 n -= v[i-1];
//                 num += 1;
//             }
//
//             i -= 1;
//         }
//
//         num
//     }
// }

impl Solution {
    pub fn hammingWeight (n: u32) -> i32 {
        let mut count = 0;
        let mut n = n;
        while n != 0 {
            n = n & (n - 1);
            count += 1;
        }
        count
    }
}