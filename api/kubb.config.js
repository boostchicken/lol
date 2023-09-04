import { defineConfig } from '@kubb/core'
import createSwagger from '@kubb/swagger'
import createSwaggerTS from '@kubb/swagger-ts'
import createSwaggerSWR from '@kubb/swagger-swr'

let skip = [{pattern: "noclient", type: "tag"}];
export default defineConfig({
  root: '.',
  input: {
    path: './openapi.yaml',
  },
  output: {
    path: './src',
    clean: true,
  },
  hooks: {
    done: ['prettier --write "**/*.{ts,tsx}"', 'eslint --fix ./'],
  },
  plugins: [createSwagger({ output: false }), createSwaggerTS({ output: 'models',skipBy: skip }), createSwaggerSWR({output: './hooks',skipBy: skip})],
})
