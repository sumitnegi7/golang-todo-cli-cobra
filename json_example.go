package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// We’ll use these two structs to demonstrate encoding and decoding
// of custom types below.
type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON.
// Fields must start with capital letters to be exported.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func m1ain(){
	bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
	slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))


	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple","peach","pear"}}
	res1B,_ := json.Marshal(res1D)
	fmt.Println(string(res1B))


	// You can use tags on struct field declarations to customize 
	// the encoded JSON key names. Check the definition of response2 
	// above to see an example of such tags.
	res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

	// decoding store data

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	err := json.Unmarshal(byt,&dat);
	if err !=nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)


	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res:= response2{}
	json.Unmarshal([]byte(str),&res)
	fmt.Println(res)
    fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out. We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}