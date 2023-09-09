import { css } from '@/styled-system/css'

type Props = {
  title: string;
}
export const PageTitle = ({ title }: Props) => {
  const styles = {
    main: css({
      fontSize: '3xl',
      fontWeight: 'bold',
      color: 'indigo.200',
      margin: '20px 0',
    })
  }

  return (
    <h3 className={styles.main}>
      {title}
    </h3>
  )
}
