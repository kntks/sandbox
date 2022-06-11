import prisma from "models/client";
import { NextFunction, Request, Response } from "express";
import { validationResult } from "express-validator";
import { getEmployees } from "models/employees";
import { departments } from "const";
import { DepartmentName, GetEmployeesRequest } from "@types";

import {
  Body,
  Controller,
  Get,
  Middlewares,
  Path,
  Post,
  Query,
  Route,
  SuccessResponse,
} from "tsoa";
import { validate, validateDepartment } from "middlewares/validation";

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
    const employees = await getEmployees(department);
    res.status(200).json(employees);
  } catch (e) {
    next(e);
  }
}


@Route("employees")
export class EmployeesController extends Controller {
  @Middlewares(validate([validateDepartment()]))
  @SuccessResponse("200", "OK")
  @Get()
  public async get(@Query() offset: number, @Query() department?: string): Promise<{id: string, name: string}[]> {
    return [{
      id: "123",
      name: "hoge"
    }]
  }
}
