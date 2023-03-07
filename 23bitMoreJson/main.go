package main

import (
	"encoding/json"
	"fmt"
)

// json: coursename sets the key name as coursename instead of Name when json is produced out of this struct...
// json:"-" removes the field in the resulting struct when json is created out of it
// json:"tags,omitempty will show the entry only n only if the firld is not null in json format
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json")
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON()  {
	lcoCourses := []course{
		{
			"ReactJs Bootcamp",
			299,
			"LearnCodeOnline.in",
			"abc123",
			[]string{
				"web-dev",
				"js",
			},
		},
		{
			"MERN Bootcamp",
			199,
			"LearnCodeOnline.in",
			"bcd123",
			[]string{
				"full-stack",
				"js",
			},
		},
		{
			"Angular Bootcamp",
			299,
			"LearnCodeOnline.in",
			"hit123",
			nil,
		},
	}

	// package this data as Json data

	// \t set the indentation to tab
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	
}

func DecodeJSON(){
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"Tags": ["web-dev","js"],
		"abc":123
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid.")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value pair
	var myonlineData map[string] interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for k,v := range myonlineData{
		fmt.Printf("key is %v and value is %v and type is: %T\n", k,v, v)
	}

}