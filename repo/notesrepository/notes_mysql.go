// Package notesrepository stores functions that directly interacts with the database
package notesrepository

import (
	"database/sql"
	"errors"
	"log"

	"../../models"
)

// NotesRepository struct for initialization
type NotesRepository struct{}

// GetNotes queries the database for all notes and returns a slice of notes and an error.
func (n NotesRepository) GetNotes(db *sql.DB, note models.Note, notes []models.Note) ([]models.Note, error) {
	sqlStmt := `SELECT
					fldID,
					fldCreatedAt,
    				fldUpdatedAt, 
    				fldTitle,
    				fldNote, 
    				fldColour, 
    				fldArchived
				FROM tblNotes;`

	// Prepare statment
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	// rows will hold all found rows
	rows, err := stmt.Query()
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	// Row counter to find out if the query returned a row
	rowCount := 0
	// Loop through each row
	for rows.Next() {
		// Increment row
		rowCount++
		// Scan each row to note struct
		err := rows.Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt, &note.Title, &note.Note, &note.Colour, &note.Archived)
		if err != nil {
			log.Fatalln(err)
		}

		// Append to slice of notes
		notes = append(notes, note)
	}

	// No rows found
	if rowCount == 0 {
		// Return error
		return []models.Note{}, errors.New("No rows found")
	}

	// Check for errors during row scan
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	// Return slice of notes
	return notes, nil
}

// GetNote queries the database for a particular note and returns a note and an error.
func (n NotesRepository) GetNote(db *sql.DB, note models.Note, id int) (models.Note, error) {
	sqlStmt := `SELECT
					fldID,
					fldCreatedAt,
    				fldUpdatedAt, 
    				fldTitle,
    				fldNote, 
    				fldColour, 
    				fldArchived
				FROM tblNotes
				WHERE fldID = ?;`

	// Prepare statment
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	// Return row and scan columns to passed note
	err = stmt.QueryRow(id).Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt, &note.Title, &note.Note, &note.Colour, &note.Archived)
	if err != nil {
		// No row found
		if err == sql.ErrNoRows {
			// Return empty note and error
			return models.Note{}, errors.New("Row not found")
		}
		// Something else went wrong
		log.Fatalln(err)
	}

	return note, nil
}

// AddNote inserts a new note in the database and returns the last insert ID
// and an error.
func (n NotesRepository) AddNote(db *sql.DB, note models.Note) (int, error) {
	sqlStmt := `INSERT INTO tblNotes(fldTitle, fldNote) 
				VALUES 
					(?, ?)`

	// Prepare statment
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	// Exec does not return any rows
	res, err := stmt.Exec(note.Title, note.Note)
	if err != nil {
		log.Fatalln(err)
	}

	// Check if any rows was affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	// Rows affected is zero, insert failed
	if rowsAffected == 0 {
		return -1, errors.New("No rows affected")
	}

	// Get last insert ID
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	// Return last insert ID
	return int(id), nil
}

// UpdateNote updates an existing note in the database and returns the number of
// rows affected and an error.
func (n NotesRepository) UpdateNote(db *sql.DB, note models.Note) (int, error) {
	sqlStmt := `UPDATE tblNotes
					SET fldTitle = ?, fldNote = ?, fldColour = ?, fldArchived = ?
				WHERE fldID = ?;`

	// Prepare statment
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	// Exec does not return any rows
	res, err := stmt.Exec(note.Title, note.Note, note.Colour, note.Archived, note.ID)
	if err != nil {
		log.Fatalln(err)
	}

	// Get the number of rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	// If number of rows affected is zero, nothing was updated
	if rowsAffected == 0 {
		return -1, errors.New("No rows affected")
	}

	// Return number of rows affected
	return int(rowsAffected), nil
}

// DeleteNote deletes an existing note in the database and returns the number of
// rows affected and an error.
func (n NotesRepository) DeleteNote(db *sql.DB, id int) (int, error) {
	sqlStmt := `DELETE FROM tblNotes
				WHERE fldID = ?`

	// Prepare statment
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	// Exec does not return any rows
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatalln(err)
	}

	// Get the number of rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	// If number of rows affected is zero, nothing was deleted
	if rowsAffected == 0 {
		return -1, errors.New("No rows affected")
	}

	// Return number of rows affected
	return int(rowsAffected), nil
}
