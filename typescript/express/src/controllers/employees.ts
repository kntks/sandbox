import { prisma } from "../handlers/prisma";
import { Request, Response } from "express";
import {
  PrismaClientKnownRequestError,
  PrismaClientValidationError,
} from "@prisma/client/runtime";
import { validationResult } from "express-validator";

import { departments } from "../const";

async function ok(req: Request, res: Response) { 
  return res.status(200).json("ok");
}

export const getAllEmployeeHandler = async (req: Request, res: Response) => {
  const allEmployees = await prisma.employees.findMany();
  console.log(allEmployees);
  res.send(JSON.stringify(allEmployees));
};

async function getDepartmentHandler(req: Request, res: Response) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(404).json({ errors: errors.array() });
  }

  const { department } = req.params
  try {
    //SELECT * FROM dept_emp WHERE dept_no="d001" AND to_date>=now();
    const employee = await prisma.dept_emp.findMany({
      where: {
        dept_no: departments.find((v) => v.dept_name === department)
          ?.dept_no,
        to_date: {
          gte: new Date(),
        },
      },
      select: {
        emp_no: true,
        employees: {
          select: {
            first_name: true,
            last_name: true,
          }
        }
      }
    });
    res.status(200).json(employee);
  } catch (e) {
    if (e instanceof Error) console.error(e);

    // TODO: error typeごとにstatusコードを修正する
    if (e instanceof PrismaClientKnownRequestError && e.code === "P2002") {
      return res.status(404).send("this employee is already exists");
    }
    if (e instanceof PrismaClientValidationError) {
      return res.status(404).send("There is a validation error");
    }
    return res.status(404).send("error");
  }
};


export {
  ok,
  getDepartmentHandler
}