mod offer;
mod question;

use offer::q22;

fn main()
{
    let mut l1 = q22::ListNode::new(10);
    let mut l2 = q22::ListNode::new(20);
    let l3 = q22::ListNode::new(30);
    l2.next = Some(Box::new(l3));
    l1.next = Some(Box::new(l2));
    let list = Some(Box::new(l1));
    // // println!("{}", p.val);
    // // if let Some(q) = &p.next {
    // //     println!("{}", q.val);
    // }
    // dbg!(reverse_print(list));
    // delete_node(list, 4);
    // reverse_list(list);
    q22::Solution::get_kth_from_end(list, 3);
}