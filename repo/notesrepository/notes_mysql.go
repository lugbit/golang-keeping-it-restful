// Package notesrepository stores functions that directly interacts with the database
package notesrepository

import (
	"database/sql"
	"log"

	"../../models"
)

// NotesRepository struct for initialization
type NotesRepository struct{}

// GetNotes interacts with the db directly to get all notes and returns
// a slice of notes
func (n NotesRepository) GetNotes(db *sql.DB, note models.Note, notes []models.Note) []models.Note {
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

	// Loop through each row
	for rows.Next() {
		// Scan each row to note struct
		err := rows.Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt, &note.Title, &note.Note, &note.Colour, &note.Archived)
		if err != nil {
			log.Fatalln(err)
		}

		// Append to slice of notes
		notes = append(notes, note)
	}

	// Check for errors during row scan
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	// Return slice of notes
	return notes
}

// GetNote serches the db for a note based on argument ID and returns a models.Note
func (n NotesRepository) GetNote(db *sql.DB, note models.Note, id int) models.Note {
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
		// No rows found
		if err == sql.ErrNoRows {
			// Return empty note
			return models.Note{}
		}
		// Something else went wrong
		log.Fatalln(err)
	}

	return note
}

// AddNote inserts a new note in the database and returns the last insert ID
func (n NotesRepository) AddNote(db *sql.DB, note models.Note) int {
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

	// Get last insert ID
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	// Return last insert ID
	return int(id)
}

// UpdateNote updates an existing note entry in the database
func (n NotesRepository) UpdateNote(db *sql.DB, note models.Note) int {
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

	// Return number of rows affected
	return int(rowsAffected)
}

// DeleteNote deletes an existing note entry in the database
func (n NotesRepository) DeleteNote(db *sql.DB, id int) int {
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

	// Return number of rows affected
	return int(rowsAffected)
}
