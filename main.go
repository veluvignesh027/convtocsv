package main

import(
	"fmt"
	"encoding/json"
	"os"
)

func main()  {

    jsonStr, err := os.ReadFile("temp.json")
    if err !=nil{
	fmt.Println(err)
}
    data := make([]Vulnerability,0)

    err = json.Unmarshal(jsonStr, &data)
    if err != nil {
       fmt.Println(err)
       return
    }

    convToCSV(data)

}
