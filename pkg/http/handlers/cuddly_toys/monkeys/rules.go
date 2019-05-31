package monkeys

import "github.com/thedevsaddam/govalidator"

func ValidationRules() govalidator.MapData {
	return govalidator.MapData{
		"name":         []string{"required"},
		"energy_level": []string{"required", "min:0", "max:100"},
	}
}
