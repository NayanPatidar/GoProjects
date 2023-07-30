Breaking down TODO-program in GoLang

func init()  	//It is a piece of code which is ran before any other part in your package.

renderer :- It is a package used for rendering (It has may uses)
- rnd = renderer.New()
  rnd.* will show many use cases like showing HTML page as response when http.StatusOk.

Complete Description of every function

1] CreateTODO

- Create a variable of type todo say 'x' which is th basic layout of your message.
- Decode the message sent (http. Request) and save it in the 'x'.
- Check if the message sent has any title by using x.title. If not then send an error.
- Now because the input is ok create again a variable of type todo Model say 'xModel'. 
- Now add data to all the fields in the xModel.
- Finally add the data to the collection name and send a Todo created successfully message.

2] Update TODO

- Create a variable id. Extract the value of the id from a request by trimming the id and assigning the value to a variable.
- Check whether the id is correct by bson.IsObjectIdHex
- Create of the variable of type todo.
- Decode the request and assign the data to the variable created. 
- Validate if the title is present or not.
- Update the todo by using db.C(collectionName).update(... changes ...)
- Give a successfully created message.

3] Delete TODO

- Create a variable id and store the value of the id by trimming the request made.
- Check whether is the id correct by using bson.IsObjectIdHex 
- Delete the todo by using db.C(collectionName).RemoveId(bson.ObjectidHex(id)) and through the error for it.
- Give a successfully deleted message. 

4] Fetch TODO 

- Create a variable todos and initialize it as a empty slice of todoModel.
- db.C(collectionName).Find(bson.M{}).All(&todos) and throw a error.
- Create a variable todoList and initialize it as an empty slice of todo.
- Using for-each loop append the todoList with todo {... changes ...} 
- Give a message of the data.
