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

    PanzerVulns := make([]Panzerformat,0)
    for _, item := range data{
        PanzerVulns = append(PanzerVulns,Panzerformat{item.VulnerabilityID,item.SeveritySource,item.LastModifiedDate,item.PublishedDate,item.Severity,item.PkgName,item.InstalledVersion,item.FixedVersion})
    }


    convToCSV(PanzerVulns)

}
