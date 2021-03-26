mod offer;
mod question;

use offer::q22;

#[derive(PartialEq, Eq, Clone, Debug)]
struct A {
    b: String
}

fn main()
{
    println!("{:?}", offer::q57::Solution::find_continuous_sequence(9));
}