import { PrismaClient } from "@prisma/client";

// https://www.prisma.io/docs/concepts/components/prisma-client/working-with-prismaclient/logging
export const prisma = new PrismaClient({
  log: [
    {
      emit: "event",
      level: "query",
    },
    {
      emit: "stdout",
      level: "error",
    },
    {
      emit: "stdout",
      level: "info",
    },
    {
      emit: "stdout",
      level: "warn",
    },
  ],
});
