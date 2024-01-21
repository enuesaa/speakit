import { globalStyle, style } from '@vanilla-extract/css'

const h1 = style({
  color: '#e0e7ff',
  fontSize: '7xl',
  fontWeight: 'bold',
  display: 'block',
  width: '500px',
  marginTop: '200px',
})

globalStyle(`${h1} a`, {
  display: 'inline-block',
  color: '#c7d2fe',
  fontSize: '3xl',
  margin: '0 30px',
})

export default {
  main: style({
    height: '100vh',
  }),
  h1,
  sideLinks: style({
    margin: '100px 0',
  })
}