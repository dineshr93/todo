package model

// code from joefazee modified by Dinesh Ravi to support various changes https://github.com/joefazee/go-toto-app/pull/1
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

const notime = string("########")

type item struct {
	Task        string
	Done        bool
	CreatedAt   string
	CompletedAt string
}

type Todos []item

func (t *Todos) Add(task string) {

	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now().Format(time.RFC822),
		CompletedAt: notime,
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index. Please mention only indexes with in range")
	}

	// ignore the CompletedAt time if already present
	if !ls[index-1].Done {
		ls[index-1].CompletedAt = time.Now().Format(time.RFC822)
		ls[index-1].Done = true
	}

	return nil
}

func (t *Todos) CompleteSA(indexes []string) error {
	ls := *t
	if len(indexes) > len(ls) {
		return errors.New("invalid index. Please mention only indexes with in range")
	}
	for _, s := range indexes {
		index, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid args. Only numbers are allowed")
			return nil
		}
		t.Complete(index)
	}

	return nil
}

func (t *Todos) Pending(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index. Please mention only indexes with in range")
	}

	ls[index-1].CompletedAt = notime
	ls[index-1].Done = false

	return nil
}

func (t *Todos) PendingSA(indexes []string) error {
	ls := *t
	if len(indexes) > len(ls) {
		return errors.New("invalid index. Please mention only indexes with in range")
	}
	for _, s := range indexes {
		index, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid args. Only numbers are allowed")
			return nil
		}
		t.Pending(index)
	}

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index. Please mention only indexes with in range")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}
func (t *Todos) DeleteSA(indexes []string) error {
	ls := *t
	if len(indexes) > len(ls) {
		return errors.New("invalid index")
	}
	// n is crucial for proper deleteion of slice in golang. refer https://stackoverflow.com/a/19954213/2018343
	n := 0
	for _, s := range indexes {

		index, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid args. Only numbers are allowed")
			return nil
		}
		index = index - n
		t.Delete(index)
		n++

	}

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
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

	return ioutil.WriteFile(filename, data, 0644)
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
		task := blue(item.Task)
		done := blue("no")
		completeat := blue(item.CompletedAt)
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")

			completeat = item.CompletedAt
			// .Format(time.RFC822)
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt},
			{Text: completeat},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	if t.CountPending() > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{

			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
		}}
	}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return total
}
