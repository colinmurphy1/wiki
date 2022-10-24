package state

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserList struct {
	Users map[string]User
}

type User struct {
	Enabled  bool   // Account enabled
	Username string // Username
	password string // Password hash
	FullName string // Full Name
	Email    string // Email address
	Id       int    // User ID
}

// Error definitions
var (
	ErrInvalidUser       = errors.New("user does not exist")
	ErrUserExists        = errors.New("user exists")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrUserDisabled      = errors.New("user account is disabled")
)

// ReadAuthDatabase - reads the database file
func (ul *UserList) ReadDatabase(dbFile string) error {
	// open the database file, and throw an error if it cannot be opened
	data, err := os.Open(dbFile)
	if err != nil {
		return err
	}

	defer data.Close()

	// iterate over each line of the file
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()

		userEnabled := true

		// ignore empty lines and lines that are commented out
		if len(line) == 0 || string(line[0]) == "#" {
			continue
		}

		// if the first character in a line is an exclamation point, the user is disabled
		if string(line[0]) == "!" {
			userEnabled = false

			// remove the first character from the user so the username of the user doesn't become !username
			line = line[1:]
		}

		// split the row to get the values at ':' and place in struct
		userInfo := strings.Split(line, ":")

		userId, _ := strconv.Atoi(userInfo[1])

		ul.Users[userInfo[0]] = User{
			Enabled:  userEnabled,
			Username: userInfo[0],
			password: userInfo[2],
			FullName: userInfo[3],
			Email:    userInfo[4],
			Id:       userId,
		}
	}

	return nil
}

// Get information about the user, returns nil if user does not exist
func (ul *UserList) GetUser(username string) (*User, error) {
	// Check if the username key exists in the ul.Users map
	u, exists := ul.Users[username]

	if !exists {
		return nil, ErrInvalidUser
	}

	return &u, nil
}

// Authenticate using username and password
func (ul *UserList) Authenticate(username string, password string) (*User, error) {
	// Get user data
	user, err := ul.GetUser(username)

	// Check if user exists
	if err != nil {
		return nil, ErrInvalidUser
	}

	// Check if user is enabled
	if !user.Enabled {
		return nil, ErrUserDisabled
	}

	// verify password against hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password)); err != nil {
		// password does not match
		return nil, ErrIncorrectPassword
	}

	return user, nil
}

// Create a new account
func (ul *UserList) Register(username string, password string, fullName string, email string) error {
	// Check if the user exists

	_, err := ul.GetUser(username)

	if err == nil {
		return ErrUserExists
	}

	// Generate hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Add the new user to the users struct
	ul.Users[username] = User{
		Enabled:  true,
		Username: username,
		password: string(hashedPassword),
		FullName: fullName,
		Email:    email,
		Id:       len(ul.Users),
	}

	// Write changes to disk
	ul.Commit()

	return nil
}

// Write any changes to disk
func (ul *UserList) Commit() error {
	// Re-create the file using information from the struct
	var lines string
	for _, usr := range ul.Users {
		var enabled string

		// Disabled user accounts have a ! in the first line
		if !usr.Enabled {
			enabled = "!"
		}

		lines += fmt.Sprintf("%s%s:%d:%s:%s:%s\n", enabled, usr.Username, usr.Id, usr.password, usr.FullName, usr.Email)
	}

	// Open db file for read/write
	db, err := os.OpenFile(Conf.Files.UsersDb, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer db.Close()

	// Write the new user file to disk
	if _, err := db.WriteString(lines); err != nil {
		return err
	}

	return nil
}
