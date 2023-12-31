package api

import "errors"

var (
	InvalidPayloadError       = errors.New("A invalid payload was detected.")
	AlreadyLoggedInError      = errors.New("You are already logged in. Please sign out.")
	UsernameOrPasswordInvalid = errors.New("The username or password you have entered is incorrect.")
	AuthenticationError       = errors.New("Ohhhhhhh.... Nobody 'ppose to be here. (nobody 'ppose to be heree-eee-eeee-eeeee).")
	UserNotFound              = errors.New("User was not found.")
	BadEmailProvidedError     = errors.New("This email address cannot be used.")
	BadNameProvidedError      = errors.New("This name cannot be used.")
	BadUsernameProvidedError  = errors.New("This username cannot be used.")
	BadPrimaryColorError      = errors.New("Sorry, the primary color you chose is not valid.")
	BadSecondaryColorError    = errors.New("Sorry, the secondary color you chose is not valid.")
	PasswordIsNotOkError      = errors.New("This password is not OK.")
	TooManyLinksProvidedError = errors.New("Too many links have been provided.")
	BadLinksProvidedError     = errors.New("Sorry, the links you have provided cannot be used.")
)
