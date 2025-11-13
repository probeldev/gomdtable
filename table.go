// Package gomdtable package is needed for generating Markdown tables.
package gomdtable

import (
	"errors"
	"unicode/utf8"
)

type tableList struct {
	tables        []Table
	maxCountChars []int
}

type Table struct {
	header []string
	rows   [][]string
}

var errorCountCellIncorrect = errors.New("count cell incorrect")

func (t *tableList) validateCountElements() error {
	count := -1
	for _, table := range t.tables {
		if count == -1 {
			count = len(table.header)
			continue
		}

		if count != len(table.header) {
			return errorCountCellIncorrect
		}

		for _, row := range table.rows {

			if count != len(row) {
				return errorCountCellIncorrect
			}
		}
	}
	return nil
}

func (t *Table) validateCountElements(row []string) error {
	if len(t.header) != 0 && len(t.header) != len(row) {
		return errorCountCellIncorrect
	}

	for _, r := range t.rows {
		if len(r) != len(row) {
			return errorCountCellIncorrect
		}
	}

	return nil
}

func (t *Table) SetHeader(header []string) error {
	t.header = []string{}

	err := t.validateCountElements(header)
	if err != nil {
		return err
	}

	t.header = header
	return nil
}

func (t *Table) AddRow(row []string) error {
	err := t.validateCountElements(row)
	if err != nil {
		return err
	}

	t.rows = append(t.rows, row)
	return nil
}

func NewTable() Table {
	t := Table{}

	return t
}

func GenerateTable(t Table) (string, error) {
	strList, err := GenerateTableList([]Table{t})
	if err != nil {
		return "", err
	}
	return strList[0], nil

}

func GenerateTableList(tables []Table) ([]string, error) {
	response := []string{}
	tl := tableList{}
	tl.tables = tables

	err := tl.validateCountElements()
	if err != nil {
		return nil, err
	}

	for i, table := range tl.tables {
		if i == 0 {
			for _, cell := range table.header {
				l := utf8.RuneCountInString(cell)
				tl.maxCountChars = append(tl.maxCountChars, l)
			}
		}
		for index, cell := range table.header {
			l := utf8.RuneCountInString(cell)
			tl.maxCountChars[index] = max(tl.maxCountChars[index], l)
		}
		for _, row := range table.rows {
			for index, cell := range row {
				l := utf8.RuneCountInString(cell)
				tl.maxCountChars[index] = max(tl.maxCountChars[index], l)
			}
		}

	}

	for _, table := range tl.tables {
		resp := ""
		delimiter := ""
		for i, h := range table.header {
			isLast := false
			if i == len(table.header)-1 {
				isLast = true
			}

			resp += generateCell(h, tl.maxCountChars[i], isLast)
			delimiter += generateDelimiter(tl.maxCountChars[i], isLast)

		}
		resp += "\n"
		resp += delimiter
		resp += "\n"
		for _, row := range table.rows {
			for i, r := range row {
				isLast := false
				if i == len(table.header)-1 {
					isLast = true
				}

				resp += generateCell(r, tl.maxCountChars[i], isLast)
			}
			resp += "\n"

		}
		response = append(response, resp)
	}

	return response, nil
}

func generateCell(cell string, countChar int, isLast bool) string {
	l := utf8.RuneCountInString(cell)
	cell = "| " + cell

	for i := l; i <= countChar; i++ {
		cell += " "
	}

	if isLast {
		cell += "|"
	}

	return cell
}

func generateDelimiter(countChar int, isLast bool) string {
	delimiter := "|"
	for i := 0; i < countChar+2; i++ {
		delimiter += "-"
	}

	if isLast {
		delimiter += "|"
	}

	return delimiter
}
