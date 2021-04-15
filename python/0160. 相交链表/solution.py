#!/usr/bin/python

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        h1, h2 = headA, headB
        while h1 != h2:
            if h1 is None:
                h1 = headB
                continue
            elif h2 is None:
                h2 = headA
                continue

            h1 = h1.next
            h2 = h2.next
        return h1
