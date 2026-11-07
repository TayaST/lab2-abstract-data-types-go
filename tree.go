package main

import "fmt"

// ===================== BST =====================

type TNode struct {
	value       int
	left, right *TNode
}

type BST struct {
	root  *TNode
	count int
}

// ---------- Вставка ----------

func (t *BST) insertRec(node *TNode, v int) (*TNode, bool) {
	if node == nil {
		return &TNode{value: v}, true
	}
	if v < node.value {
		n, added := t.insertRec(node.left, v)
		node.left = n
		return node, added
	} else if v > node.value {
		n, added := t.insertRec(node.right, v)
		node.right = n
		return node, added
	}
	// равные игнорируем
	return node, false
}

func (t *BST) Insert(v int) {
	var added bool
	t.root, added = t.insertRec(t.root, v)
	if added {
		t.count++
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR: duplicate value")
	}
}

func (t *BST) InsertSilent(v int) {
	var added bool
	t.root, added = t.insertRec(t.root, v)
	if added {
		t.count++
	}
}

// ---------- Поиск ----------

func (t *BST) Get(v int) {
	cur := t.root
	depth := 0
	for cur != nil {
		if v == cur.value {
			fmt.Printf("FOUND (depth %d)\n", depth)
			return
		}
		if v < cur.value {
			cur = cur.left
		} else {
			cur = cur.right
		}
		depth++
	}
	fmt.Println("NOT FOUND")
}

func (t *BST) IsMember(v int) {
	cur := t.root
	for cur != nil {
		if v == cur.value {
			fmt.Println("TRUE")
			return
		}
		if v < cur.value {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	fmt.Println("FALSE")
}

// ---------- Удаление ----------

func findMin(node *TNode) *TNode {
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

func (t *BST) removeRec(node *TNode, v int) (*TNode, bool) {
	if node == nil {
		return nil, false
	}
	var removed bool
	if v < node.value {
		node.left, removed = t.removeRec(node.left, v)
	} else if v > node.value {
		node.right, removed = t.removeRec(node.right, v)
	} else {
		removed = true
		// 0 или 1 потомок
		if node.left == nil {
			return node.right, true
		}
		if node.right == nil {
			return node.left, true
		}
		// 2 потомка: заменяем минимумом из правого поддерева
		mn := findMin(node.right)
		node.value = mn.value
		node.right, _ = t.removeRec(node.right, mn.value)
	}
	return node, removed
}

func (t *BST) Remove(v int) {
	var removed bool
	t.root, removed = t.removeRec(t.root, v)
	if removed {
		t.count--
		fmt.Println("OK")
	} else {
		fmt.Println("NOT FOUND")
	}
}

// ---------- IsComplete (обход BFS) ----------

func (t *BST) IsComplete() {
	if t.root == nil {
		fmt.Println("TRUE (empty tree)")
		return
	}
	queue := []*TNode{t.root}
	seenGap := false
	ok := true
	for i := 0; i < len(queue); i++ {
		n := queue[i]
		if n.left != nil {
			if seenGap {
				ok = false
				break
			}
			queue = append(queue, n.left)
		} else {
			seenGap = true
		}
		if n.right != nil {
			if seenGap {
				ok = false
				break
			}
			queue = append(queue, n.right)
		} else {
			seenGap = true
		}
	}
	if ok {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

// ---------- Обходы ----------

func inorderRec(n *TNode) {
	if n == nil {
		return
	}
	inorderRec(n.left)
	fmt.Printf("%d ", n.value)
	inorderRec(n.right)
}

func preorderRec(n *TNode) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.value)
	preorderRec(n.left)
	preorderRec(n.right)
}

func postorderRec(n *TNode) {
	if n == nil {
		return
	}
	postorderRec(n.left)
	postorderRec(n.right)
	fmt.Printf("%d ", n.value)
}

func (t *BST) Inorder() {
	inorderRec(t.root)
	fmt.Println()
}

func (t *BST) Preorder() {
	preorderRec(t.root)
	fmt.Println()
}

func (t *BST) Postorder() {
	postorderRec(t.root)
	fmt.Println()
}

func (t *BST) LevelOrder() {
	if t.root == nil {
		fmt.Println()
		return
	}
	queue := []*TNode{t.root}
	for i := 0; i < len(queue); i++ {
		n := queue[i]
		fmt.Printf("%d ", n.value)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
	fmt.Println()
}

// PreorderSlice возвращает значения в pre-order (для сохранения в файл).
func (t *BST) PreorderSlice() []int {
	out := []int{}
	var rec func(n *TNode)
	rec = func(n *TNode) {
		if n == nil {
			return
		}
		out = append(out, n.value)
		rec(n.left)
		rec(n.right)
	}
	rec(t.root)
	return out
}

func (t *BST) Len() {
	fmt.Println(t.count)
}

func (t *BST) Print() {
	if t.root == nil {
		fmt.Println("(empty tree)")
		return
	}
	fmt.Print("Inorder (sorted): ")
	inorderRec(t.root)
	fmt.Println()
	fmt.Print("Preorder: ")
	preorderRec(t.root)
	fmt.Println()
	fmt.Print("Level-order (BFS): ")
	t.LevelOrder()
	fmt.Printf("Count: %d\n", t.count)
}