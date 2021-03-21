pub struct Solution{

}

impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        if matrix.len() == 0{
            return
        }
        let l = matrix.len();

        let mut rows = vec![false; matrix.len()];
        let mut cols = vec![false; matrix[0].len()];

        for (i, v) in matrix.iter().enumerate(){
            for (j, num) in v.iter().enumerate(){
                if *num == 0{
                    rows[i] = true;
                    cols[j] = true;
                }
            }
        }

        let mut i = 0;
        while i < matrix.len() {
            let mut j = 0;
            while j < matrix[i].len() {
                if rows[i] || cols[j]{
                    matrix[i][j] = 0;
                }
                j += 1;
            }
            i += 1;
        }
    }
}