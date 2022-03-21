package user

import "errors"

type user struct {
	id   int
	Name string
	Age  int
	Sex  int
}

var Users []user
var idNow = 0

func NewUser(name string, age int, sex int) error {
	if name == "" {
		return errors.New("Name is empty")
	}
	if age < 0 {
		return errors.New("Age bellow zero")
	}
	Users = append(Users, user{
		id:   idNow,
		Name: name,
		Age:  age,
		Sex:  sex,
	})
	idNow++
	return nil
}

func GetUser() []user {
	return Users
}

func GetUserById(id int) (user, error) {
	if id < 0 || id >= idNow {
		return user{}, errors.New("bad id")
	}
	return Users[id], nil
}

func DeleteUser(id int) {
	Users = append(Users[:id], Users[id+1:]...)

}
