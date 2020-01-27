package main

import(
	"github.com/jung-kurt/gofpdf"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	err := ConvertFirstPDF("PDF/hello.pdf")
	if err != nil {
		panic(err)
	}

	err = ConvertHTMLToPDF("PDF/HTMLToPDF.pdf")
	if err != nil {
		panic(err)
	}

	err = ConvertTextWithImageToPDF("PDF/TextWithImageToPDF.pdf")
	if err != nil {
		panic(err)
	}

	err = ConvertHTMLTableToPDF("PDF/HTMLTableToPDF.pdf")
	if err != nil {
		panic(err)
	}
}

func ConvertFirstPDF(fileName string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	return pdf.OutputFileAndClose(fileName)
}

func ConvertHTMLToPDF(fileName string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize()
	pdf.Write(lineHt, "To find out what's new in this tutorial, click ")
	pdf.SetFont("", "U", 0)
	link := pdf.AddLink()
	pdf.WriteLinkID(lineHt, "here", link)
	pdf.SetFont("", "", 0)

	pdf.AddPage()
	pdf.SetLink(link, 0, -1)
	pdf.SetFontSize(14)
	_, lineHt = pdf.GetFontSize()
	htmlStr := `You can now easily print text mixing different styles: <b>bold</b>, ` +
	    `<i>italic</i>, <u>underlined</u>, or <b><i><u>all at once</u></i></b>!<br><br>` +
	    `<center>You can also center text.</center>` +
	    `<right>Or align it to the right.</right>` +
	    `You can also insert links on text, such as ` +
	    `<a href="http://www.fpdf.org">www.fpdf.org</a>, or on an image: click on the logo.`
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, htmlStr)

	pdf.AddPage()
	return pdf.OutputFileAndClose(fileName)
}

func ConvertTextWithImageToPDF(filename string) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)

    image := "Image/logo.jpg"

    // CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
    pdf.CellFormat(190, 7, "PDF With Image", "0", 0, "CM", false, 0, "")

    // ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
    pdf.ImageOptions(
    	image, 20, 20, 0, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "",
    )

    pdf.AddPage()
    pdf.ImageOptions(
        image, 20, 20, 0, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "",
    )

    pdf.ImageOptions(
        image, 20, 150, 0, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "",
    )

    return pdf.OutputFileAndClose(filename)
}

func ConvertHTMLTableToPDF(filename string) error {
 	pdf := gofpdf.New("P", "mm", "A4", "")
	type countryType struct {
	    nameStr, capitalStr, areaStr, popStr string
	}

	countryList := make([]countryType, 0, 8)
	header := []string{}
	loadData := func(fileStr string) {
	    fl, err := os.Open(fileStr)
	    if err == nil {
	        scanner := bufio.NewScanner(fl)
	        var c countryType
	        for scanner.Scan() {
	            // Austria;Vienna;83859;8075
	            lineStr := scanner.Text()
	            list := strings.Split(lineStr, ";")

	            for keyData, listData := range list {
	            	data := strings.Split(listData, ":")
	            	
	            	if len(data) == 2 {
	            		switch keyData {
						    case 0:
						    	c.nameStr = data[1]
						    	header = nil
						    	header = append(header, data[0])
						    case 1:
						        c.capitalStr = data[1]
						        header = append(header, data[0])
						    case 2:
						        c.areaStr = data[1]
						        header = append(header, data[0])
						    case 3:
						        c.popStr = data[1]
						        countryList = append(countryList, c)
						        header = append(header, data[0])
						}
	            	} else {
	                	err = fmt.Errorf("error tokenizing %s", data)
	            	}
	            }
	        }

	        fl.Close()
	        if len(countryList) == 0 {
	            err = fmt.Errorf("error loading data from %s", fileStr)
	        }
	    }

	    if err != nil {
	        pdf.SetError(err)
	    }
	}

	// Simple table
	basicTable := func() {
	    left := (210.0 - 4*40) / 2
	    pdf.SetX(left)
	    for _, str := range header {
	        pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
	    }

	    pdf.Ln(-1)
	    for _, c := range countryList {
	        pdf.SetX(left)
	        pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
	        pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
	        pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
	        pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
	        pdf.Ln(-1)
	    }
	}

	// Better table
	improvedTable := func() {
	    // Column widths
	    w := []float64{40.0, 35.0, 40.0, 45.0}
	    wSum := 0.0
	    for _, v := range w {
	        wSum += v
	    }

	    left := (210 - wSum) / 2
	    // 	Header
	    pdf.SetX(left)
	    for j, str := range header {
	        pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
	    }

	    pdf.Ln(-1)
	    // Data
	    for _, c := range countryList {
	        pdf.SetX(left)
	        pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", false, 0, "")
	        pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", false, 0, "")
	        pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
	            "LR", 0, "R", false, 0, "")
	        pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
	            "LR", 0, "R", false, 0, "")
	        pdf.Ln(-1)
	    }

	    pdf.SetX(left)
	    pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}

	// Colored table
	fancyTable := func() {
	    // Colors, line width and bold font
	    pdf.SetFillColor(255, 0, 0)
	    pdf.SetTextColor(255, 255, 255)
	    pdf.SetDrawColor(128, 0, 0)
	    pdf.SetLineWidth(.3)
	    pdf.SetFont("", "B", 0)
	    // 	Header
	    w := []float64{40, 35, 40, 45}
	    wSum := 0.0
	    for _, v := range w {
	        wSum += v
	    }

	    left := (210 - wSum) / 2
	    pdf.SetX(left)
	    for j, str := range header {
	        pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
	    }

	    pdf.Ln(-1)
	    // Color and font restoration
	    pdf.SetFillColor(224, 235, 255)
	    pdf.SetTextColor(0, 0, 0)
	    pdf.SetFont("", "", 0)
	    // 	Data
	    fill := false
	    for _, c := range countryList {
	        pdf.SetX(left)
	        pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", fill, 0, "")
	        pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", fill, 0, "")
	        pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
	            "LR", 0, "R", fill, 0, "")
	        pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
	            "LR", 0, "R", fill, 0, "")
	        pdf.Ln(-1)
	        fill = !fill
	    }

	    pdf.SetX(left)
	    pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}

	loadData("File/countries.txt")
	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()
	basicTable()
	pdf.AddPage()
	improvedTable()
	pdf.AddPage()
	fancyTable()

	return pdf.OutputFileAndClose(filename)
}

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}
