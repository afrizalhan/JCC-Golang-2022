package phone

import "strconv"

type Phone struct{
	Name, Brand string
	Year int
	Colors []string
}


type PhoneOps interface{
	AddPhone() string
}


func (ph Phone) AddPhone() string {
	name := ph.Name
	brand := ph.Brand
	year := strconv.Itoa(ph.Year)
	colorStr := ""
	for i := 0; i < len(ph.Colors) - 1; i++{
		colorStr += ph.Colors[i] + ", "
	}

	colorStr += ph.Colors[len(ph.Colors) - 1]

	return "name : " + name + "\nbrand : " + brand + "\nyear : " + year + "\ncolors : " + colorStr
}