package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bww/go-util/v1/crypto"
	"github.com/bww/go-util/v1/debug"
)

func signMessage(args []string) {
	cmdline := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	globals := bindGlobalFlags(cmdline)
	cmdline.Parse(args)

	debug.DEBUG = *globals.Debug
	debug.VERBOSE = *globals.Verbose
	debug.TRACE = *globals.Trace

	if *globals.Secret == "" {
		fmt.Println("*** No secret provided (use -secret <data>)")
		return
	}
	if *globals.Salt == "" {
		fmt.Println("*** No salt provided (use -salt <data>)")
		return
	}

	key := crypto.GenerateKey(*globals.Secret, *globals.Salt, crypto.SHA1)

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("*** Could not read resource policy set: %v\n", err)
		return
	}

	enc := base64.StdEncoding.EncodeToString(data)
	sig := crypto.Sign(key, crypto.SHA256, []byte(enc))

	fmt.Println("---")
	fmt.Println(sig + sep + enc)
}

func verifyMessage(args []string) {
	cmdline := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	globals := bindGlobalFlags(cmdline)
	cmdline.Parse(args)

	debug.DEBUG = *globals.Debug
	debug.VERBOSE = *globals.Verbose
	debug.TRACE = *globals.Trace

	if *globals.Secret == "" {
		fmt.Println("*** No secret provided (use -secret <data>)")
		return
	}
	if *globals.Salt == "" {
		fmt.Println("*** No salt provided (use -salt <data>)")
		return
	}

	key := crypto.GenerateKey(*globals.Secret, *globals.Salt, crypto.SHA1)

	for i, a := range cmdline.Args() {
		parts := strings.SplitN(a, sep, 2)
		if len(parts) != 2 {
			fmt.Printf("# % 2d *** Malformed message\n", i+1)
			continue
		}

		if !crypto.Verify(key, crypto.SHA256, parts[0], []byte(parts[1])) {
			fmt.Printf("# % 2d *** Could not validate\n", i+1)
			continue
		}

		dec, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			fmt.Printf("# % 2d *** Could not decode: %v\n", i+1, err)
			continue
		}

		fmt.Printf("# % 2d Valid: %s\n", i+1, string(dec))
	}
}
