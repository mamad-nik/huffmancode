package preprocess

// this part processes the input and build each node individually and puts them in the priority queue
import (
	"strings"

	huffman "github.com/mamad-nik/huffman"
)

func distinctchar(input string) []string {
	characters := strings.Split(input, "")
	distinct := characters[0]
	for i := 1; i < len(characters); i++ {
		if strings.ContainsAny(characters[i], distinct) {
			continue
		} else {
			distinct += characters[i]
		}
	}
	return strings.Split(distinct, "")
}
func frequency(input string, characters []string) map[string]int {
	m := make(map[string]int)
	for _, v := range characters {
		m[v] = strings.Count(input, v)
	}
	return m
}

func nodemaker(syfre map[string]int) []huffman.Nodetype {
	nodes := make([]huffman.Nodetype, 0)
	for elem := range syfre {
		nodes = append(nodes, huffman.Nodetype{
			Symbole:   elem,
			Frequency: syfre[elem],
			Left:      nil,
			Right:     nil})
	}
	return nodes
}

func Process(input string) ([]huffman.Nodetype, []string) {
	characters := distinctchar(input)
	syfre := frequency(input, characters)

	return nodemaker(syfre), characters
}
