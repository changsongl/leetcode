// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

struct Solution {
}

// impl Solution {
//     pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
//         let mut iter = head.as_mut();
//         let has_head = false;
//
//         while let Some(node) = iter{
//             let mut dup = false;
//             while let Some(nn) = node.next.as_ref(){
//                 if node.val == nn.val {
//                     dup = true;
//                     node.next = nn.next.clone();
//                 }else{
//                     if !has_head{
//                         head =
//                     }
//                     break;
//                 }
//             }
//             iter = node.next.as_mut();
//         }
//
//         new_head
//     }
// }