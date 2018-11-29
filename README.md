# Note taking app the RESTful way
## Introduction
A [REST](https://en.wikipedia.org/wiki/Representational_state_transfer)ful backend for a note taking app. Postman collection can be found [here](https://www.getpostman.com/collections/fa57fff58077138d4f68).

## API Overiew
| HTTP Request Method  | Resource               |Protected    | Description                                        |
| ---------------------|------------------------|-------------|----------------------------------------------------|
| POST                 | /v1/users/authenticate | No          | Authenticate and receive a JWT token.              |
| GET                  | /v1/notes              | Yes         | Retrieve all notes belonging to a user.            |
| GET                  | /v1/notes/:id          | Yes         | Retrieve a specific note by ID belonging to a user.| 
| POST                 | /v1/notes              | Yes         | Add new note belonging to a user.                  |
| PUT                  | /v1/notes              | Yes         | Update a user's existing note.                     |
| DELETE               | /v1/notes/:id          | Yes         | Delete a user's existing note.                     |

----

### Authenticate

Authenticate a user.

* **URL**

  /v1/users/authenticate

* **Method:**

  `POST`
  
*  **URL Params**
  
   None

* **Data Params**

  ```
  {
    "email": "Username",
    "password": "Password"
  }
  ```

* **Success Response:**

  * **Code:** 200 OK <br/>
    **Content:**
    ```
    {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZmlyc3ROYW1lI..."
    }
    ```
 
* **Error Responses:**

  * **Invalid username/password** <br/>
    **Code:** 400 Bad Request <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 3403,
            "text": "Incorrect username or password",
            "hint": "Verify that the email address and password are correct",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```
    
  * **Empty username and or password** <br/> 
    **Code:** 400 Bad Request <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 3400,
            "text": "Email cannot be empty",
            "hint": "Email is a required field",
            "info": "https://api.lugbit.com/docs/errors"
        },
        {
            "code": 3401,
            "text": "Password cannot be empty",
            "hint": "Password is a required field",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```
    
  * **Invalid email** <br/> 
    **Code:** 400 Bad Request <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 3499,
            "text": "Email is invalid",
            "hint": "Ensure the email is formatted in example@email.com",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```
    
### Get notes

Returns JSON data of every note.

* **URL**

  /v1/notes

* **Method:**

  `GET`
  
*  **URL Params**
  
   None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 OK <br/>
    **Content:**
    ```
    [
      {
        "id": 2,
        "created": "2018-11-29 20:10:29",
        "updated": "2018-11-29 20:10:29",
        "title": "Groceries",
        "note": "Milk, bread and ice-cream",
        "colour": "#ffffff",
        "archived": false,
        "userID": 1
      },
      {
        "id": 4,
        "created": "2018-11-29 20:11:01",
        "updated": "2018-11-29 20:11:01",
        "title": "New Note",
        "note": "This note is brought to you by Postman",
        "colour": "#ffffff",
        "archived": false,
        "userID": 1
      }
    ]
    ```
 
* **Error Response:**

  * **Empty notes collection** <br/> 
    **Code:** 404 Not Found <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 6900,
            "text": "No notes found",
            "hint": "Ensure you have at least one note",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```
 
 ### Get note

Returns JSON data of a particular note.

* **URL**

  /v1/notes/:id  

* **Method:**

  `GET`
  
*  **URL Params**
  
   id=[integer]

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 OK <br/>
    **Content:**
    ```
    {
      "id": 1,
      "created": "2018-11-29 20:52:14",
      "updated": "2018-11-29 20:52:14",
      "title": "First Note",
      "note": "This is my first note",
      "colour": "#ffffff",
      "archived": false,
      "userID": 1
    }
    ```
 
* **Error Response:**

  * **Note not found** <br/> 
    **Code:** 404 Not Found <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 6901,
            "text": "Note doesn't exist",
            "hint": "There is no note belonging to you with that ID",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```

### Add note

Adds a new note.

* **URL**

  /v1/notes 

* **Method:**

  `POST`
  
*  **URL Params**
  
   None

* **Data Params**

  ```
  {
    "title":"New Note",
    "note":"This note is brought to you by Postman"
  }
  ```

* **Success Response:**

  * **Code:** 201 Created <br/>
    **Content:**
    ```
    {
      "id": 4,
      "created": "2018-11-29 20:53:36",
      "updated": "2018-11-29 20:53:36",
      "title": "New Note",
      "note": "This note is brought to you by Postman",
      "colour": "#ffffff",
      "archived": false,
      "userID": 1
    }
    ```
 
* **Error Response:**

  * **Internal server error** <br/>
    **Code:** 400 Bad Request <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 6906,
            "text": "Unable to add note",
            "hint": "Internal server error",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```
  * **Empty title and or body** <br/>
    **Code:** 400 Bad Request <br/>
    **Content:**
      ```
      {
        "errors": [
          {
            "code": 6800,
            "text": "Title cannot be empty",
            "hint": "Note title is required",
            "info": "https://api.lugbit.com/docs/errors"
          },
          {
            "code": 6850,
            "text": "Note text cannot be empty",
            "hint": "Note body is required",
            "info": "https://api.lugbit.com/docs/errors"
          }
      ]
      }
      ```

### Update note (Adds note if note doesn't exist)

Updates an existing note.

* **URL**

  /v1/notes   

* **Method:**

  `PUT`
  
*  **URL Params**
  
   None

* **Data Params**

  ```
  {
	  "id": 1,
	  "title":"Note updated",
	  "note":"This note was updated successfully.",
	  "colour":"#0000ff",
	  "archived": false
  }
  ```

* **Success Response:**

  * **Code:** 201 Created <br/>
    **Content:**
    ```
    {
      "id": 1,
      "created": "2018-11-29 20:52:14",
      "updated": "2018-11-29 21:00:57",
      "title": "Note updated",
      "note": "This note was updated successfully.",
      "colour": "#0000ff",
      "archived": false,
      "userID": 1
    }
    ```
 
* **Error Response:**

  * ** Update failed** <br/>
    **Code:** 400 Bad Request <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 6901,
            "text": "Updating note failed",
            "hint": "Ensure the note ID you are updating exists.",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```
    
   * ** Empty required fields** <br/>
    **Code:** 400 Bad Request <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 6888,
            "text": "Note ID is required",
            "hint": "A note ID is required when updating a note",
            "info": "https://api.lugbit.com/docs/errors"
        },
        {
            "code": 6800,
            "text": "Title cannot be empty",
            "hint": "Note title is required",
            "info": "https://api.lugbit.com/docs/errors"
        },
        {
            "code": 6850,
            "text": "Note text cannot be empty",
            "hint": "Note body is required",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```
    
### Delete note

Deletes an existing note.

* **URL**

  /v1/notes/:id 

* **Method:**

  `DELETE`
  
*  **URL Params**
  
   id=[integer]

* **Data Params**

  None

* **Success Response:**

  * **Code:** 204 No Content<br/>
 
* **Error Response:**

  * **Code:** 400 Bad Request <br/>
    **Content:**
    ```
    {
    "errors": [
        {
            "code": 6901,
            "text": "Deleting note failed",
            "hint": "Ensure the ID of the note you are trying to delete exists.",
            "info": "https://api.lugbit.com/docs/errors"
        }
    ]
    }
    ```


