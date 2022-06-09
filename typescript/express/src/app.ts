import express from "express";
import swaggerJsDoc from "swagger-jsdoc";
import swaggerUi from "swagger-ui-express";

import employees from "routes/employees";
import { RegisterRoutes } from "./.build/routes";

const app: express.Express = express();
if (process.env.NODE_ENV === "dev") {
  app.use("/api/docs", swaggerUi.serve, async (req: express.Request, res: express.Response) => {
    return res.send(
      swaggerUi.generateHTML(await import("./.build/swagger.json"))
      );
    }
  )
  // app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerJsDoc(options)));
}

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
RegisterRoutes(app);
const port = process.env.PORT || 3000;
// app.use(
//   (req: express.Request, res: express.Response, next: express.NextFunction) => {
//     res.header("Access-Control-Allow-Origin", "*");
//     res.header("Access-Control-Allow-Methods", "*");
//     res.header("Access-Control-Allow-Headers", "*");
//     next();
//   }
// );
// const BASE_URL = "/api/v1";
// app.use(BASE_URL, employees);

app.listen(port, () => {
  console.log(`Start on port :${port}.`);
});
