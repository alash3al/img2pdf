package main

import "fmt"
import "flag"
import "path/filepath"
import "github.com/signintech/gopdf"

var (
	src = flag.String("src", "./", "the images directory path (png, jpg & jpeg) files")
	as  = flag.String("as", "./result.pdf", "the result filename")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("Reading files from (", *src, ") and saving the result as (", *as,")")
	fmt.Println("-----------------------")
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: UnitPT, PageSize: gopdf.Rect{W: 595.28, H: 841.89}})
	jpgs, _ := filepath.Glob(*src + "/*.jpg")
	pngs, _ := filepath.Glob(*src + "/*.png")
	jpegs, _ := filepath.Glob(*src + "/*.jpeg")
	files := append(jpgs, append(jpegs, pngs...)...)
	for i := 0; i < len(files); i++ {
		fmt.Println(i+1, ")- adding ", files[i])
		x := float64(0)
		if x < 0 {
			continue
		}
		pdf.AddPage()
		pdf.Image(files[i], x, 0, &gopdf.Rect{W: 595.28, H: 840})
	}
	fmt.Println("saving to ", *as, " ...")
	pdf.WritePdf(*as)
	fmt.Println("-----------------------")
	fmt.Println("Done, have fun ;)")
	fmt.Println("Created by Mohammed Al Ashaal <https://www.alash3al.xyz>")
}
