# Survey

Survey is an API for manage questions that uses Rest and GRPC.


## Prerequisits

* [Docker](https://www.docker.com/get-started) installed in the machine 
* [Docker Compose](https://docs.docker.com/compose/install/) installed in the machine

## Usage

We can obtain, create, edit, delete questions through the api using two protocols, REST or GRPC.
To use REST we can use POSTMAN or any client like this and to use GRPC we can run the ***client.go*** that is in the folder ***pckg/client***

To validate the records created, the questions table of the survey database can be consulted, the connection data can be found in the .env file 

### *Rest*

__Endpoints__


* **GET** all the questions

_Request_
```
curl --location --request GET 'localhost:8080/questions'
```

_Response_


***200***
```
 [
   {
        "id": "1",
        "content": "In what year was america discovered?",
        "description": "About history",
        "user_created": "Luis",
        "answer": "1492"
    },
    {
        "id": "2",
        "content": "Who was Steve Jobs?",
        "description": "About Technology",
        "answer": "The man who created Apple Company"
    }
]
```

* **GET** one question by Id

_Request_
```
curl --location --request GET 'localhost:8080/questions/{id}'
```

_Response_

***200***
```
    {
        "id": "2",
        "content": "Who was Steve Jobs?",
        "description": "About Technology",
        "answer": "The man who created Apple Company"
    }
```

* POST
_Resquest_

**200**
```
curl --location --request POST 'localhost:8080/question' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "Who was Steve Jobs?",
    "description": "About Technology",
    "answer": "The man who created Apple Company"
}'
```

_Response_
```
{
    "id": "9",
    "content": "Who was Steve Jobs?",
    "description": "About Technology",
    "answer": "The man who created Apple Company"
}
```

* PUT

**200**
```
curl --location --request PUT 'localhost:8080/questions/{id}' \
--header 'Content-Type: application/json' \
--data-raw ' {
        "content": "Who was Steve Jobs?",
        "description": "About Technology",
        "answer": "The man who created Apple Company and Pixar"
    }'
```

_Response_
```
{
    "id": "9",
    "content": "Who was Steve Jobs?",
    "description": "About Technology",
    "answer": "The man who created Apple Company and Pixar"
}
```

* DELETE

**200**
```
curl --location --request DELETE 'localhost:8080/questions/{id}'
```

_Response_
```
   "The Question with ID % has been deleted successfully"
```

### *GRPC*

To use GRPC we can run the ***client.go*** that is in the folder ***pckg/client***

* Create, option 1 allows to create a new question, It needs to receive a escaped json parameter with the corresponding structure.

```
go run client.go 1 "{\"content\": \"Who was Steve Jobs?\",\"description\": \"About Technology\",\"answer\": \"The man who created Apple Company and Pixar\"}"
```

* List, option 2 returns all the questions.
```
go run client.go 2
```

* List, option 3 allows to update a question, It needs to receive a escaped json parameter with the corresponding structure.
```
go run client.go 3 "{\"content\": \"Who was Steve Jobs?\",\"description\": \"About Technology\",\"answer\": \"The man who created Apple Company and Pixar\"}"
```

* Delete, option 4 allows to delete a question, It needs to receive the id to delete.
```
go run client.go 4 8
```

* Get by Id, option 5 allows search question by id , It needs to receive the id to seaarch.
```
go run client.go 5 9
```
## Tools 🛠️

* [Golang](https://golang.org/)
* [PostgreSQL](https://www.postgresql.org/) 

