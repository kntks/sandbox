import {
  PrismaClientKnownRequestError,
  PrismaClientValidationError,
} from "@prisma/client/runtime";
import { NextFunction, Request, Response } from "express";

export function errorHandler(
  err: Error,
  req: Request,
  res: Response,
  next: NextFunction
) {
  if (err instanceof Error) console.error(err);

  // TODO: error typeごとにstatusコードを修正する
  if (err instanceof PrismaClientKnownRequestError && err.code === "P2002") {
    return res.status(404).send("this employee is already exists");
  }
  if (err instanceof PrismaClientValidationError) {
    return res.status(404).send("There is a validation error");
  }
  return res.status(404).send("error");
}
