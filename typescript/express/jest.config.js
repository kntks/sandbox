module.exports = {
  roots: [
    "<rootDir>/src"
  ],
  testMatch: [
    "**/__tests__/*.test.ts"
  ],

  testPathIgnorePatterns: [
    "<rootDir>/node_modules/",
  ],

  // コンパイル対象外のフォルダーを指定
  transformIgnorePatterns: ["/node_modules/"],

  setupFilesAfterEnv: ["<rootDir>/src/__tests__/singleton.ts"],

  moduleDirectories: ["node_modules", "src"],
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
