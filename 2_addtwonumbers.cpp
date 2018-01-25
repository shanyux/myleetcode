/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int result = 0;
		ListNode* head = new ListNode(0);
        ListNode* pnode = head;
		while(true){
			if(l1){
				result += l1->val;
				l1 = l1->next;
			}
			if(l2){
				result += l2->val;
				l2 = l2->next;
			}
			pnode->val = result % 10;
			result = result / 10;
			if ( l1 || l2 || result){
				pnode->next = new ListNode(0);
                pnode = pnode->next;
            }
			else
				break;

		}
		return head;        
    }
};


class Solution {
public:
	ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
		int result = 0;
		ListNode headnode(0), *p = &headnode;
		while (l1 || l2 || result){
			result = (l1 ? l1->val : 0) + (l2 ? l2->val : 0) + result;
			p->next = new ListNode(result % 10);
			result = result / 10;
			p = p->next;
			l1 = l1 ? l1->next : l1;
			l2 = l2 ? l2->next : l2;

		}
		return headnode.next;
	}
};