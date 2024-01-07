import { css } from '@/styled-system/css'
import Link from 'next/link'

type Props = {
  href: string;
  name: string;
}
export const SideLink = ({ href, name }: Props) => {
  const styles = {
    main: css({
      background: 'indigo.200',
      color: 'indigo.950',
      fontWeight: 'bold',
      display: 'block',
      borderRadius: '5px',
      boxShadow: '0 5px 5px rgba(1,1,1,0.5)',
      padding: '5px 10px',
      margin: '10px 5px',
      width: '100px',
      textAlign: 'center',
      _hover: {
        background: 'indigo.300',
      }
    }),
  }

  return (
    <Link href={href} className={styles.main}>{name}</Link>
  )
}