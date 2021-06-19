package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/martinhackl/tqs/internal/lib"
)

type subsituteFlags []string

func (i *subsituteFlags) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *subsituteFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var substitutes subsituteFlags

	flag.Var(&substitutes, "s", "substitute a placeholder with given variable: `KEY=value`")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Required parameter path to config file missing")
		os.Exit(1)
	}

	configFile := flag.Arg(0)
	session, err := lib.ParseJSONFile(configFile)
	if err != nil {
		fmt.Println("An error occurred:", err)
		os.Exit(1)
	}

	if len(substitutes) > 0 {
		sTuples := make(map[string]string)
		for _, s := range substitutes {
			kv := strings.Split(s, "=")
			sTuples[kv[0]] = kv[1]
		}
		lib.Substitute(session, sTuples)
	}

	if err := lib.CreateSession(*session); err != nil {
		fmt.Println("error while creating session:", err)
		os.Exit(1)
	}
}
