# Go Markdown Table

Package gomdtable package is needed for generating Markdown tables.

## Install

```bash
go get github.com/probeldev/gomdtable
```


## Usage
example:

```go
package main

import (
	"fmt"
	"log"

	"github.com/probeldev/gomdtable"
)

func main() {
	fmt.Println("Table:")
	writeTable()

	fmt.Println("Table list:")
	writeTableList()
}

func writeTableList() {
	table := gomdtable.NewTable()
	err := table.SetHeader([]string{"Имя", "Возраст"})
	if err != nil {
		log.Panic(err)
	}

	err = table.AddRow([]string{"Сергей", "31"})
	if err != nil {
		log.Panic(err)
	}
	err = table.AddRow([]string{"Ян", "9"})
	if err != nil {
		log.Panic(err)
	}
	err = table.AddRow([]string{"Яна", "110"})
	if err != nil {
		log.Panic(err)
	}

	table2 := gomdtable.NewTable()
	err = table2.SetHeader([]string{"Имя", "Возраст"})
	if err != nil {
		log.Panic(err)
	}

	err = table2.AddRow([]string{"Сергей Игоревич", "31"})
	if err != nil {
		log.Panic(err)
	}
	err = table2.AddRow([]string{"Ян Васильев", "9"})
	if err != nil {
		log.Panic(err)
	}
	err = table2.AddRow([]string{"Яна Александровна", "110"})
	if err != nil {
		log.Panic(err)
	}

	tableList, err := gomdtable.GenerateTableList([]gomdtable.Table{
		table,
		table2,
	})
	if err != nil {
		log.Panic(err)
	}

	for _, tableStr := range tableList {
		fmt.Println(tableStr)
	}

}

func writeTable() {

	table := gomdtable.NewTable()
	err := table.SetHeader([]string{"Имя", "Возраст"})
	if err != nil {
		log.Panic(err)
	}

	err = table.AddRow([]string{"Сергей", "31"})
	if err != nil {
		log.Panic(err)
	}
	err = table.AddRow([]string{"Ян", "9"})
	if err != nil {
		log.Panic(err)
	}
	err = table.AddRow([]string{"Яна", "110"})
	if err != nil {
		log.Panic(err)
	}

	tableStr, err := gomdtable.GenerateTable(table)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(tableStr)
}

```



result:

```bash
❯ go run main.go
Table:
| Имя    | Возраст |
|--------|---------|
| Сергей | 31      |
| Ян     | 9       |
| Яна    | 110     |

Table list:
| Имя               | Возраст |
|-------------------|---------|
| Сергей            | 31      |
| Ян                | 9       |
| Яна               | 110     |

| Имя               | Возраст |
|-------------------|---------|
| Сергей Игоревич   | 31      |
| Ян Васильев       | 9       |
| Яна Александровна | 110     |
```
