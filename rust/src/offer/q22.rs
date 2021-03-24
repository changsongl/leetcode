use std::ops::Deref;

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

pub struct Solution{

}

impl Solution {
    pub fn get_kth_from_end(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        let (mut fast, mut slow) = (&head, &head);
        let mut i = 0;
        loop {
            match fast {
                Option::None => {
                    break;
                },
                Option::Some(b) => {
                    fast = &b.next;
                }
            }

            if i >= k {
                match slow {
                    Option::None => {
                        break;
                    },
                    Option::Some(b) => {
                        slow = &b.next;
                    }
                }
            }

            i += 1;
        }

        slow.clone()
    }
}

// impl Solution {
//     pub fn get_kth_from_end(mut head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
//         let mut length = 0;
//
//         let mut node = &mut head;
//         loop {
//             if let Some(x) = node {
//                 length += 1;
//                 node = &mut x.next;
//             } else {
//                 break;
//             }
//         }
//
//         let length = length;
//         let mut node = head;
//         for i in 0..length-k{
//             if let Some(x) = node {
//                 node = x.next;
//             }
//         }
//         node
//     }
// }

// impl Solution {
//     pub fn get_kth_from_end(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
//         let (mut quick, mut slow) = (&head, &head);
//         for _ in 1..k {
//             quick = &quick.as_ref().unwrap().next
//         }
//
//         while quick.as_ref().unwrap().next.is_some(){
//             quick = &quick.as_ref().unwrap().next;
//             slow = &slow.as_ref().unwrap().next;
//         }
//
//         slow.clone()
//     }
// }