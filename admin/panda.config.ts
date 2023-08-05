import { defineConfig } from '@pandacss/dev'

export default defineConfig({
  preflight: true,
  include: ['./src/components/**/*.{ts,tsx,js,jsx}', './src/app/**/*.{ts,tsx}'],
  exclude: [],
  theme: {
    extend: {}
  },
  outdir: 'src/styled-system',
})
