import { prismaMock } from "./singleton";
import { getEmployees } from "models/employees";
import { dept_emp, Prisma } from "@prisma/client";

describe("test employees", () => {
  test("should get empolyees", async () => {
    const employees = [
      {
        employees: {
          first_name: "a",
          last_name: "a",
        },
        emp_no: 1,
      },
    ] as unknown as Prisma.Prisma__dept_empClient<dept_emp[]>;
    prismaMock.dept_emp.findMany.mockResolvedValue(employees);

    await expect(getEmployees("marketing")).resolves.toEqual([
      {
        firstName: "a",
        lastName: "a",
        empNo: 1,
      },
    ]);
  });
});
