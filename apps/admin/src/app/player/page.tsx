'use client'
import { container } from '@/styled-system/patterns'
import { KeyConsole } from './KeyConsole'
import { useGetfeeds } from '@/lib/api'

export default function Page() {
  const {data, status} = useGetfeeds()
  console.log(data)
  console.log(status)

  return (
    <>
      <div className={container()}>
        <h1>Player</h1>
        <KeyConsole />
      </div>
    </>
  )
}
