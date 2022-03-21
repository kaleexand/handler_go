package random

func GetNames() string {
	var names [5]string
	names[0] = "Anton"
	names[1] = "Maks"
	names[2] = "Yuliya"
	names[3] = "Alex"
	names[4] = "John"
	res := names[GetNumber(0, 4)]
	return res
}
