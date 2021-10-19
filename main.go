package main

import (
	u "Golang-HTML-TO-PDF-Converter/pdfGenerator"
	"fmt"
)

func main() {

	r := u.NewRequestPdf("")

	//html template path
	templatePath := "templates/demo.html"

	//path for download pdf
	outputPath := "storage/pranjit4.pdf"

	//html template data
	templateData := struct {
		// Title1       string
		// Description1 string
		// Company1     string
		// Contact1     string
		// Country1     string
		// Address1     string
		//	Url            string
		//RespondentName string
		// CustomerName   string
		// Method1        string
		// MethodType1    string
		// Desc           string
		AlfredsFutterkiste string
		MariaAnders        string
		Germany            string
	}{
		// Title1:       "HTML to PDF generator",
		// Description1: "This is the simple HTML to PDF file.",
		// Company1:     "Jhon Lewis",
		// Contact1:     "Maria Anders",
		// Country1:     "Germany",
		// Address1:     "assam",
		//Url:            "add",
		//	RespondentName: "pranjit dgfdfg ",
		// CustomerName:   "ashu",
		// Method1:        "unic",
		// MethodType1:    "rahul",
		// Desc:           "rabi",
		AlfredsFutterkiste: "HI",
		MariaAnders:        "hello",
		Germany:            "yes",
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
