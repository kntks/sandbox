import { prisma } from "./prisma";
import { Request, response, Response } from "express";
import { PrismaClientKnownRequestError, PrismaClientValidationError } from "@prisma/client/runtime";

export const getAllUserHandler = async (req: Request, res: Response) => {
  const allUsers = await prisma.user.findMany();
  console.log(allUsers);
  res.send(JSON.stringify(allUsers));
};

export const createUserHandler = async (req: Request, res: Response) => {
  console.log("req.body", req.body);
  try {
    const user = await prisma.user.create({ data: req.body });
    res.status(200).send(JSON.stringify(user));
  } catch (e) {
    if (e instanceof Error) console.error(e);
    
    // TODO: error typeごとにstatusコードを修正する
    if(e instanceof PrismaClientKnownRequestError && e.code === "P2002") {
      res.status(404).send("this user is already exists");
      return 
    }
    if(e instanceof PrismaClientValidationError) {
      res.status(404).send("There is a validation error");
    }
    res.status(404).send("error");
  }
};
