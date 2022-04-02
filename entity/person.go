package entity

//Person object for REST(CRUD)
type Phonebook struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phonenumber int64  `json:"phonenumber"`
}
