package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Berita struct {
	ID    int    `json:"id"`
	Date  string `json:"date"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
