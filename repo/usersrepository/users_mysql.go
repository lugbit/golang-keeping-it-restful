package usersrepository

import (
	"database/sql"
	"errors"
	"log"

	usersmodels "../../models"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository struct{}

// UserEmail queries the database for a user based on the argument email.
//
// UserEmail will return the User and an error.
func (u UsersRepository) UserEmail(db *sql.DB, email string, user usersmodels.User) (usersmodels.User, error) {
	sqlStmt := `SELECT
					fldID,
    				fldFirstName,
    				fldLastName,
    				fldEmail,
    				fldPassword
				FROM tblUsers
				WHERE fldEmail = ?;`

	// Prepare statment
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	// Return row and scan columns to passed note
	err = stmt.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		// No row found
		if err == sql.ErrNoRows {
			// Return empty user and error
			return usersmodels.User{}, errors.New("Row not found")
		}
		// Something else went wrong
		log.Fatalln(err)
	}

	return user, nil
}

// VerifyPassword verifies the user provided password with the hashed password entry in the database.
func (u UsersRepository) VerifyPassword(db *sql.DB, providedPassword string, user usersmodels.User) error {
	// Compare hashed password in the DB with the user provided password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		// Password's don't match
		return errors.New("Password's do not match")
	}

	// No errors, passwords match.
	return nil
}
