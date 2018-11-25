CALL insertNewUser("Marck", "Munoz", "Marck527@gmail.com", "$2a$10$FbqifHl4ISCSt9NqkxWdueSitWGl.azwn2FrWO46tTcQtfCBf8hqS"); -- PW: Password1
CALL insertNewUser("Brayden", "Gravestock", "braydengravestock@gmail.com", "$2a$10$H1PYw2zR1.sl/OFUNqJE/.LVmXdWplqCZiTuO7Mv.HAZ5PquqsgS."); -- PW: hunter2

CALL getAllUsers();

CALL insertNewNote("First Note", "This is my first note", 1);
CALL insertNewNote("Groceries", "Milk, bread and ice-cream", 1);
CALL insertNewNote("Reminder", "Drink water", 2);

CALL getAllNotes();

CALL getAllUserNotes();
