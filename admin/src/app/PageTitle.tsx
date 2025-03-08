import styles from './PageTitle.css'

type Props = {
  title: string
}
export const PageTitle = ({ title }: Props) => {
  return <h3 className={styles.main}>{title}</h3>
}
