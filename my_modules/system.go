package my_modules

import (
	"fmt"
)

type User struct {
	Username string
	Password string
	Email    string
}

type Item struct {
	Name   string
	Price  float64
	Rating float64
}

func (u *User) Register(username, password, email string) {
	u.Username = username
	u.Password = password
	u.Email = email
	fmt.Println("User Registered:", u.Username)
}

func (u *User) Login(username, password string) bool {
	if u.Username == username && u.Password == password {
		fmt.Println("User Logged in:", u.Username)
		return true
	}
	fmt.Println("Login Failed")
	return false
}

func (i *Item) GiveRating(rating float64) {
	i.Rating = rating
	fmt.Println("Item Rated:", i.Name)
}

type ItemStore struct {
	Items []Item
}

func (is *ItemStore) Search(name string) []Item {
	var result []Item
	for _, item := range is.Items {
		if item.Name == name {
			result = append(result, item)
		}
	}
	return result
}

func (is *ItemStore) Filter(price, rating float64) []Item {
	var result []Item
	for _, item := range is.Items {
		if item.Price <= price && item.Rating >= rating {
			result = append(result, item)
		}
	}
	return result
}
