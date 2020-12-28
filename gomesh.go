package gomesh

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Msh struct {
	Version   string
	IsAscii   bool
	Precision int

	Nodes    []Node
	Elements []Element
}

type Node struct {
}
type Element struct {
}

func Parse(filename string) Msh {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	msh := Msh{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "$MeshFormat":
			msh.Version, msh.IsAscii, msh.Precision, err = ParseFormat(scanner)
		case "$Nodes":
			msh.Nodes, err = ParseNodes(scanner)
		case "$Elements":
			err = ParseElements(scanner)
		default:
			log.Printf("WARN: Tag %v not implemented, skipping", line)
		}
		if err != nil {
			log.Fatal("Error in file parsing: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return Msh{} //TODO:
}
func ParseFormat(sc *bufio.Scanner) (string, bool, int, error) {
	line := sc.Text()
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
	prec, err := strconv.Atoi(x[2])
	if err != nil {
		return "", false, 0, errors.New("invalid precision value in the format: " + line)
	}
	sc.Text() // read last $EndMeshFormat line
	return x[0], isAscii, prec, nil
}

func ParseNodes(sc *bufio.Scanner) ([]Node, error) {
	line := sc.Text()
	meta := strings.Split(line, " ")
	if len(meta) != 4 {
		return nil, errors.New("nodes meta len!=4: " + line)
	}
	numEnts, err := strconv.Atoi(meta[0])
	if err != nil {
		return nil, errors.New("invalid numEnts: " + line)
	}
	numNods, err := strconv.Atoi(meta[1])
	if err != nil {
		return nil, errors.New("invalid numNods: " + line)
	}
	/*	minNodeTag, err := strconv.Atoi(meta[2])
		if err!= nil {
			return nil, errors.New("invalid minNodeTag: "+ line)
		}
		maxNodeTag, err := strconv.Atoi(meta[3])
		if err!= nil {
			return nil, errors.New("invalid maxNodeTag: "+ line)
		}*/
	for i := 0; i < numEnts; i++ {
		line := sc.Text()
		meta := strings.Split(line, " ")
		if len(meta) != 4 {
			return nil, errors.New("ents meta len!=4: " + line)
		}
		for j := 0; j < numNods; j++ {
			//read node tags
		}
		for j := 0; j < numNods; j++ {
			//read coordinates for tags
		}
		//mesh.add entity()
		return nil, errors.New("Not implemented")
	}

	sc.Text()       // read last $EndNodes line
	return nil, nil //TODO:
}

func ParseElements(sc *bufio.Scanner) error {
	line := sc.Text()
	_ = strings.Split(line, " ")

	sc.Text()                            // read last $EndNodes line
	return errors.New("Not implemented") //TODO:
}
