import { prisma } from "../handlers/prisma";
import { NextFunction, Request, Response } from "express";
import { validationResult } from "express-validator";
import { getEmployees } from "models/employees";
import { departments } from "const";
import { DepartmentName, GetEmployeesRequest } from "@types";

export async function ok(req: Request, res: Response) {
  return res.status(200).json("ok");
}

export const getAllEmployeeHandler = async (req: Request, res: Response) => {
  const allEmployees = await prisma.employees.findMany();
  console.log(allEmployees);
  res.send(JSON.stringify(allEmployees));
};

export async function getEmployeesController(
  req: GetEmployeesRequest<typeof departments>,
  res: Response,
  next: NextFunction
) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(404).json({ errors: errors.array() });
  }
  const { department } = req.params;
  try {
    //SELECT * FROM dept_emp WHERE dept_no="d001" AND to_date>=now();
    const employees = await getEmployees(department)
    res.status(200).json(employees);
  } catch (e) {
    next(e);
  }
}