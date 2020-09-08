package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"encoding/csv"
)

func main()  {
	
	column := 0
	var colname []string
	read := make( chan []string )

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
				colname = append(record) 
				column++
			} else {
				read <- record
			}
		}
	}( column )

	go func() {
		for {
			data, hasData := <-read

			if hasData {
				fmt.Println(data, colname)
			} else {
				fmt.Println("received all Data")
				return
			}

		}
	}()

	// for res := range colname { 
    //     fmt.Println(res) 
    // }

	// for res := range read { 
    //     fmt.Println(res) 
	// }
	<- read
}