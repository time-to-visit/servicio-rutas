package utils

func GetWhere(data map[string]interface{}) (string, []interface{}) {
	var i = 0
	var command = ""
	var request []interface{}
	for index, value := range data {
		command += index + " = ? "
		if i == 0 {
			command += " and "
		}
		request = append(request, value)
	}
	return command, request
}
