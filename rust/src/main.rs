mod offer;
mod question;

use offer::q22;

#[derive(PartialEq, Eq, Clone, Debug)]
struct A {
    b: String
}

fn main()
{
    let l1 = Box::new(offer::q25::ListNode::new(1));
    let l2 = Box::new(offer::q25::ListNode::new(2));
    println!("{:?}", offer::q25::Solution::merge_two_lists(Some(l1), Some(l2)));
}