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
import { param } from "express-validator";

import {
  getAllEmployeeHandler,
  getDepartmentHandler,
} from "./handlers/employees";

app.get("/employees", getAllEmployeeHandler);

const departments = [
  {
    dept_no: "d001",
    dept_name: "marketing",
  },
  {
    dept_no: "d002",
    dept_name: "finance",
  },
  {
    dept_no: "d003",
    dept_name:"human_resources",
  },
  {
    dept_no: "d004",
    dept_name: "production",
  },
  {
    dept_no: "d005",
    dept_name: "development"
  },
  {
    dept_no: "d006",
    dept_name: "quality_management",
  },
  {
    dept_no: "d007",
    dept_name: "sales"
  },
  {
    dept_no: "d008",
    dept_name:  "research",
  },
  {
    dept_no: "d009",
    dept_name: "customer_service"
  }
] as const;

type Department = typeof departments[number]["dept_name"];
const DepartmentNames = departments.map(v => v.dept_name);

// 部署ごとにいる従業員の一覧
app.get(
  "/employees/:department",
  param("department").isIn(DepartmentNames),
  getDepartmentHandler
);

// app.post("/create", createEmployeeHandler);

app.listen(3000, () => {
  console.log("Start on port 3000.");
});
