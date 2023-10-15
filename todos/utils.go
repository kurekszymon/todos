package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

func (t *Todos) CheckConstraints(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("Invalid index")
	}

	return nil
}

func (t *Todos) Load(filename string) error {

	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0666)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return nil
}

func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		cells = append(cells, []*simpletable.Cell{
			{Text: strconv.Itoa(idx)},
			{Text: item.Task},
			{Text: strconv.FormatBool(item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CreatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignRight, Span: 5, Text: fmt.Sprintf("You have %d pending todos", t.Total())},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
