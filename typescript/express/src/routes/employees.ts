import { Router } from "express";

import { param } from "express-validator";

import { ok, getDepartmentHandler } from "controllers/employees";
import { DepartmentNames } from "const";
import { errorHandler } from "middlewares/error";

const router = Router();
router.use(errorHandler);

router.get("/", ok);
// 部署ごとにいる従業員の一覧
router.get(
  "/:department",
  param("department").isIn(DepartmentNames),
  getDepartmentHandler
);

export default router;
