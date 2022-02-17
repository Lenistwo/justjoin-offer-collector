package csv

import (
	"github.com/lenistwo/model"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	FileName       = "output/out.csv"
	ValueDelimiter = ";"
	OfferURL       = "https://justjoin.it/offers/"
)

func Write(arr []model.Offer) {
	file, err := os.OpenFile(FileName, os.O_CREATE, 655)
	checkError(err)
	defer file.Close()
	_, err = file.WriteString("title;employment;salary from;salary to;skills;url\n")
	checkError(err)

	for _, o := range arr {
		employmentType := ""
		from := 0
		to := 0
		for _, e := range o.EmploymentTypes {
			employmentType += e.Type + " "
			from = int(math.Max(float64(from), float64(e.Salary.From)))
			to = int(math.Max(float64(to), float64(e.Salary.To)))
		}

		skills := ""
		for _, s := range o.Skills {
			skills += s.Name + " " + strconv.Itoa(s.Level) + "  "
		}

		_, err := file.WriteString(strings.Join([]string{o.Title, employmentType, strconv.Itoa(from), strconv.Itoa(to), skills, OfferURL + o.Id}, ValueDelimiter) + "\n")
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
