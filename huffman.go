package huffmancode

// data structures declarations
type Nodetype struct {
	Symbole   string
	Frequency int

	Left  *Nodetype
	Right *Nodetype
}
type Pqueue struct {
	Nodes []Nodetype
	Cap   int
}
