'use client'
import { useState } from 'react'
import {useKey} from 'react-use'

export const KeyConsole = () => {
  const [count, set] = useState(0)
  const increment = () => set(count => ++count)
  useKey('ArrowRight', increment)

  return (
    <div>
      Press arrow up: {count}
    </div>
  )
}
