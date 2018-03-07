package main

import (
	"fmt"
	"os"

	"github.com/giuliocalzo/eks-token"


)


func main() {
    // fmt.Println("start")

    if len(os.Args) != 3 {
  		fmt.Fprintf(os.Stderr,"error argv required: e.g. %v roleARN clusterID \n", os.Args[0])
      os.Exit(1)
  	}

    roleARN := os.Args[1]
    if roleARN == "" {
			fmt.Fprintf(os.Stderr, "roleARN not valid: \n")
			os.Exit(1)
		}
    clusterID := os.Args[2]
    if clusterID == "" {
			fmt.Fprintf(os.Stderr, "clusterID not valid: \n")
			os.Exit(1)
		}
		var tok string
		var err error
		gen, err := token.NewGenerator()
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not get token: %v\n", err)
			os.Exit(1)
		}
		tok, err = gen.GetWithRole(clusterID, roleARN)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not get token: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(tok)
}
