const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const packageDefination = protoLoader.loadSync("protobufs/todo.proto", {});
const grpcObj = grpc.loadPackageDefinition(packageDefination);
const todoPackage = grpcObj.todoapp;

const readlineSync = require("readline-sync");

const client = new todoPackage.TodoList("localhost:50051", grpc.credentials.createInsecure())

console.log(`\n
1. GetTodos
2. AddTodos
3. DeleteTodos
`)
var choice = parseInt(readlineSync.question("\nEnter your choice: "));

switch (choice) {
    case 1:

        client.GetTodos({}, (err, res) => {
            console.log(JSON.stringify(res, null, 2))
            console.error(err)
        })

        break;

    case 2:

        client.AddTodo({
            "title": readlineSync.question("Enter the title of the todo: "),
            "description": readlineSync.question("Enter the description of the todo: ")
        }, (err, res) => {
            if (!err) {
                console.log("New todo has been added, Id: " + JSON.stringify(res))
            } else {
                console.error(err)
            }
        })

        break;

    case 3:

        client.DeleteTodo({
            "id": parseInt(readlineSync.question("Enter the Id of the todo to be deleted: "))
        }, (err, response) => {
            if (!err) {
                console.log("Todo has been deleted... " + JSON.stringify(response))
            } else {
                console.error(err)
            }
        })

        break;

    case 4:

        return;

    default:
        console.log("Invalid choice")
        break;
}