import { style } from '@vanilla-extract/css'

export default {
  main: style({
    fontWeight: 'bold',
    color: '#6366f1',
    padding: '3',
    cursor: 'pointer',
    ':hover': {
      color: '#f97316',
    },
  }),
}
