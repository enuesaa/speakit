import { style } from '@vanilla-extract/css'

export default {
  main: style({
    background: '#312e81',
    fontWeight: 'bold',
    color: '#c7d2fe',
    padding: '3',
    cursor: 'pointer',
    borderRadius: '5px',
    ':hover': {
      background: '#4338ca',
    },
  })
}
