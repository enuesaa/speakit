'use client'
import {useKey} from 'react-use'

export const KeyConsole = () => {
  useKey('ArrowRight', () => {
    console.log('ArrowRight Event')
  })
  useKey('ArrowLeft', () => {
    console.log('ArrowLeft Event')
  })

  return <></>
}
