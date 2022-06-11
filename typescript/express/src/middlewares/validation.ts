import { DepartmentNames } from "const";
import { NextFunction, Request, Response } from "express";
import { param, query, validationResult, type ValidationChain } from "express-validator";


export function validate(validations: ValidationChain[]) {
  return async (req: Request, res: Response, next: NextFunction) => {
    await Promise.all(validations.map(validation => validation.run(req)));

    const errors = validationResult(req);
    if (errors.isEmpty()) {
      return next();
    }
    res.status(422).json({ errors: errors.array() });
  }
}

export function validateDepartment(): ValidationChain {
  return query("department").optional().isIn(DepartmentNames).withMessage("department param is invalid value")
}

export function validateOffset(): ValidationChain {
  return query("offset").not().isEmpty().withMessage("offset is required").isInt({gt: 0}).withMessage("offset must be Int and greater than 0")
}