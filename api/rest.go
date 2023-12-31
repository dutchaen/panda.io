package api

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func WriteMeRest(ctx *fiber.Ctx) error {
	cookie_jar := ctx.Locals("cookie_jar").(*session.Store)
	database, err := sql.Open("mysql", DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	session, err := cookie_jar.Get(ctx)
	if err != nil {
		log.Println(err)
		return send_redirect(ctx, "/login")
	}

	if session.Fresh() {
		log.Println(err)
		return send_redirect(ctx, "/login")
	}

	id := session.Get("id").(int64)
	stmt, err := database.Prepare("SELECT id FROM users WHERE id = ?;")
	if err != nil {
		return send_internal_error(ctx)
	}
	row := stmt.QueryRow(id)

	var _id int64
	if err := row.Scan(&_id); err != nil {
		log.Println(err)
		return send_internal_error(ctx)
	}
	stmt.Close()

	admin_user := AdministerUser{}
	if err := json.Unmarshal(ctx.Body(), &admin_user); err != nil {
		log.Println(err)
		return send_internal_error(ctx)
	}

	username_available, err := is_username_available_to_id(database, admin_user.Username, id)
	if err != nil || !username_available {
		log.Println(err)
		return send_error(ctx, BadUsernameProvidedError)
	}

	email_available, err := is_email_available_to_id(database, admin_user.Email, id)
	if err != nil || !email_available {
		log.Println(err)
		return send_error(ctx, BadEmailProvidedError)
	}

	if !is_name_ok(admin_user.Name) {
		log.Println(err)
		return send_error(ctx, BadNameProvidedError)
	}

	// if colors is in rgb range
	if admin_user.PrimaryColor > 0xFFFFFF {
		log.Println(err)
		return send_error(ctx, BadPrimaryColorError)
	}

	if admin_user.SecondaryColor > 0xFFFFFF {
		log.Println(err)
		return send_error(ctx, BadSecondaryColorError)
	}

	if len(admin_user.Links) > 5 {
		return send_error(ctx, TooManyLinksProvidedError)
	}

	if !is_links_ok(&admin_user.Links) {
		return send_error(ctx, BadLinksProvidedError)
	}

	links_json, _ := json.Marshal(&admin_user.Links)

	stmt, err = database.Prepare("UPDATE users SET username = ?, name = ?, email = ?, bio = ?, profile_photo_b64 = ?, primary_color = ?, secondary_color = ?, links_json = ? WHERE id = ?")
	if err != nil {
		log.Println(err)
		return send_internal_error(ctx)
	}
	_, err = stmt.Exec(
		admin_user.Username,
		admin_user.Name,
		admin_user.Email,
		admin_user.Bio,
		admin_user.ProfilePhotoB64,
		admin_user.PrimaryColor,
		admin_user.SecondaryColor,
		string(links_json),
		id,
	)
	if err != nil {
		log.Println(err)
		return send_internal_error(ctx)
	}
	stmt.Close()

	return send_ok_json(ctx, SetResponseOk{
		Message: "Successfully saved changes.",
		Set:     true,
	})
}

func GetMeRest(ctx *fiber.Ctx) error {

	cookie_jar := ctx.Locals("cookie_jar").(*session.Store)
	database, err := sql.Open("mysql", DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	session, err := cookie_jar.Get(ctx)
	if err != nil {
		log.Println(err)
		return send_redirect(ctx, "/login")
	}

	if session.Fresh() {
		log.Println(err)
		return send_redirect(ctx, "/login")
	}

	id := session.Get("id").(int64)
	stmt, err := database.Prepare("SELECT * FROM users WHERE id = ?;")
	if err != nil {
		return send_internal_error(ctx)
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)

	user_model := User{}
	err = row.Scan(
		&user_model.ID,
		&user_model.Username,
		&user_model.PasswordHash,
		&user_model.PasswordSalt,
		&user_model.Email,
		&user_model.Name,
		&user_model.Bio,
		&user_model.ProfilePhotoB64,
		&user_model.PrimaryColor,
		&user_model.SecondaryColor,
		&user_model.IsBitcoinBaller,
		&user_model.LinksJSON,
		&user_model.CreatedAt,
	)

	if err != nil {
		log.Println(err)
		return send_internal_error(ctx)
	}

	admin_user := user_model.ToAdministerUser()
	return ctx.JSON(admin_user)
}

func UserLookupRest(ctx *fiber.Ctx) error {
	database, err := sql.Open("mysql", DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	username := ctx.Params("username")
	stmt, err := database.Prepare("SELECT * FROM users WHERE username = ?;")
	if err != nil {
		return send_internal_error(ctx)
	}

	defer stmt.Close()
	row := stmt.QueryRow(username)

	user_model := User{}

	err = row.Scan(
		&user_model.ID,
		&user_model.Username,
		&user_model.PasswordHash,
		&user_model.PasswordSalt,
		&user_model.Email,
		&user_model.Name,
		&user_model.Bio,
		&user_model.ProfilePhotoB64,
		&user_model.PrimaryColor,
		&user_model.SecondaryColor,
		&user_model.IsBitcoinBaller,
		&user_model.LinksJSON,
		&user_model.CreatedAt,
	)

	if err != nil {
		return send_error(ctx, UserNotFound)
	}

	public_user := user_model.ToPublicUser()
	return ctx.JSON(public_user)
}

func RegisterRest(ctx *fiber.Ctx) error {

	cookie_jar := ctx.Locals("cookie_jar").(*session.Store)
	database, err := sql.Open("mysql", DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	session, err := cookie_jar.Get(ctx)
	if err != nil {
		log.Println(err)
		return send_error(ctx, AlreadyLoggedInError)
	}

	if !session.Fresh() {
		log.Println(err)
		return send_error(ctx, AlreadyLoggedInError)
	}

	request := RegisterRequestJson{}
	if err := json.Unmarshal(ctx.Body(), &request); err != nil {
		log.Println(err)
		session.Destroy()
		return send_error(ctx, InvalidPayloadError)
	}

	email_ok, err := is_email_available(database, request.Email)
	if err != nil || !email_ok {
		log.Println(err)
		session.Destroy()
		return send_error(ctx, BadEmailProvidedError)
	}

	username_ok, err := is_username_available(database, request.Username)
	if err != nil || !username_ok {
		log.Println(err)
		session.Destroy()
		return send_error(ctx, BadUsernameProvidedError)
	}

	if !is_password_ok(request.Password) {
		log.Println("err=" + request.Password)
		session.Destroy()
		return send_error(ctx, PasswordIsNotOkError)
	}

	registered_id, err := register_user_to_db(database, &request)
	if err != nil {
		log.Println(err)
		session.Destroy()
		return send_internal_error(ctx)
	}

	session.Set("id", registered_id)
	session.Save()

	return send_ok_json(ctx, RegisterResponseOk{
		Message:    "Aloha!",
		Registered: true,
	})
}

func LoginRest(ctx *fiber.Ctx) error {
	cookie_jar := ctx.Locals("cookie_jar").(*session.Store)
	database, err := sql.Open("mysql", DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	session, err := cookie_jar.Get(ctx)
	if err != nil {
		return send_error(ctx, AlreadyLoggedInError)
	}

	if !session.Fresh() {
		return send_error(ctx, AlreadyLoggedInError)
	}

	request := LoginRequestJson{}
	if err := json.Unmarshal(ctx.Body(), &request); err != nil {
		session.Destroy()
		return send_error(ctx, InvalidPayloadError)
	}

	stmt, err := database.Prepare("SELECT id, password_hash, password_salt FROM users WHERE username = ? OR email = ?;")
	if err != nil {
		session.Destroy()
		return send_internal_error(ctx)
	}

	row := stmt.QueryRow(request.Username, request.Username)
	stmt.Close()

	var id int64
	var password_hash string
	var password_salt string

	err = row.Scan(&id, &password_hash, &password_salt)
	if err != nil {
		session.Destroy()
		return send_internal_error(ctx)
	}

	// if password is invalid
	if !cmp_password(request.Password, password_salt, password_hash) {
		session.Destroy()
		return send_error(ctx, UsernameOrPasswordInvalid)
	}

	session.Set("id", id)
	session.Save()

	return send_ok_json(ctx, LoginResponseOk{
		Message:  "Welcome!",
		LoggedIn: true,
	})
}
