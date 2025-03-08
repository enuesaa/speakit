'use client'
import { FetchFeedButton } from './FetchFeedButton'
import styles from './FeedCard.css'

type Props = {
  name: string
  url: string
  id: string
}
export const FeedCard = ({ name, url, id }: Props) => {
  return (
    <div className={styles.main}>
      <div className={styles.item}>
        <b>name</b> {name}
      </div>
      <div className={styles.item}>
        <b>url</b> {url}
      </div>
      <FetchFeedButton id={id} />
    </div>
  )
}
