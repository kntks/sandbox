import prisma from "models/client";
import { departments } from "const";
import { DepartmentName, Employee } from "@types";


export async function getEmployees(
  department: DepartmentName<typeof departments>
): Promise<Employee[]> {
  //SELECT * FROM dept_emp WHERE dept_no="d001" AND to_date>=now();
  const employees = await prisma.dept_emp.findMany({
    where: {
      dept_no: departments.find((v) => v.deptName === department)?.deptNo,
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
        },
      },
    },
  });
  return employees.map(({ emp_no, employees }) => ({
    empNo: emp_no,
    firstName: employees.first_name,
    lastName: employees.last_name,
  }));
}
