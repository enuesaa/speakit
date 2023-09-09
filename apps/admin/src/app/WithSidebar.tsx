import { css } from '@/styled-system/css'
import { container } from '@/styled-system/patterns'
import { ReactNode } from 'react'

const Sidebar = () => {
  const styles = {
    main: css({
      height: '100vh',
      '& h1': {
        color: 'indigo.100',
        fontSize: '7xl',
        fontWeight: 'bold',
        display: 'block',
        width: '500px',
        marginTop: '200px',
      }
    }),
  }

  return (
    <nav className={styles.main}>
      <h1>speakit</h1>
    </nav>
  )
}

type Props = {
  children: ReactNode,
}
export const WithSidebar = ({ children }: Props) => {
  const styles = {
    main: container({
      display: 'flex',
    }),
  }

  return (
    <main className={styles.main}>
      <Sidebar />
      <section>
        {children}
      </section>
    </main>
  )
}