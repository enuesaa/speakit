import { css } from '@/styled-system/css'
import { container } from '@/styled-system/patterns'
import Link from 'next/link'
import { SideLink } from './SideLink'
import { ReactNode } from 'react'
import { AiFillHome } from 'react-icons/ai'

const Sidebar = () => {
  const styles = {
    main: css({
      height: '100vh',
    }),
    h1: css({
      color: 'indigo.100',
      fontSize: '7xl',
      fontWeight: 'bold',
      display: 'block',
      width: '500px',
      marginTop: '200px',
      '& a': {
        display: 'inline-block',
        color: 'indigo.200',
        fontSize: '3xl',
        margin: '0 30px',
        _hover: {
          color: 'indigo.300',
        },
      }
    }),
    sideLinks: css({
      margin: '100px 0',
    })
  }

  return (
    <nav className={styles.main}>
      <h1 className={styles.h1}>
        speakit
        <Link href='/'><AiFillHome /></Link>
      </h1>
      <div className={styles.sideLinks}>
        <SideLink href='/feeds' name='Feeds' />
        <SideLink href='/programs' name='Programs' />
      </div>
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