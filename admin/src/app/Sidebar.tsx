import Link from 'next/link'
import styles from './Sidebar.css'
import { AiFillHome } from 'react-icons/ai'

export const Sidebar = () => {
  return (
    <nav className={styles.main}>
      <h1 className={styles.h1}>
        speakit
        <Link href='/'><AiFillHome /></Link>
      </h1>
      <div className={styles.sideLinks}>
        <Link href='/feeds'>Feeds</Link>
        <Link href='/programs'>Programs</Link>
      </div>
    </nav>
  )
}
