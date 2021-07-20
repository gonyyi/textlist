package main

import (
	"fmt"
	"github.com/gonyyi/textlist"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <mode> <opt>...\n", os.Args[0])
		return
	}

	switch os.Args[1] {
	case "compare", "cmp":
		if len(os.Args) > 3 {
			added, removed, err := compare(os.Args[2], os.Args[3])
			if err != nil {
				println(err.Error())
				return
			}
			for _, v := range added {
				println("+ " + v)
			}
			for _, v := range removed {
				println("- " + v)
			}
		} else {
			fmt.Printf("Usage: %s %s <FILE:FROM> <FILE:TO>\n", os.Args[0], os.Args[1])
		}
	default:
		println("Unknown")
	}
}

func compare(fileFrom, fileTo string) (added []string, removed []string, err error) {
	lFrom, err := textlist.NewListFromFile(fileFrom, textlist.FILE_TRIMSPACE|textlist.FILE_NEWLINE)
	if err != nil {
		return nil, nil, err
	}
	lTo, err := textlist.NewListFromFile(fileTo, textlist.FILE_TRIMSPACE|textlist.FILE_NEWLINE)
	if err != nil {
		return nil, nil, err
	}
	added, removed = textlist.Compare(lFrom, lTo)
	return
}
