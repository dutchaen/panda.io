package api

import "encoding/json"

type User struct {
	ID              int64  `json:"id"`
	Username        string `json:"username"`
	PasswordHash    string `json:"password_hash"`
	PasswordSalt    string `json:"password_salt"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Bio             string `json:"bio"`
	ProfilePhotoB64 string `json:"profile_photo_b64"`
	PrimaryColor    int    `json:"primary_color"`
	SecondaryColor  int    `json:"secondary_color"`
	IsBitcoinBaller bool   `json:"is_bitcoin_baller"`
	LinksJSON       string `json:"links_json"`
	CreatedAt       string `json:"created_at"`
}

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PublicUser struct {
	ID              int64  `json:"id"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	Bio             string `json:"bio"`
	ProfilePhotoB64 string `json:"profile_photo_b64"`
	PrimaryColor    int    `json:"primary_color"`
	SecondaryColor  int    `json:"secondary_color"`
	IsBitcoinBaller bool   `json:"is_bitcoin_baller"`
	Links           []Link `json:"links"`
	CreatedAt       string `json:"created_at"`
}

type AdministerUser struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Bio             string `json:"bio"`
	ProfilePhotoB64 string `json:"profile_photo_b64"`
	PrimaryColor    int    `json:"primary_color"`
	SecondaryColor  int    `json:"secondary_color"`
	Links           []Link `json:"links"`
}

type LoginResponseOk struct {
	Message  string `json:"message"`
	LoggedIn bool   `json:"logged_in"`
}

type RegisterResponseOk struct {
	Message    string `json:"message"`
	Registered bool   `json:"registered"`
}

type SetResponseOk struct {
	Message string `json:"message"`
	Set     bool   `json:"set"`
}

func (u *User) ToPublicUser() PublicUser {

	links := []Link{}
	json.Unmarshal([]byte(u.LinksJSON), &links)

	public_user := PublicUser{
		u.ID,
		u.Username,
		u.Name,
		u.Bio,
		u.ProfilePhotoB64,
		u.PrimaryColor,
		u.SecondaryColor,
		u.IsBitcoinBaller,
		links,
		u.CreatedAt,
	}

	return public_user
}

func (u *User) ToAdministerUser() AdministerUser {
	links := []Link{}
	json.Unmarshal([]byte(u.LinksJSON), &links)

	admin_user := AdministerUser{
		u.Username,
		u.Name,
		u.Email,
		u.Bio,
		u.ProfilePhotoB64,
		u.PrimaryColor,
		u.SecondaryColor,
		links,
	}

	return admin_user
}

type LoginRequestJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequestJson struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
