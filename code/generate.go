package code

// this part generates the code initially for each of the symbols and then for the input string itself
import (
	"fmt"
	"strings"

	huffman "github.com/mamad-nik/huffman"
)

func traverstree(node *huffman.Nodetype, input, code string) string {
	if node.Left != nil {
		if strings.Contains(node.Left.Symbole, input) {
			code += "0"
			code = traverstree(node.Left, input, code)
		} else if strings.Contains(node.Right.Symbole, input) {
			code += "1"
			code = traverstree(node.Right, input, code)
		}
	}
	return code
}
func sycogen(node *huffman.Nodetype, charachters []string) map[string]string {
	syco := make(map[string]string)
	for _, v := range charachters {
		syco[v] = traverstree(node, v, "")
	}
	fmt.Println("codes individually: ")
	for elem := range syco {
		fmt.Println(elem, ": ", syco[elem])
	}
	return syco
}
func Get(node *huffman.Nodetype, input string, characters []string) string {
	code := ""
	syco := sycogen(node, characters)

	inpchars := strings.Split(input, "")
	for _, v := range inpchars {
		code += syco[v]
	}
	return code
}
