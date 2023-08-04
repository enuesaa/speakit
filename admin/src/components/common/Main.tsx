import { useStyles } from '@/styles/use'
import { ReactNode } from 'react'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  const styles = useStyles(theme => ({
    main: theme().css({
      height: '100vh',
    })
  }))

  return <section css={styles.main}>{children}</section>
}
