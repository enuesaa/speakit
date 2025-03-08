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

const sideLinks = style({
  display: 'inline-block',
  margin: '100px 0',
})

globalStyle(`${sideLinks} a`, {
  background: '#c7d2fe',
  color: '#1e1b4b',
  fontWeight: 'bold',
  display: 'block',
  borderRadius: '5px',
  boxShadow: '0 5px 5px rgba(1,1,1,0.5)',
  padding: '5px 10px',
  margin: '10px 5px',
  width: '100px',
  textDecoration: 'none',
  textAlign: 'center',
})
globalStyle(`${sideLinks} a:hover`, {
  background: '#a5b4fc',
})

export default {
  main: style({
    height: '100vh',
    textAlign: 'center',
  }),
  h1,
  sideLinks,
}
