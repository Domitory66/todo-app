# todo-app
Web application Todo lists for tasks
Learning "Rest-API" using the Gin web framework in language Golang.

Use Cases:          Http method     Path

Registration            POST        /auth/sign-up
Authorization           POST        /auth/sign-in

Show lists              GET         /lists
Go to the list          GET         /lists/{id}
Create list             POST        /lists
Update list             PUT         /lists/{id}
Delete list             DELETE      /lists/{id}

Show tasks              GET         /lists/{id}/items
Add task                POST        /lists/{id}/items