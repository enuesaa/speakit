import { useStyles } from '@/styles/use'

type Props = {
  name: string
}
export const Title = ({ name }: Props) => {
  const styles = useStyles(theme => ({
    h2: theme({ surf: 'main', size: 'x2' })
  }))

  return <h2 css={styles.h2}>{name}</h2>
}
