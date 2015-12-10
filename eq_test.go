package goquery_test

import (
	"fmt"
	"github.com/eaciit/goquery"
	"testing"
)

type Data map[string]interface{}
type ColumnSetting struct {
	Title    string
	Selector string
}

func TestDirectEq(t *testing.T) {
	ds := []Data{}
	settings := []ColumnSetting{
		{"Code", "td:eq(0)"},
		{"ListingDate", "td:eq(1)"},
		{"Price", "td:eq(5)"},
	}

	doc, _ := goquery.NewDocument("http://www.shfe.com.cn/en/products/Gold/")
	rows := doc.Find(".listshuju:eq(0) tr")
	rows.Each(func(i int, sel *goquery.Selection) {
		d := Data{}
		for _, s := range settings {
			d[s.Title] = sel.Find(s.Selector).Text()
		}
		ds = append(ds, d)
	})

	fmt.Printf("Data Grabbed:\n%v\n", ds)
}
