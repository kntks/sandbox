// const DepartmentNames = departments.map((v) => v.dept_name);

import { Request } from "express";

type GetEmployeesRequest<T> = Request<
  GetEmployeesRequestParam<T>,
  unknown,
  unknown,
  Record<string, any> | undefined
>;

type GetEmployeesRequestParam<T> = {
  department: DepartmentName<T>;
};

type Department = {
  no: string;
  name: string;
};
type DepartmentName<T extends ReadonlyArray<{ deptName: string }>> =
  T[number]["deptName"];

type Employee = {
  empNo: number;
  firstName: string;
  lastName: string;
};
