import { globalStyle, style } from '@vanilla-extract/css'

const item = style({
  padding: '1',
})

globalStyle(`${item} b`, {
  fontWeight: 'bold',
  textAlign: 'center',
  display: 'inline-block',
  width: '50px',
})

export default {
  main: style({
    border: 'solid 1px rgba(255,255,255,0.3)',
    padding: '3',
    borderRadius: '2px',
    color: '#c7d2fe',
    margin: '10px 0',
  }),
  item,
}
