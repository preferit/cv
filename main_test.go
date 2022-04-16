package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	os.Args = []string{"cv",
		"-cv", "example.yaml",
		"-co", "preferit.yaml",
		"-s", "cv.html",
	}
	main()

	got, err := os.ReadFile("cv.html")
	if err != nil {
		t.Fatal(err)
	}

	cases := []string{
		"One sentence description NASA", // last one
	}
	for _, exp := range cases {
		if !bytes.Contains(got, []byte(exp)) {
			t.Fatal("missing:", exp)
		}
	}
	os.RemoveAll("cv.html")
}
