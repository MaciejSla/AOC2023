package main

import "github.com/docopt/docopt-go"

func main() {
	usage := `
Usage: Day3 <ex> <input file>
`
	args, _ := docopt.ParseDoc(usage)
	ex, _ := args.Int("<ex>")
	input, _ := args.String("<input file>")

	switch ex {
	case 1:
		Ex1(input)
		//case 2:
		//	Ex2(input)
	}
}
