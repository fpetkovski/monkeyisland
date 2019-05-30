package weapons

import "github.com/thedevsaddam/govalidator"

func ValidationRules() govalidator.MapData {
	return govalidator.MapData{
		"name":        []string{"required"},
		"power_level": []string{"required", "min:0", "max:100"},
	}
}
