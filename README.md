# retriever-local

## To run the server
Run :
```
go run .
```
Server will start up on 8080.
 
GET-call on 
```
localhost:8080/events
```
with the following as body:
```
{
    "artist": "Coldplay"
}
```
GET-call on 
```
localhost:8080/events-by-country-date
```
with the following as body:
```
{
    "artist": "Coldplay",
    "country": "IN",
    "start": "XXXX-XX-XX",
    "end": "XXXX-XX-XX"
}
```
GET-call on 
```
localhost:8080/events-by-country
```
with the following as body:
```
{
    "artist": "Coldplay",
    "country": "IN"
}
```
