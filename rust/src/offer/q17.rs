pub fn print_numbers(n: i32) -> Vec<i32> {
    let mut max = 0;
    let mut i = 1;
    while i <= n {
        max = max * 10 + 9;
        i += 1
    }

    let mut ans = Vec::new();

    for i in 1..max + 1 {
        ans.push(i)
    }

    return ans
}