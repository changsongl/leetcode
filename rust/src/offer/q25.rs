use std::cmp::min;

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
pub struct Solution {

}

impl Solution {
    pub fn merge_two_lists(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>)
        -> Option<Box<ListNode>> {
        let (mut l1, mut l2) = (l1, l2);
        let mut dummy = Box::new(ListNode::new(0));
        let mut tail = &mut dummy;
        while let (Some(head1), Some(head2)) = (&l1, &l2) {
            if head1.val < head2.val {
                tail.next = l1;
                tail = tail.next.as_mut().unwrap();
                l1 = tail.next.take();
            }else{
                tail.next = l2;
                tail = tail.next.as_mut().unwrap();
                l2 = tail.next.take();
            }
        }

        if l1.is_some() {
            tail.next = l1
        }else if l2.is_some(){
            tail.next = l2
        }

        return dummy.next
    }
}