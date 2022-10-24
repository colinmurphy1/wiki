package state

import (
	"testing"
)

// Test the ReadDatabase function with a file that exists and one that does not
func TestReadDatabase(t *testing.T) {
	users := UserList{
		Users: make(map[string]User),
	}

	dbfiles := make(map[string]bool)

	dbfiles["./auth_test.db"] = true   // auth_test.db exists
	dbfiles["./auth_test.dbx"] = false // auth_test.dbx does not exist

	for file, exists := range dbfiles {
		o := users.ReadDatabase(file)
		// If there is not error and there should be, fail the test
		if o != nil && exists {
			t.Error("Could not read database, received:", o)
			continue
		}
		t.Logf("%s passed\n", file)
	}
}

// Tests GetUser function
func TestGetUser(t *testing.T) {
	users := UserList{Users: make(map[string]User)}

	// Open test database
	if err := users.ReadDatabase("./auth_test.db"); err != nil {
		t.Fatalf("Could not open test database, received: %s", err)
	}

	// Test a user that exists and does not exist
	testusers := make(map[string]bool)
	testusers["admin"] = true
	testusers["idontexist"] = false

	for user, exists := range testusers {
		_, err := users.GetUser(user)
		if err != nil && exists {
			t.Errorf("Test user \"%s\" FAILED, received: %s\n", user, err)
			continue
		}
		t.Logf("%s passed\n", user)
	}

}

type testuser struct {
	usage    string
	username string
	password string
	success  bool // Expected result
}

// Tests authenticate function
func TestAuthenticate(t *testing.T) {
	users := UserList{Users: make(map[string]User)}

	// Open test database
	if err := users.ReadDatabase("./auth_test.db"); err != nil {
		t.Fatalf("Could not open test database, received: %s", err)
	}

	// Test every possible result
	testusers := []testuser{
		{usage: "Normal signin", username: "admin", password: "admin", success: true},
		{usage: "Incorrect password", username: "admin", password: "root", success: false},
		{usage: "Nonexistent user", username: "idontexist", password: "idontexist", success: false},
		{usage: "Disabled user", username: "disabledUser", password: "hello", success: false},
	}

	for _, u := range testusers {
		_, err := users.Authenticate(u.username, u.password)
		if err != nil && u.success {
			t.Errorf("Test \"%s\" FAILED: %s\n", u.usage, err)
			continue
		}
		t.Logf("Test \"%s\" passed\n", u.usage)
	}

}
