import { prisma } from "./prisma";
import { Request, Response } from "express";

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
    res.status(404).send("error");
  }
};
