export const departments = [
  {
    deptNo: "d001",
    deptName: "marketing",
  },
  {
    deptNo: "d002",
    deptName: "finance",
  },
  {
    deptNo: "d003",
    deptName: "human_resources",
  },
  {
    deptNo: "d004",
    deptName: "production",
  },
  {
    deptNo: "d005",
    deptName: "development",
  },
  {
    deptNo: "d006",
    deptName: "quality_management",
  },
  {
    deptNo: "d007",
    deptName: "sales",
  },
  {
    deptNo: "d008",
    deptName: "research",
  },
  {
    deptNo: "d009",
    deptName: "customer_service",
  },
] as const;

export const DepartmentNames = departments.map(({ deptName }) => deptName);