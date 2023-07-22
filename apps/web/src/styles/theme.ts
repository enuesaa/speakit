import { type Theme } from '@emotion/react'

export const baseTheme: Theme = {
  surf: {
    main: {
      color: '#fafafa',
      background: '#222222',
    },
    sub: {
      color: '#fafafa',
      background: '#333333',
    },
    reverse: {
      color: '#111111',
      background: '#fafafa',
    },
    transparent: {
      color: 'none',
      background: 'transparent',
    },
  },
  size: {
    x1: {
      fontSize: '16px',
      fontWeight: '400',
      // padding: '2px 5px',
    },
    x2: {
      fontSize: '20px',
      fontWeight: '600',
      // padding: '5px 10px',
    },
    x3: {
      fontSize: '25px',
      fontWeight: '800',
      // padding: '10px 15px',
    },
  },
  decorate: {
    shadow: {
      border: 'none',
      borderRadius: '0',
      boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
    },
    rounded: {
      border: 'none',
      borderRadius: '5px',
      boxShadow: 'none',
    },
    card: {
      border: 'solid 1px rgba(255,255,255,0.2)',
      borderRadius: '5px',
    }
  },
  hover: {
    shadow: {
      '&:hover': {
        boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
      },
    }
  },
  around: {
    x1: {
      padding: '3px',
      margin: '3px',
    },
    x1tb: {
      padding: '3px 0',
      margin: '3px 0',
    },
    x2: {
      padding: '6px',
      margin: '6px',
    },
    x2tb: {
      padding: '6px 0',
      margin: '6px 0',
    },
    x3: {
      padding: '9px',
      margin: '9px',
    },
    x3tb: {
      padding: '9px 0',
      margin: '9px 0',
    },
  },
}
