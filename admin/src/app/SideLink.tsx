import Link from 'next/link'
import styles from './SideLink.css'

type Props = {
  href: string;
  name: string;
}
export const SideLink = ({ href, name }: Props) => {
  return (
    <Link href={href} className={styles.main}>{name}</Link>
  )
}