# encoding: utf-8
class Node(object):

    def __init__(self, data=None, next_node=None):
        self.data = data
        self.next = next_node


def CompareLists(headA, headB):
    if headA is None and headB is None:
        return 1
    if headA is None or headB is None:
        return 0

    nodeA = headA
    nodeB = headB
    while nodeA is not None and nodeB is not None:
        if nodeA.data != nodeB.data:
            return 0
        nodeA = nodeA.next
        nodeB = nodeB.next
        if nodeA is None or nodeB is None:
            return 0
    return 1


def Reverse(head):
    if head is None:
        return None
    node = head
    values = []
    while node is not None:
        n = Node(node.data, None)
        values.append(n)
        node = node.next
    values.reverse()
    last = None
    for v in values:
        if last is not None:
            last.next = v
        last = v

    return values[0]


def ReversePrint(head):
    node = head
    values = []
    while node is not None:
        values.append(node.data)
        node = node.next
        if node is None:
            break
    values.reverse()
    for v in values:
        print(v)


def InsertNth(head, data, position):
    current = 0
    prev = None
    node = head
    if head is None:
        node = Node(data)
        return node
    while node is not None:
        if position == current:
            n = Node(data, node)
            if prev is not None:
                prev.next = n
            else:
                return n
            #node.next = n
            #return head
        prev = node
        node = node.next
        current += 1
    return head

# 0: a, next=b
# 1: b, next=c
# 2: c, next=None

# 0: a, next=x # next=xにする
# 1: x, next=b # next=bにする
# 2: b, next=c
# 3: c, next=None


def print_node(head):
    n = head
    while n is not None:
        print(n.data)
        n = n.next


if __name__ == '__main__':
    c = Node('c', None)
    b = Node('b', c)
    a = Node('a', b)
    print_node(a)
    print("---")
    # head = InsertNth(a, 'x', 0)
    # print_node(head)
    head = Reverse(a)
    print_node(head)

