package my_modules

import (
	"fmt"
	"sort"
)

type User struct {
	Username string `json:"username"`
	Password string	`json:"password"`
	Email    string	`json:"email"`
	Items    ItemStore `json:"items list"`
}

type Database struct {
	Users []User `json:"users"`
}

type Item struct {
	Name   string `json:"name"`
	Price  float64 `json:"price"`
	Rating float64 `json:"rating"`
}

type ItemStore struct {
	Items []Item `json:"items list"`
}

func (db *Database) Register(username, password, email string) User{
	u := User{username, password, email, ItemStore{}}
	fmt.Println("User Registered: ", u)
	db.Users = append(db.Users, u)
	return u
}

func (db *Database) SeeUsersList() {
	for _, user := range db.Users {
		fmt.Println(user)
	}
}

func (db *Database) Login(username, password string) User {
	for _, u := range db.Users {
		if u.Username == username && u.Password == password {
			fmt.Println("User Logged in:", u.Username)
			return u
		}
	}
	fmt.Println("Login Failed")
	return User{}
}

func (u *User) GetItems() {
	for _, item := range u.Items.Items {
		fmt.Println(item)
	}
}

func (u User) String() string {
	return fmt.Sprintf("username: %s, email: %s", u.Username, u.Email)
}

func (i *Item) GiveRating(rating float64) {
	i.Rating = rating
	fmt.Println(i)
}

func (i Item) String() string {
	return fmt.Sprintf("name: %s, price: %f, rating: %f", i.Name, i.Price, i.Rating)
}

func (is *ItemStore) AddingItem(name string, price float64) {
	var i = Item{name, price, 0}
	is.Items = append(is.Items, i)
}


func (u *User) Search(name string) Item {
	var result Item
	for _, item := range u.Items.Items {
		if item.Name == name {
			return result
		}
	}
	return result
}

func (is *ItemStore) FilterByPrice() []Item {
	result := make([]Item, len(is.Items))
	copy(result, is.Items)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Price < result[j].Price
	})
	return result
}

func (is *ItemStore) FilterByRating() []Item {
	result := make([]Item, len(is.Items))
	copy(result, is.Items)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Rating < result[j].Rating
	})
	return result
}
