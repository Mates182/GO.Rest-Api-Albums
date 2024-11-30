# GO.Rest-Api-Albums

This project is a basic RESTful API developed in Go using the **Gin** framework. The API manages information about music albums, allowing CRUD (Create, Read, Update, Delete) operations and partial updates.

## 🚀 Features

- **RESTful routes** for managing albums:
  - Get all albums (`GET /albums`).
  - Get an album by ID (`GET /albums/:id`).
  - Create a new album (`POST /albums`).
  - Fully update an album by ID (`PUT /albums/:id`).
  - Partially update an album by ID (`PATCH /albums/:id`).
  - Delete an album by ID (`DELETE /albums/:id`).
- **Structured JSON responses**.
- Error handling with clear messages.
- **Simple web server** rendering an HTML home page.

## 🛠️ Prerequisites

- Go (version 1.23.3 or later).
- Project dependencies:
  - [Gin](https://github.com/gin-gonic/gin).


## 💻 Installation and Running

1. Clone this repository:

   ```bash
   git clone https://github.com/Mates182/GO.Rest-Api-Albums.git
   ```
   ```
   cd GO.Rest-Api-Albums
   ```

2. Install the required dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. Open your browser and go to [http://localhost:8080](http://localhost:8080).

## 🌐 API Endpoints

### **GET /albums**
Retrieve all albums.

**Response example:**
```json
[
  {
    "id": "1",
    "title": "Born to Die",
    "artist": "Lana del Rey",
    "year": 2012
  },
  ...
]
```

---

### **GET /albums/:id**
Retrieve a specific album by its ID.

**Response example:**
```json
{
  "id": "1",
  "title": "Born to Die",
  "artist": "Lana del Rey",
  "year": 2012
}
```

---

### **POST /albums**
Create a new album.

**Request body (JSON):**
```json
{
  "id": "6",
  "title": "Midnights",
  "artist": "Taylor Swift",
  "year": 2022
}
```

---

### **PUT /albums/:id**
Fully update an existing album.

**Request body (JSON):**
```json
{
  "id": "1",
  "title": "Born to Die - The Paradise Edition",
  "artist": "Lana del Rey",
  "year": 2012
}
```

---

### **PATCH /albums/:id**
Partially update an existing album.

**Request body (JSON):**
```json
{
  "title": "Updated Title"
}
```

---

### **DELETE /albums/:id**
Delete a specific album.

---

## 🧪 Testing the API

You can use tools like [Postman](https://www.postman.com/) or `cURL` to test the endpoints.

### Example using `cURL`:

1. Retrieve all albums:
   ```bash
   curl -X GET http://localhost:8080/albums
   ```

2. Create a new album:
   ```bash
   curl -X POST http://localhost:8080/albums -H "Content-Type: application/json" -d '{"id":"6","title":"New Album","artist":"Midnights","year":2022}'
   ```

3. Update an album:
   ```bash
   curl -X PUT http://localhost:8080/albums/1 -H "Content-Type: application/json" -d '{"id":"1","title":"Born to Die - The Paradise Edition","artist":"Lana del Rey","year":2012}'
   ```

4. Delete an album:
   ```bash
   curl -X DELETE http://localhost:8080/albums/1
   ```

