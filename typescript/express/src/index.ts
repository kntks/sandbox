import express from "express";
const app: express.Express = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// app.use(
//   (req: express.Request, res: express.Response, next: express.NextFunction) => {
//     res.header("Access-Control-Allow-Origin", "*");
//     res.header("Access-Control-Allow-Methods", "*");
//     res.header("Access-Control-Allow-Headers", "*");
//     next();
//   }
// );

import { getAllUserHandler, createUserHandler } from "./handlers/user";

app.get("/users", getAllUserHandler);

app.post("/create", createUserHandler);

app.listen(3000, () => {
  console.log("Start on port 3000.");
});
