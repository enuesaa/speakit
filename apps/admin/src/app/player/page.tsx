import { container } from '@/styled-system/patterns'
import { KeyConsole } from './KeyConsole'

export default async function Page() {
  return (
    <>
      <div className={container()}>
        <h1>Player</h1>
        <KeyConsole />
      </div>
    </>
  )
}
