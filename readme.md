# My quizz API

### Routes

`/quiz/` -> Random question

Example:
```json
{
    "id":581,
    "question":"Which Italian dessert has a name that can be translated as ",
    "propositions":[
        "Panna cotta",
        "Amatriciana",
        "Tiramisu",
        "Panforte"
    ],
    "réponse":"Panna cotta",
    "anecdote":"The panna cotta is made from cream, milk and sugar, to which is added some Latin."
}
```

`/quiz/[0-899]` -> Fetch question by ID

Example with `/quiz/87`:
```json
{
    "id":87,
    "question":"How many times a second can a hummingbird flap its wings ?",
    "propositions":[
        "78 times",
        "58 times",
        "38 times",
        "18 times"
    ],
    "réponse":"78 times",
    "anecdote":"Hummingbirds are also known as hummingbirds because of their small size and fast wing beats."}
```

