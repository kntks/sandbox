import { Request } from "express";

// https://tsoa-community.github.io/docs/authentication.html

export async function expressAuthentication(
  request: Request,
  securityName: string,
  scopes?: string[]
): Promise<void> {
  console.log("exressAuthentication")
}