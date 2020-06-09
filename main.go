package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	//emailList := make([]string, 10)
	//sendEmail(emailList)
	allLawyers := readLawyers("lawyers.csv")
	fmt.Println(match(allLawyers, "NA"))
}

func readLawyers(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': '%s\n", name, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows
}

func match(rows [][]string, location string) []string {
	var lawyers []string
	for i := range rows {
		if i != 0 {
			lawyers = append(lawyers, rows[i][3])
		}

	}
	return lawyers
}

func sendEmail([]string) {
	from := mail.NewEmail("Me", "austinlh.business@gmail.com")
	to := mail.NewEmail("Austin", "austinlh.business@gmail.com")
	subject := "Something here"
	plainTextContent := "something"
	htmlContent := "<strong>Something"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("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"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}