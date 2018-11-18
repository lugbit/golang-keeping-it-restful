DROP DATABASE IF EXISTS notesDB;
CREATE DATABASE notesDB;
USE notesDB;

CREATE TABLE tblNotes (
	fldID BIGINT AUTO_INCREMENT PRIMARY KEY,
    fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
    fldTitle VARCHAR(50),
    fldNote TEXT NOT NULL,
    fldColour CHAR(7) DEFAULT "#ffffff",
    fldArchived BOOL DEFAULT FALSE
);