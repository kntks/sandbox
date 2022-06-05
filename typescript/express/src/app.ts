import express from "express";
import swaggerJsDoc from "swagger-jsdoc";
import swaggerUi from "swagger-ui-express";

import employees from "routes/employees";

const app: express.Express = express();
if (process.env.NODE_ENV === "dev") {
  const options = {
    swaggerDefinition: {
      info: {
        title: "Express TypeScript",
        version: "1.0.0",
      },
    },
    apis: ["src/routes/*"],
  };
  app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerJsDoc(options)));
}

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
const BASE_URL = "/api/v1";
app.use(BASE_URL, employees);

app.listen(3000, () => {
  console.log("Start on port 3000.");
});
