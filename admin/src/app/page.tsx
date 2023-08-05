import { css } from '@/styled-system/css'
import { container } from '@/styled-system/patterns'
import Link from 'next/link'

export default function Page() {
  const styles = {
    main: css({
      fontSize: '2xl',
      fontWeight: 'bold',
      color: 'violet.700',
      _hover: {
        color: 'red.300',
      },
    })
  }

  return (
    <>
    <div className={container()}>
      <h1 className={styles.main}>
        Hello 🐼!
      </h1>
      <Link href='/about'>about</Link>
    </div>
    </>
  )
}
