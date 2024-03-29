import prisma from "repositories/client";
import { Request, Response } from "express";
import { getEmployees, type Employee } from "repositories/employees";

import {
  Body,
  Controller,
  Get,
  Middlewares,
  Path,
  Post,
  Query,
  Route,
  Security,
  SuccessResponse,
} from "tsoa";
import { validate, validateDepartment, validateOffset } from "middlewares/validation";
import { DepartmentName } from "@types";
import { departments } from "const";

@Route("employees")
export class EmployeesController extends Controller {

  @Middlewares(validate([validateDepartment(), validateOffset()]))
  @SuccessResponse("200", "OK")
  @Get()
  public async get(@Query() offset: string, @Query() department?: string) {
    return await getEmployees(parseInt(offset), department as DepartmentName<typeof departments>);
  }
  
  @Security("accessToken")
  @Security("sessionId")
  @Get("/test")
  async test() {
    return { msg: "OK" }
  }
}
