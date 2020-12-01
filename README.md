# Library API with GOLang
## Running the code:  
 ```bash
go run cmd\library\librarysrv.go ./config/config.yaml
```
## Endpoints :
#### - GET: 
``` bash 
http://localhost:8080/books
```
#### - GET: 
``` bash 
http://localhost:8080/book/:id
```
#### - POST: 
``` bash 
http://localhost:8080/book
```
##### Request body: 
 ``` bash 
    {
        "title": "string",
        "author": "string",
        "price": number
    }
```

#### - PUT: 
``` bash 
http://localhost:8080/book/:id
```

##### Request body: 
 ``` bash 
    {
        "title": "string",
        "author": "string",
        "price": number
    }
```
#### - DELETE: 
``` bash 
http://localhost:8080/book/:id
```
