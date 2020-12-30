package gomesh

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type nextLineFunc func() string

func nextLineReader(scanner *bufio.Scanner) nextLineFunc {
	return func() string {
		scanner.Scan()
		return scanner.Text()
	}
}
type Msh struct {
	Version  string
	IsAscii  bool
	DataSize int

	Nodes    []Node
	Elements []Element
}

type Node struct {
	Tag, X, Y, Z string
}

func NewNode(tag, x, y, z string) Node {
	return Node{
		Tag: tag,
		X:   x,
		Y:   y,
		Z:   z,
	}
}

type Element struct {
	Tag, Type string
	Tags      []string
	NodeTags  []string
	Data      []string
}

func NewElement(tag, typ string, tags, nodeTags []string) Element {
	return Element{
		Tag:      tag,
		Type:     typ,
		Tags:     tags,
		NodeTags: nodeTags,
	}
}

func (m *Msh) SetElementData(data map[string][]string) {
	ind := map[string]int{}
	for i, el := range m.Elements {
		ind[el.Tag] = i
	}
	for k, v := range data {
		m.Elements[ind[k]].Data = v
	}
}

func ReadVersion(filename string) (string, bool, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "$MeshFormat":
			return ParseFormat(	nextLineReader(scanner)	)
		}
	}
	return "", false, 0, errors.New("Mesh Format not found ")
}

func Parse(filename string) (*Msh, error) {
	version, _, _, err := ReadVersion(filename)
	if err != nil {
		log.Fatalf("Failed read version %v", err)
	}
	//0 Some mesh files out there have the version specified as version "2" when it really is
	// "2.2". Same with "4" vs "4.1".
	switch version {
	case "2":
		return Parse22(filename)
	case "2.2":
		return Parse22(filename)
	case "4.0":
		// implement me?
		break
	case "4":
		return Parse41(filename)
	case "4.1":
		return Parse41(filename)
	}

	return nil, fmt.Errorf("version is not recognized %v", version)
}

func ParseFormat(nextLine nextLineFunc) (string, bool, int, error) {
	line := nextLine()
	x := strings.Split(line, " ")
	if len(x) != 3 {
		return "", false, 0, errors.New("format len!=3: " + line)
	}
	isAscii := false
	switch x[1] {
	case "0":
		isAscii = true
	case "1":
		isAscii = false
	default:
		return "", false, 0, errors.New("invalid isAscii value in the format: " + line)
	}
	datasize, err := strconv.Atoi(x[2])
	if err != nil {
		return "", false, 0, errors.New("invalid datasize value in the format: " + line)
	}
	nextLine() // read last $EndMeshFormat line
	return x[0], isAscii, datasize, nil
}
