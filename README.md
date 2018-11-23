# Note taking app the RESTful way
## Introduction
A [REST](https://en.wikipedia.org/wiki/Representational_state_transfer)ful backend for a note taking app. Postman collection can be found [here](https://www.getpostman.com/collections/fa57fff58077138d4f68).

## API Overiew
| HTTP Request Method  | Resource      | Description                   |
| ---------------------|---------------|-------------------------------|
| GET                  | /notes        | Retrieve all notes            |
| GET                  | /notes/:id    | Retrieve a specific note by ID| 
| POST                 | /notes        | Add new note                  |
| PUT                  | /notes        | Update existing note          |
| DELETE               | /noted/:id    | Delete existing note          |

----
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


