package tree

type CharNode struct {
	Char  rune
	Count int
	Left  *CharNode
	Right *CharNode
}
