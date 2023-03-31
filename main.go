package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {

	// mengubah angka bulan menjadi nama bulan dalam bahasa inggris
	var formattedMonth string
	switch month {
	case 1:
		formattedMonth = "January"
	case 2:
		formattedMonth = "February"
	case 3:
		formattedMonth = "March"
	case 4:
		formattedMonth = "April"
	case 5:
		formattedMonth = "May"
	case 6:
		formattedMonth = "June"
	case 7:
		formattedMonth = "July"
	case 8:
		formattedMonth = "August"
	case 9:
		formattedMonth = "September"
	case 10:
		formattedMonth = "October"
	case 11:
		formattedMonth = "November"
	case 12:
		formattedMonth = "December"
	default:
		formattedMonth = "Error"
	}

	// mengubah angka yang kurang dari 10
	var formattedDay string
	if day < 10 {
		formattedDay = "0" + fmt.Sprint(day)
	} else {
		formattedDay = fmt.Sprint(day)
	}

	result := formattedDay + "-" + formattedMonth + "-" + fmt.Sprint(year)
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
