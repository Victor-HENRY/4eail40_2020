package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}
func (tr spaceEraser) Read(bufer []byte) (int, error) {
    n, error := tr.r.Read(bufer)
    var maliste []string
    sli := strings.Split(string(bufer), "")
    for _, el := range sli {
        if el != " " {
            maliste = append(maliste, el)
        }
    }
    final := []byte(strings.Join(maliste, ""))
    copy(bufer, final)
    return n, error
}


func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
