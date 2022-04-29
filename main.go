package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/tredoe/osutil/user/crypt"
	"github.com/tredoe/osutil/user/crypt/md5_crypt"
	"github.com/tredoe/osutil/user/crypt/sha256_crypt"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var (
	Name    string
	Version string
	GitHash string
)

const (
	Prefix_1    = "$1$"
	Prefix_5    = "$5$"
	Prefix_6    = "$6$"
)

var (
	V     = flag.Bool("v", false, "show version")
	H     = flag.Bool("h", false, "show help and exit")
	_1    = flag.Bool("1", false, "use MD5 based Unix password algorithm 1")
	_5    = flag.Bool("5", false, "use SHA256 based Unix password algorithm 5")
	_6    = flag.Bool("6", false, "use SHA512 based Unix password algorithm 6 (default)")
	S     = flag.String("s", "", "salt")
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [password]\n", filepath.Base(os.Args[0]))
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "options\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if !flag.Parsed() {
		flag.Usage()
		return
	}

	if *H {
		flag.Usage()
		return
	}

	if *V {
		fmt.Printf("%s version %s (%s)\n", Name, Version, GitHash)
		return
	}

	// prefix: algorithm type
	var prefix = Prefix_6
	if *_5 {
		prefix = Prefix_5
	} else if *_1 {
		prefix = Prefix_1
	}

	// password
	var password []byte
	if flag.NArg() > 0 {
		password = []byte(flag.Arg(0))
	} else {
		var err error
		password, err = readPassword()
		if err != nil {
			log.Fatal(err)
		}
	}

	// salt: length must be 8~16 bytes
	var salt []byte
	if *S != "" {
		if len(*S) < 8 || len(*S) > 16 {
			log.Fatal("salt length must be 8~16 bytes")
		}
		salt = []byte(prefix + *S)
	} else {
		// generate random salt
		salt = append([]byte(prefix), randomSalt(16)...)
	}

	// create hash
	var h crypt.Crypter
	switch prefix {
	case Prefix_1:
		h = md5_crypt.New()
	case Prefix_5:
		h = sha256_crypt.New()
	default:
		h = sha512_crypt.New()
	}
	hash, err := h.Generate(password, salt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash)
}

func readPassword() (password []byte, err error) {
	fmt.Print("Enter password: ")
	stdin := int(os.Stdin.Fd())
	password, err = terminal.ReadPassword(stdin)
	fmt.Println()
	if err != nil {
		return
	}

	fmt.Print("Confirm password: ")
	confirm, err := terminal.ReadPassword(stdin)
	fmt.Println()
	if err != nil {
		return
	}
	if !bytes.Equal(confirm, password) {
		return nil, errors.New("password mismatch")
	}
	return
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
// generate alphanumeric salt
func randomSalt(n int) []byte {
	result := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return result
}
