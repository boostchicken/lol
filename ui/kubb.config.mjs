import { defineConfig } from '@kubb/core'
import createSwagger from '@kubb/swagger'
import createSwaggerTS from '@kubb/swagger-ts'
import createSwaggerSWR from '@kubb/swagger-swr'


export default defineConfig({
  root: '.',
  input: {
    path: './public/openapi.yaml',
  },
  output: {
    path: './src/gen',
    clean: true,
  },
  hooks: {
    done: ['prettier --write "**/*.{ts,tsx}"', 'eslint --fix ./src/gen'],
  },
  plugins: [createSwagger({ output: false }), createSwaggerTS({ output: 'models' }), createSwaggerSWR({ output: './hooks',skipBy:[{pattern: "getConfig"},{pattern: "executeLol", type: "operationId"}] })],
})
