import { style } from '@vanilla-extract/css'

export default {
  main: style({
    border: 'solid 1px rgba(255,255,255,0.3)',
    padding: '3',
    borderRadius: '2px',
    color: '#c7d2fe',
    margin: '10px 0',
    display: 'flex',
  }),
  left: style({
    flex: '1 1 auto',
  }),
  right: style({
    flex: '0 0 80px',
  }),
  delete: style({
    flex: '0 0 20px',
  }),
}