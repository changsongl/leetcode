mod offer;
mod question;

use offer::q22;

#[derive(PartialEq, Eq, Clone, Debug)]
struct A {
    b: String
}

fn main()
{
    println!("{}", offer::q42::Solution::max_sub_array(vec![-2,1,-3,4,-1,2,1,-5,4]));
}