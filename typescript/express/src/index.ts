import express from "express";
const app: express.Express = express();
const router = express.Router();
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

import { getAllEmployeeHandler, getDepartmentHandler } from "./handlers/user";

app.get("/employees", getAllEmployeeHandler);

// 部署ごとにいる従業員の一覧
app.get("/employees/:department", getDepartmentHandler)

// app.post("/create", createEmployeeHandler);

app.listen(3000, () => {
  console.log("Start on port 3000.");
});
