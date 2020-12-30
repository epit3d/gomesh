package gomesh

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Parse22(filename string) (*Msh, error) {
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
			msh.Nodes, err = ParseNodes22(nextLineReader(scanner))
		case "$Elements":
			msh.Elements, err = ParseElements22(nextLineReader(scanner))
		case "$ElementData":
			var elData map[string][]string
			elData, err = ParseElementData22(nextLineReader(scanner))
			msh.SetElementData(elData)
		default:
			log.Printf("WARN: Tag %v not implemented, skipping", line)
		}
		if err != nil {
			log.Fatalf("Error in file parsing22: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &msh, nil
}

func ParseNodes22(nextLine nextLineFunc) ([]Node, error) {
	line := nextLine()
	numNods, err := strconv.Atoi(line)
	if err != nil {
		return nil, errors.New("invalid numNods: " + line)
	}
	var nodes []Node
	for i := 0; i < numNods; i++ {
		line := nextLine()
		nod := strings.Split(line, " ")
		if len(nod) != 4 {
			return nil, errors.New("nod len!=4: " + line)
		}
		nodes = append(nodes, NewNode(nod[0], nod[1], nod[2], nod[3]))
	}
	nextLine() // read last $EndNodes line
	return nodes, nil
}

func ParseElements22(nextLine nextLineFunc) ([]Element, error) {
	line := nextLine()
	numEls, err := strconv.Atoi(line)
	if err != nil {
		return nil, errors.New("invalid numEls: " + line)
	}
	//1 4 0 203 170 206 163
	//2 4 0 755 313 307 320
	//3 4 0 96 85 62 812
	//elm-number elm-type number-of-tags < tag > … node-number-list
	var els []Element
	for i := 0; i < numEls; i++ {
		line := nextLine()
		el := strings.Split(line, " ")
		elTag := el[0]
		elType := el[1]
		numTags, err := strconv.Atoi(el[2])
		if err != nil {
			return nil, errors.New("invalid numTags: " + line)
		}
		elTags := el[3 : 3+numTags]
		elNodes := el[3+numTags:]
		els = append(els, NewElement(elTag, elType, elTags, elNodes))
	}
	nextLine() // read last $EndElements line
	return els, nil
}

func ParseElementData22(nextLine nextLineFunc) (map[string][]string, error) {
	/*	$ElementData
		number-of-string-tags
		< "string-tag" >
		…
		number-of-real-tags
		< real-tag >
		…
		number-of-integer-tags
		< integer-tag >
		…
		elm-number value …
		…
		$EndElementData*/
	line := nextLine()
	numStringTags, err := strconv.Atoi(line)
	if err != nil {
		return nil, errors.New("invalid numStringTags: " + line)
	}
	var strTags []string
	for i := 0; i < numStringTags; i++ {
		str := nextLine()
		str = strings.ReplaceAll(str, "\"", "")
		strTags = append(strTags, str)
	}
	//The real tags typically only contain one value, the time.
	//Discard it.
	line = nextLine()
	numRealTags, err := strconv.Atoi(line)
	if err != nil {
		return nil, errors.New("invalid numRealTags: " + line)
	}
	//var realTags []float64
	for i := 0; i < numRealTags; i++ {
		nextLine()
		//realTags = append(realTags, sc.Text())
	}
	line = nextLine()
	numIntTags, err := strconv.Atoi(line)
	if err != nil {
		return nil, errors.New("invalid numIntTags: " + line)
	}
	if numIntTags < 3 {
		return nil, errors.New("invalid numIntTags, less 3: " + line)
	}
	var intTags []int
	for i := 0; i < numIntTags; i++ {
		line = nextLine()
		v, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.New("invalid int value in data: " + line)
		}
		intTags = append(intTags, v)
	}
	//numValues := intTags[1]
	numData := intTags[2]
	elData := map[string][]string{}
	for i := 0; i < numData; i++ {
		line := nextLine()
		el := strings.Split(line, " ")
		elData[el[0]] = el[1:]
	}
	nextLine() // read last $EndElementData line
	return elData, nil
}
