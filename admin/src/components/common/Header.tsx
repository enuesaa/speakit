import { useStyles } from '@/styles/use'
import Link from 'next/link'

export const Header = () => {
  const styles = useStyles(theme => ({
    top: theme().css({
      height: '1.0',
      minHeight: '100px',
      boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
    }),
    title: theme().css({
      color: '#fafafa',
      textShadow: '2px 2px 2px #000',
      fontSize: '45px',
      height: '100px',
      lineHeight: '100px',
      fontWeight: '800',
      margin: '0 auto',
      textAlign: 'center',
      userSelect: 'none',
      '&:hover': {
        textShadow: '3px 3px 2px #000',
      },
    })
  }))

  return (
    <>
      <header css={styles.top}>
        <Link href={{ pathname: `/` }}>
          <div css={styles.title}>my-nextjs-template</div>
        </Link>
      </header>
    </>
  )
}
