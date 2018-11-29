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

  * **Invalid username/password**
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
    
  * **Empty username or password field** 
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
    
  * **Invalid email** 
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

  /notes

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
          "id": 1,
          "created": "2018-11-18 17:55:01",
          "updated": "2018-11-18 17:55:01",
          "title": "First Note",
          "note": "My first note.",
          "colour": "#ffffff",
          "archived": false
      },
      {
          "id": 2,
          "created": "2018-11-18 17:55:13",
          "updated": "2018-11-18 17:55:13",
          "title": "Second Note",
          "note": "My second note.",
          "colour": "#ffffff",
          "archived": false
      }
    ]
    ```
 
* **Error Response:**

  * **Code:** 404 Not Found <br/>
    **Content:** `{"message:":"empty collection"}`
 
 ### Get note

Returns JSON data of a particular note.

* **URL**

  /notes/:id

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
      "created": "2018-11-18 17:55:01",
      "updated": "2018-11-18 17:55:01",
      "title": "First Note",
      "note": "My first note.",
      "colour": "#ffffff",
      "archived": false
      }
    ```
 
* **Error Response:**

  * **Code:** 404 Not Found <br/>
    **Content:** `{"message:":"resource not found"}`

### Add note

Adds a new note.

* **URL**

  /notes

* **Method:**

  `POST`
  
*  **URL Params**
  
   None

* **Data Params**

  ```
  {
    "title":"Third Note",
    "note":"My third note."
  }
  ```

* **Success Response:**

  * **Code:** 200 OK <br/>
    **Content:**
    ```
    3
    ```
 
* **Error Response:**

  * **Code:** 400 Bad Request <br/>
    **Content:** `{"message":"add unsuccessful"}`

### Update note (Adds note if note doesn't exist)

Updates an existing note.

* **URL**

  /notes

* **Method:**

  `PUT`
  
*  **URL Params**
  
   None

* **Data Params**

  ```
  {
    "id": 3,
    "title":"Third note edit",
    "note":"My third note edited.",
    "colour":"#0000ff",
    "archived": false
  }
  ```

* **Success Response:**

  * **Code:** 200 OK <br/>
    **Content:**
    ```
    1
    ```
 
* **Error Response:**

  * **Code:** 400 Bad Request <br/>
    **Content:** `{"message":"add unsuccessful"}`
    
### Delete note

Deletes an existing note.

* **URL**

  /notes/:id

* **Method:**

  `PUT`
  
*  **URL Params**
  
   id=[integer]

* **Data Params**

  None

* **Success Response:**

  * **Code:** 204 No Content<br/>
 
* **Error Response:**

  * **Code:** 400 Bad Request <br/>
    **Content:** `{"message:":"delete unsuccessful"}`


