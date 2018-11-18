# Note taking app the RESTful way
## Introduction
A note taking app back end the [REST](https://en.wikipedia.org/wiki/Representational_state_transfer)ful way. A Postman collection can be found [here](https://www.getpostman.com/collections/fa57fff58077138d4f68)

## API Overiew
| HTTP Request Method  | Resource      | Description                   |
| ---------------------|---------------|-------------------------------|
| GET                  | /notes        | Retrieve all notes            |
| GET                  | /notes/{id}   | Retrieve a specific note by ID| 
| POST                 | /notes        | Add new note                  |
| PUT                  | /notes        | Update existing note          |
| DELETE               | /noted/{id}   | Delete existing note          |

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
    **Content:** `{"message":"Empty set"}`
