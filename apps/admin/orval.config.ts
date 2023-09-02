import { defineConfig } from 'orval'

export default defineConfig({
  speakitapi: {
    output: {
      client: 'react-query',
      schemas: 'src/lib/schema',
      target: 'src/lib/api.ts',
      override: {
        mutator: {
          path: './src/lib/client.ts',
          name: 'useClient',
        },
      },
    },
    input: './openapi.yaml',
  },
})
