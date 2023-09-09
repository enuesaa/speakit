import { defineConfig, defineGlobalStyles } from '@pandacss/dev'

const globalCss = defineGlobalStyles({
  'html, body': {
    background: 'indigo.950',
  },
})

export default defineConfig({
  preflight: true,
  include: ['./src/components/**/*.{ts,tsx,js,jsx}', './src/app/**/*.{ts,tsx}'],
  exclude: [],
  theme: {
    extend: {}
  },
  outdir: 'src/styled-system',
  globalCss,
})
