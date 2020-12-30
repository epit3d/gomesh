package gomesh

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Parse41(filename string) (* Msh, error) {
	msh := Msh{}

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
			msh.Version, msh.IsAscii, msh.DataSize, err = ParseFormat(nextLineReader(scanner))
		case "$Nodes":
			msh.Nodes, err = ParseNodes41(scanner)
		case "$Elements":
			err = ParseElements41(scanner)
		default:
			log.Printf("WARN: Tag %v not implemented, skipping", line)
		}
		if err != nil {
			log.Fatalf("Error in file parsing: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &msh, nil
}


func ParseNodes41(sc *bufio.Scanner) ([]Node, error) {
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

func ParseElements41(sc *bufio.Scanner) error {
	line := sc.Text()
	_ = strings.Split(line, " ")

	sc.Text()                            // read last $EndNodes line
	return errors.New("Not implemented") //TODO:
}
