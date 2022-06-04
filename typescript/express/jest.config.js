module.exports = {
  roots: [
    "<rootDir>/src"
  ],
  testMatch: [
    "**/__tests__/**/*.+(ts|tsx|js)",
    "**/?(*.)+(spec|test).+(ts|tsx|js)"
  ],

  testPathIgnorePatterns: [
    "<rootDir>/node_modules/",
  ],

  // コンパイル対象外のフォルダーを指定
  transformIgnorePatterns: ["/node_modules/"],

  transform: {
    ".+\\.(t|j)sx?$": [
      "@swc/jest",
      {
        sourceMaps: true,

        module: {
          type: "commonjs", 
        },

        jsc: {
          parser: {
            syntax: "typescript",
          },
        },
      },
    ],
  },
}
