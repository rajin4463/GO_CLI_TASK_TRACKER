package utils

import (
	"github.com/markkurossi/tabulate"
	"os"
)

func FormatTab(data [][]string){
	tab := tabulate.New(tabulate.Github)
	for index, _ := range data[0]{
		tab.Header(data[0][index]).SetAlign(tabulate.MR)
	}
	for i := 1; i < len(data); i++{
		row := tab.Row()
		for j, _ := range data[i]{
			row.Column(data[i][j])
		}
	}
	tab.Print(os.Stdout)
}