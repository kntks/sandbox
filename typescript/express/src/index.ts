import express from "express";
const app: express.Express = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// app.use(
//   (req: express.Request, res: express.Response, next: express.NextFunction) => {
//     res.header("Access-Control-Allow-Origin", "*");
//     res.header("Access-Control-Allow-Methods", "*");
//     res.header("Access-Control-Allow-Headers", "*");
//     next();
//   }
// );

app.listen(3000, () => {
  console.log("Start on port 3000.");
});

import { PrismaClient } from '@prisma/client';
const prisma = new PrismaClient();

app.get("/users", async (req: express.Request, res: express.Response) => {
  const allUsers = await prisma.user.findMany()
  console.log(allUsers)
  res.send(JSON.stringify(allUsers))
})

app.post("/create", async (req: express.Request, res: express.Response) => {
  console.log("req.body",req.body);
  const user = await prisma.user.create({data: req.body});
  res.send(JSON.stringify(user));
})
