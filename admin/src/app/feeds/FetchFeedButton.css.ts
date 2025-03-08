import { style } from '@vanilla-extract/css'

export default {
  main: style({
    display: 'inline-block',
    padding: '5px 10px',
    borderRadius: '10px',
    background: '#fb923c',
    cursor: 'pointer',
    margin: '2',
    fontWeight: 'bold',
    ':hover': {
      background: '#f97316',
    },
  }),
}
