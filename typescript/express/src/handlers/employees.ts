import { prisma } from "./prisma";
import { Request, response, Response } from "express";
import { PrismaClientKnownRequestError, PrismaClientValidationError } from "@prisma/client/runtime";

export const getAllEmployeeHandler = async (req: Request, res: Response) => {
  const allEmployees = await prisma.employees.findMany();
  console.log(allEmployees);
  res.send(JSON.stringify(allEmployees));
};

export const getDepartmentHandler = async (req: Request, res: Response) => {
  console.log("req.body", req.params);
  try {
    //SELECT * FROM dept_emp WHERE dept_no="d001" AND to_date>=now();
    const employee = await prisma.dept_emp.findMany({
      where: {
        dept_no: req.params.dept_no,
        to_date: {
          gte: new Date()
        }
      }
    });
    res.status(200).send(JSON.stringify(employee));
  } catch (e) {
    if (e instanceof Error) console.error(e);
    
    // TODO: error typeごとにstatusコードを修正する
    if(e instanceof PrismaClientKnownRequestError && e.code === "P2002") {
      res.status(404).send("this employee is already exists");
      return 
    }
    if(e instanceof PrismaClientValidationError) {
      res.status(404).send("There is a validation error");
    }
    res.status(404).send("error");
  }
};
