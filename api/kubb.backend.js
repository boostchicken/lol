import { defineConfig } from '@kubb/core'
import createSwagger from '@kubb/swagger'
import createSwaggerTS from '@kubb/swagger-ts'
import createSwaggerSWR from '@kubb/swagger-swr'

export default defineConfig({
  root: './passkeys',
  input: {
    path: './backend.yaml',
  },
  output: {
    path: './src',

  },
  hooks: {
    done: ['prettier --write "**/*.{ts,tsx}"'],
  },
  plugins: [
    createSwagger({ output: false }), 
    createSwaggerTS({ output: { path: './models'}}),
    createSwaggerSWR({ output: { path: './clients'}})]
})
