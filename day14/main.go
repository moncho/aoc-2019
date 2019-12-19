package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	fuel = "FUEL"
	ore  = "ORE"
)

type chemical struct {
	name   string
	amount int
}

type reaction struct {
	input  map[string]chemical
	output chemical
}

type reactions struct {
	reactions []reaction
	indices   map[string]int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reactions := parseInstructions(file)

	fmt.Printf("what is the minimum amount of ORE required to produce exactly 1 FUEL? %d\n", reactions.ore(fuel))

}
func parseInstructions(r io.Reader) reactions {
	var rr []reaction
	scanner := bufio.NewScanner(r)
	indices := make(map[string]int)
	i := 0
	for scanner.Scan() {
		reaction := parse(scanner.Text())
		rr = append(rr, reaction)
		indices[reaction.output.name] = i
		i++
	}
	return reactions{
		reactions: rr,
		indices:   indices,
	}
}
func parse(s string) reaction {
	r := reaction{
		input: make(map[string]chemical),
	}
	formula := strings.Split(s, "=>")

	for _, input := range strings.Split(formula[0], ",") {
		chem := parseChemical(input)
		r.input[chem.name] = chem
	}
	r.output = parseChemical(formula[1])

	return r
}

func parseChemical(s string) chemical {
	chem := strings.Split(strings.TrimSpace(s), " ")
	amount, err := strconv.Atoi(chem[0])
	if err != nil {
		panic(err)
	}

	return chemical{
		name:   chem[1],
		amount: amount,
	}
}

func (rr reactions) ore(s string) int {
	baseElements := make(map[string]int)
	rr.base(rr.reactions[rr.indices[s]].output, baseElements)

	ores := 0
	for element, count := range baseElements {
		reaction := rr.reactions[rr.indices[element]]
		mul := 1
		if count > reaction.output.amount {
			mul = (count + reaction.output.amount - 1) / reaction.output.amount
		}
		ores += mul * reaction.input[ore].amount
	}
	return ores
}

func (rr reactions) ore2(s string) int {

	chem := map[string]int{
		rr.reactions[rr.indices[s]].output.name: rr.reactions[rr.indices[s]].output.amount,
	}

	for {
		if len(chem) == 1 && chem[ore] > 1 {
			break
		}
		for name, amount := range chem {
			if name == ore {
				continue
			}
			reaction := rr.reactions[rr.indices[name]]

			factor := (amount + reaction.output.amount - 1) / reaction.output.amount
			chem[name] -= factor * reaction.output.amount
			for _, input := range reaction.input {
				chem[input.name] += factor * input.amount
			}
			if chem[name] <= 0 {
				delete(chem, name)
			}
		}
	}
	return chem[ore]
}

func (rr reactions) base(chem chemical, base map[string]int) {
	reaction := rr.reactions[rr.indices[chem.name]]
	for _, input := range reaction.input {
		if input.name == ore {
			base[chem.name] += chem.amount
		} else {
			mul := (chem.amount + reaction.output.amount - 1) / reaction.output.amount

			rr.base(chemical{
				name:   input.name,
				amount: input.amount * mul,
			}, base)
		}
	}
}
