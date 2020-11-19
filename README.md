Questions APIs
=================

This project creates apis to manage questions. The endpoints are:
- /question method POST ==> Create a question
- /question/{id} method PUT ==> Update a question
- /question/{id} method DELETE ==> Delete a question
- /question method GET ==> Get all the questions
- /question/user/{user_id} method GET ==> Get all the questions from a specific user

## How to run

The apis are in docker containers and you can use the command bellow to star the services

```sh
$ docker-compose build
$ docker-compose up
```

## Tests
With the service running it is possible to call the apis.

**Create a question**
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"statement": "question 1", "user_id": 1}' \
  http://localhost:8080/question
```

Expected return
```
{
  "ID":1,
  "statement":"question 1",
  "answer":"",
  "user_id":1
}
```

**Update a question**
```
curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{"id": 1, "statement": "updated question 1", "user_id": 1}' \
  http://localhost:8080/question/1
```

Expected return
```
{
  "ID":1,
  "statement":"updated question 1",
  "answer":"",
  "user_id":1
}
```

**Delete a question**
```
curl --request DELETE http://localhost:8080/question/1
```

Expected return
```
"File deleted"
```

**Get all questions**
```
curl --request GET http://localhost:8080/question
```

Expected return
```
[
	{
		"ID":1,
		"statement":"Statement 1",
		"answer":"",
		"user_id":1
	},
	{
		"ID":2,
		"statement":"Statement 2",
		"answer":"",
		"user_id":2
	}
]
```

**Get all questions by user**
```
curl --request GET http://localhost:8080/question/user/1
```

Expected return

```
[
	{
		"ID":1,
		"statement":"Statement 1",
		"answer":"","
		user_id":1
	},
	{
		"ID":2,
		"statement":"Statement 2",
		"answer":"",
		"user_id":1
	}
]
```