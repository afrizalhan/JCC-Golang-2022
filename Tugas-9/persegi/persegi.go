package persegi

import "strconv"

func LuasPersegi(num int, action bool) interface{} {
	var any interface{}
	if num > 0 && action == true{
		luas := num * num
		any = "Luas Persegi dengan sisi " + strconv.Itoa(num) + "cm adalah " + strconv.Itoa(luas) +"cm"
	} else if num > 0 && action == false{
		any = num
	} else if num == 0 && action == true {
		any = "Maaf anda belum menginput sisi dari persegi"
	} else {
		any = nil
	}

	return any

}