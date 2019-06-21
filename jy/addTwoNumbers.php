<?php
/**
 *   给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 *   如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 *   您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 *   示例：
 *
 *   输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 *   输出：7 -> 0 -> 8
 *   原因：342 + 465 = 807
 */

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution {

    /**
     * 两数相加
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $t = $l1;
        $k = 0;
        do{

            $val = $t->val + $l2->val + $k;

            $t->val = $val % 10;

            $k = $val >= 10 ? 1 : 0;

            if (!$l2->next && !$t->next && $k) {
                $t->next = new ListNode(1);
                break;
            }

            if ($t->next && !$l2->next) {
                $l2->next = new ListNode(0);
            }


            if ($l2->next && !$t->next) {
                $t->next = new ListNode(0);
            }

            $t = $t->next;

            $l2 = $l2->next;

        }while ($t);

        return $l1;

    }
}