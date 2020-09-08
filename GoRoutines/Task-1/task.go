package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"encoding/csv"
	"encoding/json"
)

func main()  {
	
	forever := make(chan bool, 1)

	column := 0
	read := make( chan []string )
	final := make( chan string )
	parseData := make([]map[string]interface{}, 0, 0)

	csvfile, err := os.Open("dfy.csv")	
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader( csvfile )

	go func( column int ) {
		for {
			record, err := r.Read()

			if err == io.EOF {
				close(read)
				break
			}

			if err != nil {
				log.Fatal(err)
				close(read)
			}

			if column == 0 {
				column++
			} else {
				read <- record
			}
		}
	}( column )

	go func() {
		for {
			data, hasData := <-read
			var singleMap = make(map[string]interface{})
			if hasData {
				singleMap["id"] = data[0]
				singleMap["name"] = data[1]
				singleMap["description"] = data[2]
				singleMap["image_url"] = data[3]
				singleMap["color"] = data[4]
				parseData = append(parseData, singleMap)
			} else {
				b, _:= json.Marshal(parseData)
				fmt.Println("received all Data")				
				final <- string(b)
				close(final)
				return
			}

		}
	}()

	fmt.Println(<-final)
	<- forever
}