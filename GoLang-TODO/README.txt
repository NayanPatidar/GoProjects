Breaking down TODO-program in GoLang

func init()  	//It is a piece of code which is ran before any other part in your package.

renderer :- It is a package used for rendering (It has may uses)
- rnd = renderer.New()
  rnd.* will show many use cases like showing HTML page as response when http.StatusOk.

Complete Description of every function

1] CreateTODO

- Create a variable of type todo say 'x' which is th basic layout of your message.
- Decode the message sent (http.Request) and save it in the 'x'.
- Check if the message sent has any title by using x.title. If not then send a error.
- Now because the input is ok create a again a variable of type todo Model say 'xModel'. 
- Now add data to all the fields in the xModel.
- Finally add the data to the collection name and send a Todo created successfully message.

2] Update TODO

- Create a variable id. Extract the value of id from request by trimming the id and assign the value to variable.
- Check whether is the id is correct by bson.IsObjectIdHex
- Create of cariable of type todo.
- Decode the request and assign the data to the variable created. 
- Validate if the title is present or not.
- Update the todo by using db.C(collectionName).update(... changes ...)
- Give a successfully created message.

3] Delete TODO

- Create a variable id and store the value of id by trimming the request made.
- Check whether is the id correct by using bson.IsObjectIdHex 
- Delete the todo by using db.C(collectionName).RemoveId(bson.ObjectidHex(id)) and through the error for it.
- Give a successfully deleted message. 
