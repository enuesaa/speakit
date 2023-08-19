import { defineConfig } from 'orval'

export default defineConfig({
  speakitapi: {
    output: {
      client: 'react-query',
      target: 'src/lib/api.ts',
    },
    input: '../../openapi.yaml',
  },
})
