import Link from 'next/link'
import { SideLink } from './SideLink'
import { ReactNode } from 'react'
import { AiFillHome } from 'react-icons/ai'
import styles from './WithSidebar.css'

const Sidebar = () => {
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
  return (
    <main style={{display: 'flex'}}>
      <Sidebar />
      <section>
        {children}
      </section>
    </main>
  )
}