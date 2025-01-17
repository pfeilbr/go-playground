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

// Person defines Person and it's data
type Person struct {
	Name      string
	Age       int
	Nicknames []string
}

// Description description
func (p *Person) Description() string {
	return fmt.Sprintf("Name: %s, Age: %d, Nicknames: %s", p.Name, p.Age, strings.Join(p.Nicknames, ","))
}

// SetName add a setter.  note that p is a pointer and must be to change it's state
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
	fmt.Printf("host = %v, useTLS = %v", *host, *useTLS)

	// execute external program
	cmd := exec.Command("/bin/ls", "-l")
	buf, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error %v", err)
	}
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

	// ---
	// type assertions. similar to casting. interface to more specialized type
	var general interface{} // variable that can hold anything
	general = "brian"       // assing a string

	if name, ok := general.(string); ok { // cast
		fmt.Printf("name = %v\n", name)
	} else {
		fmt.Println(`failed to cast var general to string`)
	}

	// assign a struct
	general = &Person{"Trica", 42, []string{"Trish", "Dish"}}

	if tricia, ok := general.(*Person); ok { // cast.  note pointer dereference with *
		fmt.Printf("tricia = %v\n", tricia)
	} else {
		fmt.Println(`failed to cast var general to *Person`)
	}
	// ---

}
