package my_modules

import (
	"fmt"
	"sort"
)

type User struct {
	Username string
	Password string
	Email    string
}

type Database struct {
	Users []User
}

type Item struct {
	Name   string
	Price  float64
	Rating float64
}

type ItemStore struct {
	Items []Item
}

func (db *Database) Register(username, password, email string) {
	u := User{username, password, email}
	fmt.Println("User Registered: ", u)
	db.Users = append(db.Users, u)
}

func (db *Database) SeeUsersList() {
	for _, user := range db.Users {
		fmt.Println(user)
	}
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

func (is *ItemStore) AddingItem(name string, price float64) {
	var i = Item{name, price, 0}
	is.Items = append(is.Items, i)
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

func (is *ItemStore) FilterByPrice(price float64) []Item {
	result := make([]Item, len(is.Items))
	copy(result, is.Items)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Price < result[j].Price
	})
	return result
}

func (is *ItemStore) FilterByRating(price float64) []Item {
	result := make([]Item, len(is.Items))
	copy(result, is.Items)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Rating < result[j].Rating
	})
	return result
}
