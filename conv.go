package main

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"os"
)

func convToCSV(Vulns []Vulnerability){
	file , err:= os.Create("new.csv")
	if err != nil{
		fmt.Println(err)
	}


	w := csv.NewWriter(file)

	t := reflect.TypeOf(Vulns[0])

	header := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		header[i] = t.Field(i).Name
	}
	w.Write(header)

	for _, p := range Vulns {
		data := make([]string, t.NumField())
		val := reflect.ValueOf(p)
		for i := 0; i < t.NumField(); i++ {
			data[i] = fmt.Sprintf("%v", val.Field(i))
		}
		w.Write(data)
	}

	w.Flush()

}
