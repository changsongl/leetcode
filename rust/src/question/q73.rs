// pub struct SolutionMN{
//
// }
//
// impl SolutionMN {
//     pub fn set_zeroes_mn(matrix: &mut Vec<Vec<i32>>) {
//         if matrix.len() == 0{
//             return
//         }
//         let l = matrix.len();
//
//         let mut rows = vec![false; matrix.len()];
//         let mut cols = vec![false; matrix[0].len()];
//
//         for (i, v) in matrix.iter().enumerate(){
//             for (j, num) in v.iter().enumerate(){
//                 if *num == 0{
//                     rows[i] = true;
//                     cols[j] = true;
//                 }
//             }
//         }
//
//         let mut i = 0;
//         while i < matrix.len() {
//             let mut j = 0;
//             while j < matrix[i].len() {
//                 if rows[i] || cols[j]{
//                     matrix[i][j] = 0;
//                 }
//                 j += 1;
//             }
//             i += 1;
//         }
//     }
// }

// impl Solution {
//     pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
//         let rows = matrix.len();
//         let cols = matrix[0].len();
//         let mut row_zeros = Vec::new();
//         let mut col_zeros = Vec::new();
//         for row in 0..rows {
//             for col in 0..cols {
//                 if matrix[row][col] == 0 {
//                     row_zeros.push(row);
//                     col_zeros.push(col);
//                 }
//             }
//         }
//         for row in row_zeros.into_iter() {
//             for col in 0..cols {
//                 matrix[row][col] = 0;
//             }
//         }
//         for col in col_zeros.into_iter() {
//             for row in 0..rows {
//                 matrix[row][col] = 0;
//             }
//         }
//     }
// }

pub struct Solution{
}

impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let mut m = matrix.len();
        let mut n = matrix[0].len();
        let mut col_zero = false;

        let mut row = 0;
        while row < m {
            if matrix[row][0] == 0{
                col_zero = true;
            }

            let mut i = 1;
            while i < n{
                if matrix[row][i] == 0{
                    matrix[row][0] = 0;
                    matrix[0][i] = 0;
                }
                i += 1;
            }
            row += 1;
        }

        let mut i:i32 = (matrix.len() - 1) as i32;
        while i >= 0 {
            let mut j = n - 1;
            while j > 0 {
                if matrix[i as usize][0] == 0 || matrix[0][j] == 0{
                    matrix[i as usize][j] = 0;
                }
                j -= 1;
            }

            if col_zero{
                matrix[i as usize][0] = 0;
            }

            i -= 1;
        }
    }
}