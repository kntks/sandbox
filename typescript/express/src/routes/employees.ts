import { Router } from "express";

import { param } from "express-validator";

import { ok, getDepartmentHandler } from "controllers/employees";
import { DepartmentNames } from "const";
import { errorHandler } from "middlewares/error";

const router = Router();
router.use(errorHandler);

/**
 * @swagger
 * /:
 *   get:
 *     description: 部署ごとにいる従業員の一覧取得
 *     produces:
 *       - application/json
 *     responses:
 *       200:
 *         description: タイトル
 */
router.get("/", ok);

/**
 * @swagger
 * /employees/{department}:
 *   get:
 *     description: 部署ごとにいる従業員の一覧取得
 *     produces:
 *       - application/json
 *     responses:
 *       200:
 *         description: タイトル
 */
router.get(
  "/:department",
  param("department").isIn(DepartmentNames),
  getDepartmentHandler
);

export default router;
