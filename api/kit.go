package api

import (
	"crypto/md5"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const DATABASE_AUTHENTICATION = "root:Login1234@/panda"

var EMAIL_REGEX = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")

var BANNED_DOMAINS_OR_KEYWORDS = []string{
	"onlyfans.com",
	"doxbin",
	"grabify.io",
	"spylink.net",
	"porn",
	"xxx",
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func send_ok_json(ctx *fiber.Ctx, msg interface{}) error {
	if _, err := json.Marshal(msg); err != nil {
		return send_internal_error(ctx)
	}

	fmt.Println(msg)

	return ctx.Status(200).JSON(msg)
}

func send_redirect(ctx *fiber.Ctx, to string) error {
	return ctx.Redirect(to)
}

func send_error(ctx *fiber.Ctx, msg error) error {
	err := ErrorMessage{
		Error: msg.Error(),
	}

	byte_array, _ := json.Marshal(err)

	return ctx.Status(400).SendString(string(byte_array))
}

func send_internal_error(ctx *fiber.Ctx) error {

	err := ErrorMessage{
		Error: "An internal server error has occured.",
	}

	byte_array, _ := json.Marshal(err)

	return ctx.Status(500).SendString(string(byte_array))
}

func generate_salt() string {
	writer := md5.New()

	unix := time.Now().UnixMilli()
	seed := rand.Int63()

	value := fmt.Sprint(unix & seed)

	writer.Write([]byte(value))

	return hex.EncodeToString(writer.Sum(nil))
}

func hash_password(text string, salt string) string {

	writer := sha512.New()
	writer.Write([]byte(salt + text))

	return hex.EncodeToString(writer.Sum(nil))
}

func is_username_available(db *sql.DB, username string) (bool, error) {
	if !is_username_ok(username) {
		return false, nil
	}

	stmt, err := db.Prepare("SELECT id FROM users WHERE username = ?;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var id int
	row := stmt.QueryRow(username)
	if err := row.Scan(&id); err != nil {
		return true, nil
	}

	return false, nil
}

func is_username_available_to_id(db *sql.DB, username string, m_id int64) (bool, error) {
	if !is_username_ok(username) {
		return false, nil
	}

	stmt, err := db.Prepare("SELECT id FROM users WHERE username = ?;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var id int64
	row := stmt.QueryRow(username)
	if err := row.Scan(&id); err != nil {
		return true, nil
	}

	if m_id == id {
		return true, nil
	}

	return false, nil
}

func is_email_available_to_id(db *sql.DB, email string, m_id int64) (bool, error) {
	if !is_email_ok(email) {
		return false, nil
	}

	stmt, err := db.Prepare("SELECT id FROM users WHERE email = ?;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var id int64
	row := stmt.QueryRow(email)
	if err := row.Scan(&id); err != nil {
		return true, nil
	}

	if m_id == id {
		return true, nil
	}

	return false, nil
}

func is_email_available(db *sql.DB, email string) (bool, error) {
	if !is_email_ok(email) {
		return false, nil
	}

	stmt, err := db.Prepare("SELECT id FROM users WHERE email = ?;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var id int
	row := stmt.QueryRow(email)
	if err := row.Scan(&id); err != nil {
		return true, nil
	}

	return false, nil
}

func is_email_ok(email string) bool {
	return EMAIL_REGEX.MatchString(email)
}

func is_username_ok(username string) bool {
	if len(username) > 18 {
		return false
	}

	return true
}

func is_name_ok(name string) bool {
	if len(name) == 0 || len(name) > 20 {
		return false
	}

	return true
}

func extract_domain_from_url(urlLikeString string) string {

	urlLikeString = strings.TrimSpace(urlLikeString)

	if regexp.MustCompile(`^https?`).MatchString(urlLikeString) {
		read, _ := url.Parse(urlLikeString)
		urlLikeString = read.Host
	}

	if regexp.MustCompile(`^www\.`).MatchString(urlLikeString) {
		urlLikeString = regexp.MustCompile(`^www\.`).ReplaceAllString(urlLikeString, "")
	}

	return regexp.MustCompile(`([a-z0-9\-]+\.)+[a-z0-9\-]+`).FindString(urlLikeString)
}

func is_links_ok(links *[]Link) bool {

	if len(*links) > 5 {
		return false
	}

	for _, link := range *links {

		if _, err := url.Parse(link.URL); err != nil {
			log.Println(err)
			return false
		}

		domain := extract_domain_from_url(link.URL)
		for _, banned_domain := range BANNED_DOMAINS_OR_KEYWORDS {
			if strings.Contains(domain, banned_domain) {
				return false
			}
		}
	}
	return true
}

func is_password_ok(password string) bool {
	if len(password) <= 8 {
		return false
	}

	return true
}

func register_user_to_db(db *sql.DB, request *RegisterRequestJson) (int64, error) {
	now := time.Now()

	id := now.UnixNano()
	date_string := now.String()
	salt := generate_salt()
	hashed_password := hash_password(request.Password, salt)

	stmt, err := db.Prepare("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		id,
		request.Username,
		hashed_password,
		salt,
		request.Email,
		request.Name,
		"No information given.",
		"data:image/png;base64,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",
		0xEEEEEE,
		0x888888,
		false,
		"[]",
		date_string,
	)

	if err != nil {
		log.Println(err)
		return -1, err
	}

	return id, nil

}

func cmp_password(password string, salt string, password_hash string) bool {
	new_hash := hash_password(password, salt)
	return new_hash == password_hash
}
