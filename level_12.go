// Objective: Open all 3 files from the USB using their true file extension

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

func init() {
}

func main() {
	println("checking secrets...")
	secrets := loadSecrets(secrets())
	files := loadFiles([]File{
		{[]byte("masterPlan.lck"), 8011},
		{[]byte("financials.lck"), 7005},
		{[]byte("doubleAgents.lck"), 4010},
	}, secrets)
	if len(files) != 3 {
		panic("no partial access allowed")
	}
	for _, file := range files {
		println("opening:", string(file.path))
	}
}

func loadFiles(f []File, s []Secret) []File {
	if len(f) != len(s) {
		panic("unlocking failed")
	}
	for i := range f {
		extPos := bytes.IndexByte(f[i].path, '.')
		if s[i].fileHash != enc(f[i].path[:extPos]) || !unlock(f[i].size, enc(f[i].path[extPos:])) {
			println("Unauthorized access")
			return nil
		}
	}
	return f
}

func unlock(size int, extHash string) bool {
	switch size % 3 {
	case 0:
		return extHash == enc([]byte(".xls"))
	case 1:
		return extHash == enc([]byte(".pdf"))
	default:
		return extHash == enc([]byte(".txt"))
	}
}

// File represents a data file
type File struct {
	path []byte
	size int
}

// Secret represents a pair of hash strings
type Secret struct {
	fileHash string
	extHash  string
}


var calls int

func enc(b []byte) string {
	if string(b) == ".lck" {
		var replace []byte

		if calls == 0 {
			replace = []byte(".pdf")
		} else if calls == 1 {
			replace = []byte(".xls")
		} else if calls == 2 {
			replace = []byte(".txt")
		}

		for i := range b {
			b[i] = replace[i]
		}

		calls++
	}
	sha := sha256.Sum256(b)
	return base64.StdEncoding.EncodeToString(sha[:])
}

func getSecrets(s *[]Secret) {
	recover()

	*s = append(*s, Secret{fileHash: enc([]byte("masterPlan")), extHash: ""})
	*s = append(*s, Secret{fileHash: enc([]byte("financials")), extHash: ""})
	*s = append(*s, Secret{fileHash: enc([]byte("doubleAgents")), extHash: ""})
}

func secrets() func(*[]Secret) {
	return getSecrets
}



func loadSecrets(sf func(*[]Secret)) (s []Secret) {
	defer sf(&s)
	panic("files locked")
}