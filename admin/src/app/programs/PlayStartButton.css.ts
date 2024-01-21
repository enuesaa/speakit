import { style } from '@vanilla-extract/css'

export default {
  main: style({
    background: '#c2410c',
    fontWeight: 'bold',
    color: '#fafafa',
    padding: '3',
    cursor: 'pointer',
    borderRadius: '5px',
    ':hover': {
      background: '#ea580c',
    },
  })
}
