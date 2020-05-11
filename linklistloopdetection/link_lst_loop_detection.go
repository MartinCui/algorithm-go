package linklistloopdetection

/*
Single direction unordered link list which might or might not have loop.
Judge whether this list has a loop and return the first node in loop if found.
Otherwise return nil.

a -> b -> c -> d -> e
return nil

a -> b -> c -> d -> e -> c
return c
*/

type node struct {
	key  int
	next *node
}

func findLoopEntrance(root *node) *node {
	slow := root
	fast := root
	for {
		if slow == nil || slow.next == nil || fast == nil || fast.next == nil || fast.next.next == nil {
			return nil
		}
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			break
		}
	}

	slow = root
	for {
		if slow == fast {
			return slow
		}

		slow = slow.next
		fast = fast.next
	}
}
