/* 
	STORED PROCEDURES FOR TESTING PURPOSES
*/


-- Return all users
DELIMITER //
CREATE PROCEDURE getAllUsers()
	BEGIN
		SELECT *  FROM tblUsers;
	END //
DELIMITER ;

-- Return all notes
DELIMITER //
CREATE PROCEDURE getAllNotes()
	BEGIN
		SELECT *  FROM tblNotes;
	END //
DELIMITER ;

-- Returns all notes belonging to a user
DELIMITER //
CREATE PROCEDURE getAllUserNotes()
	BEGIN
		SELECT
			tblUsers.fldID,
			CONCAT(tblUsers.fldFirstName, " ",tblUsers.fldLastName) AS fldName, 
			tblUsers.fldEmail,
			tblUsers.fldPassword,
			tblUsers.fldCreatedAt,
			tblUsers.fldUpdatedAt,
			tblNotes.fldID AS fldNoteID,
			tblNotes.fldTitle,
			tblNotes.fldNote,
			tblNotes.fldColour,
			tblNotes.fldArchived,
			tblNotes.fldCreatedAt AS fldNoteCreated,
			tblNotes.fldUpdatedAt AS fldNoteUpdated
		FROM tblUsers
		INNER JOIN tblNotes
			ON tblNotes.fldFKUserID = tblUsers.fldID;
	END //
DELIMITER ;

-- Insert new user
DELIMITER //
CREATE PROCEDURE insertNewUser(IN firstName VARCHAR(50), IN lastName VARCHAR(50), IN email VARCHAR(50), IN userPassword VARCHAR(255))
	BEGIN
		INSERT INTO tblUsers (fldFirstName, fldLastName, fldEmail, fldPassword) VALUES 
							 (firstName, lastName, email, userPassword);
	END //
DELIMITER ;

-- Insert new note
DELIMITER //
CREATE PROCEDURE insertNewNote(IN title VARCHAR(50), IN note TEXT, in userID BIGINT)
	BEGIN
		INSERT INTO tblNotes (fldTitle, fldNote, fldFKUserID) VALUES 
							 (title, note, userID);
	END //
DELIMITER ;


 
