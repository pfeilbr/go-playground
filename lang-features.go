package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// define Person and it's data
type Person struct {
	Name      string
	Age       int
	Nicknames []string
}

// add a Person method
func (p *Person) Description() string {
	return fmt.Sprintf("Name: %s, Age: %d, Nicknames: %s", p.Name, p.Age, strings.Join(p.Nicknames, ","))
}

// add a setter.  note that p is a pointer and must be to change it's state
func (p *Person) SetName(name string) {
	p.Name = name
}

func main() {
	p := &Person{"Brian", 36, []string{"Bri", "b-dog", "Brizilla"}}
	p.SetName("Bryan")
	fmt.Printf("p.Description(): %s", p.Description())

	// fetch the body of a web page
	resp, err := http.Get("http://google.com/")
	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)

	// access the response headers
	for key, val := range resp.Header {
		fmt.Printf("%s: %s\n\n", key, val)
	}

	// display current working directory
	wd, _ := os.Getwd()
	fmt.Printf("wd = %s", wd)

	// print commad line args
	fmt.Printf("os.Args = %s", strings.Join(os.Args, ","))

	// read the contents of a file
	path := "/Users/brianpfeil/Dropbox/address_state.txt"
	body, _ := ioutil.ReadFile(path)
	fmt.Printf("contents of \"%s\": \n%s\n\n", path, string(body))

	// append element to array example
	names := []string{"Brian", "Tricia", "Wyatt"}
	names = append(names, "Max")
	fmt.Printf("names = %s\n", strings.Join(names, ","))

	// command line flags
	var host = flag.String("host", ":8080", "host to bind to")
	var useTLS = flag.Bool("tls", false, "enable TLS")
	flag.Parse()
	fmt.Printf("host = %s, useTLS = %s", *host, *useTLS)

	// execute external program
	cmd := exec.Command("/bin/ls", "-l")
	buf, err := cmd.Output()
	fmt.Printf("buf = %v", string(buf))

	// JSON examples
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	fmt.Printf("f = %v\n", m)

	wrd := []byte{'a', 'b', 'c'}
	fmt.Println(wrd)
	fmt.Println(string(wrd))

}
