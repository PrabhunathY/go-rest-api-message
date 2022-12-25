# go-rest-api-message

## GET API
Get All Message API: http://localhost:9000/messages
Response: 
[
    {
        "id": 6129484611666145821,
        "title": "Title 1",
        "text": "Text 1"
    }
]

## POST API
POST a Message API: http://localhost:9000/messages
Resuest Body:
    {
        "title": "Title 3",
        "text": "Text 3 desc"
    }
Response: 
{
    "id": 5577006791947779410,
    "title": "Title 3",
    "text": "Text 3 desc"
}