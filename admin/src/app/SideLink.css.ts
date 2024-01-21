import { style } from '@vanilla-extract/css'

export default {
  main: style({
    background: '#c7d2fe',
    color: '#1e1b4b',
    fontWeight: 'bold',
    display: 'block',
    borderRadius: '5px',
    boxShadow: '0 5px 5px rgba(1,1,1,0.5)',
    padding: '5px 10px',
    margin: '10px 5px',
    width: '100px',
    textAlign: 'center',
    ':hover': {
      background: '#a5b4fc',
    }
  }),
}