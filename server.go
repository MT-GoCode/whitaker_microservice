package main

import (
	"bufio"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
    "io"
	"os/exec"
	"net/http" // http package
	// "reflect"
)
// Send a request to /process with text query parameter containing the word to be processed, result is sent as response
// running python file from go based off of this:
// https://stackoverflow.com/questions/41415337/running-external-python-in-golang-catching-continuous-exec-command-stdout

var textile string;

func hello(w http.ResponseWriter, req *http.Request) { // handler for hello path
	// arguments are http.ResponseWriter to write the response and a pointer to the http.Request
	// var names come before type

	// https://www.informit.com/articles/article.aspx?p=2861456&seqNum=6 
	// read this to see how to respond to different REST requests
	// basically its just accessing req.Method value within this function , which will contain the request type (GET, POST, PUT, DELETE, etc.)
	fmt.Fprintf(w, "hello\n")
	// simply writes hello TO the http.responseWriter - similar to py print(text, file=)
}

func runPy(w http.ResponseWriter, r *http.Request) { // handler for hello path
	
	// b, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// // Unmarshal
	// var res textrequest
	// err = json.Unmarshal(b, &res)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// output, err := json.Marshal(res)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	keys, ok := r.URL.Query()["text"]
	
	fmt.Println(keys[0], ok)
	
	if r.Method == "GET" {
		results := runFile(keys[0])
		fmt.Fprintf(w, results) //results)
	}
}



func headers(w http.ResponseWriter, req *http.Request) {

	
	for name, headers := range req.Header { // extracts name and header.
		// short variable declaration := is for implicit type, only available in functions

        for _, h := range headers {  // goes through the header
			fmt.Fprintf(w, "%v: %v\n", name, h) //  and prints to the response all this info
        }
    }
}

func main() {
	
	fmt.Printf("running file\n")

	http.HandleFunc("/hello", hello) // 'attaches' the hello handler to /"hello" endpoint
	// that creates a default server so in line 37, nil refers to the default server
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/runPy", runPy)

	http.ListenAndServe(":8090", nil)
	
	// to run this boi
	// go run server.go
	// curl localhost:8090/hello
	// or you can go into web browser and send the request
	
}


func runFile(text string) string {
    cmd := exec.Command("python", "./open_words/execute.py", text)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
		panic(err)
		return ""
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
		
		panic(err)
		return ""
    }
    err = cmd.Start()
    if err != nil {
		
		panic(err)
		return ""
    }
	
	copyOutput(stderr)
    // cmd.Wait()
	return copyOutput(stdout)
}

func copyOutput(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	var textll string;
    for scanner.Scan() {
		textll = scanner.Text()
		
	}
	return textll
}

// package main

// import (
// 	"bufio"
// 	// "encoding/json"
// 	"fmt"
// 	// "io/ioutil"
//     "io"
// 	"os/exec"
// 	"net/http" // http package
// 	// "reflect"
// )

// var textile string;

// func hello(w http.ResponseWriter, req *http.Request) { // handler for hello path
// 	// arguments are http.ResponseWriter to write the response and a pointer to the http.Request
// 	// var names come before type

// 	// https://www.informit.com/articles/article.aspx?p=2861456&seqNum=6 
// 	// read this to see how to respond to different REST requests
// 	// basically its just accessing req.Method value within this function , which will contain the request type (GET, POST, PUT, DELETE, etc.)
// 	fmt.Fprintf(w, "hello\n")
// 	// simply writes hello TO the http.responseWriter - similar to py print(text, file=)
// }

// type textrequest struct {
//     text string `json:"text"`
// }

// func runPy(w http.ResponseWriter, r *http.Request) { // handler for hello path
	
// 	// b, err := ioutil.ReadAll(r.Body)
// 	// defer r.Body.Close()
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), 500)
// 	// 	return
// 	// }
// 	// // Unmarshal
// 	// var res textrequest
// 	// err = json.Unmarshal(b, &res)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), 500)
// 	// 	return
// 	// }

// 	// output, err := json.Marshal(res)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), 500)
// 	// 	return
// 	// }
// 	keys, ok := r.URL.Query()["text"]
	
// 	fmt.Println(keys[0], ok)
	
// 	if r.Method == "GET" {

// 		fmt.Fprintf(w, "hello\n")
// 		// runFile()
// 	}
	
// }



// func headers(w http.ResponseWriter, req *http.Request) {

	
// 	for name, headers := range req.Header { // extracts name and header.
// 		// short variable declaration := is for implicit type, only available in functions

//         for _, h := range headers {  // goes through the header
// 			fmt.Fprintf(w, "%v: %v\n", name, h) //  and prints to the response all this info
//         }
//     }
// }

// func main() {
// 	runFile()
// 	fmt.Printf(textile)
// 	fmt.Printf("running file\n")

// 	http.HandleFunc("/hello", hello) // 'attaches' the hello handler to /"hello" endpoint
// 	// that creates a default server so in line 37, nil refers to the default server
// 	http.HandleFunc("/headers", headers)
// 	http.HandleFunc("/runPy", runPy)

// 	http.ListenAndServe(":8090", nil)
	
// 	// to run this boi
// 	// go run server.go
// 	// curl localhost:8090/hello
// 	// or you can go into web browser and send the request
	
// }


// func runFile() {
//     cmd := exec.Command("python", "./open_words/execute.py")
//     stdout, err := cmd.StdoutPipe()
//     if err != nil {
		
// 		panic(err)
// 		// return ""
//     }
//     stderr, err := cmd.StderrPipe()
//     if err != nil {
		
// 		panic(err)
// 		// return ""
//     }
//     err = cmd.Start()
//     if err != nil {
		
// 		panic(err)
// 		// return ""
//     }

// 	copyOutput(stdout)
	

// 	copyOutput(stderr)
// 	cmd.Wait()
//     // return text
    
// }

// func copyOutput(r io.Reader)  {
// 	scanner := bufio.NewScanner(r)
//     for scanner.Scan() {
// 		textile = scanner.Text()
		
// 	}
// 	// return toRet
// }