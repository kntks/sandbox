const { build } = require('esbuild')

build({
  // entryPoints,
  entryPoints: ["src/main.ts"],
  outbase: './src', 
  // bundle: true,
  outdir: './lib', 
  platform: 'node',
  external: [],
  watch: false 
})