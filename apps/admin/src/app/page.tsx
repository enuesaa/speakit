import { css } from '@/styled-system/css'
import { container } from '@/styled-system/patterns'
import Link from 'next/link'

export default function Page() {
  const styles = {
    main: css({
      fontWeight: 'bold',
    })
  }

  return (
    <>
    <div className={container()}>
      <h1 className={styles.main}>
        speakit admin
      </h1>
      <Link href='/feeds'>Feeds</Link>
    </div>
    </>
  )
}
