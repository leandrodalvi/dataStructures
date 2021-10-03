package list

type List struct {
	first *Node
	last  *Node
}

type Node struct {
	key   string
	value int
	next  *Node
}

func create() *List {
	return &List{
		first: nil,
		last:  nil,
	}
}

func insert(list *List, key string, value int) {
	node := &Node{key: key, value: value, next: nil}
	if list.first == nil {
		list.first = node
		list.last = node
	} else {
		aux := list.last
		aux.next = node
		list.last = node
	}
}

func remove(list *List, key string) bool {
	//checking trivial cases
	if list == nil || list.first == nil {
		return false
	}
	if list.first.key == key {
		list.first = list.first.next
	}

	//searching for node to remove
	var nodeBeforeAux *Node = nil
	aux := list.first
	for aux != nil && aux.key != key {
		nodeBeforeAux = aux
		aux = aux.next
	}

	//checking if we found it, and if we did, removing it
	if aux.key != key {
		return false
	} else {
		nodeBeforeAux.next = aux.next
		aux = nil
		return true
	}
}

func search(list *List, key string) *Node {
	if list == nil {
		return nil
	}
	aux := list.first
	for aux != nil && aux.key != key {
		aux = aux.next
	}
	return aux
}
