import { TopLink } from './TopLink'

export default function Page() {
  return (
    <>
      <TopLink href='/feeds' name='Feeds' />
      <TopLink href='/player' name='Player' />
    </>
  )
}
