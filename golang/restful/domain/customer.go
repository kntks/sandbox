package domain

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Customer struct {
	ID   string `json:"id" copier:"-"`
	Name Name   `json:"name"`
}
