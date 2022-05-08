export const departments = [
  {
    dept_no: "d001",
    dept_name: "marketing",
  },
  {
    dept_no: "d002",
    dept_name: "finance",
  },
  {
    dept_no: "d003",
    dept_name: "human_resources",
  },
  {
    dept_no: "d004",
    dept_name: "production",
  },
  {
    dept_no: "d005",
    dept_name: "development",
  },
  {
    dept_no: "d006",
    dept_name: "quality_management",
  },
  {
    dept_no: "d007",
    dept_name: "sales",
  },
  {
    dept_no: "d008",
    dept_name: "research",
  },
  {
    dept_no: "d009",
    dept_name: "customer_service",
  },
] as const;

type Department = typeof departments[number]["dept_name"];

export const DepartmentNames = departments.map((v) => v.dept_name);
