import {
  PrismaClientKnownRequestError,
  PrismaClientValidationError,
} from "@prisma/client/runtime";
import { NextFunction, Request, Response } from "express";

const BAD_REQUEST = "BAD REQUEST"
const INTERNAL_SERVER_ERROR = "INTERNAL SERVER ERROR"
export function errorHandler(
  err: Error,
  req: Request,
  res: Response,
  next: NextFunction
) {
  if (err instanceof Error) console.error(err);

  // TODO: error typeごとにstatusコードを修正する
  if (err instanceof PrismaClientKnownRequestError && err.code === "P2002") {
    return res.status(400).send({error: BAD_REQUEST, message: "this employee is already exists"});
  }
  if (err instanceof PrismaClientValidationError) {
    return res.status(400).send({error: BAD_REQUEST, message: "There is a validation error"});
  }
  return res.status(500).send({erorr: INTERNAL_SERVER_ERROR, message: "error"});
}
