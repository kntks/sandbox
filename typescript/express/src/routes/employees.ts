import { Router } from "express";

import { param } from "express-validator";

import { ok, getEmployeesController } from "controllers/employees";
import { DepartmentNames } from "const";
import { errorHandler } from "middlewares/error";

const router = Router();
router.use(errorHandler);

/**
 * @swagger
 *  definitions:
 *    Employees:
 *      description: 従業員の一覧取得
 *      type: array
 *      items:
 *        type: object
 *        properties:
 *          empNo:
 *            type: number
 *          firstName:
 *            type: string
 *          lastName:
 *            type: string
 */

/**
 * @swagger
 *  parameters:
 *    departmentPathParam:
 *      description: 部署名
 *      in: path
 *      name: department
 *      type: string
 *      required: true
 */

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
 *    get:
 *      description: 部署ごとにいる従業員の一覧取得
 *      tags:
 *        - employees
 *      produces:
 *        - application/json
 *      parameters:
 *        - $ref: "#/parameters/departmentPathParam"
 *      responses:
 *        200:
 *          description: 一覧取得
 *          schema:
 *            $ref: "#/definitions/Employees"
 */
router.get(
  "employees/:department",
  param("department").isIn(DepartmentNames),
  getEmployeesController
);

export default router;
