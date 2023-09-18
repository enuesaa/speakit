'use client'
import { css } from '@/styled-system/css'
import { FetchFeedButton } from './FetchFeedButton'

type Props = {
  name: string;
  url: string;
  id: string;
}
export const FeedCard = ({ name, url, id }: Props) => {
  const styles = {
    main: css({
      border: 'solid 1px rgba(255,255,255,0.3)',
      padding: '3',
      borderRadius: '2px',
      color: 'indigo.200',
      margin: '10px 0',
    }),
    item: css({
      padding: '1',
      '& b': {
        fontWeight: 'bold',
        textAlign: 'center',
        display: 'inline-block',
        width: '50px',
      },
    }),
  }

  return (
    <div className={styles.main}>
      <div className={styles.item}><b>name</b> {name}</div>
      <div className={styles.item}><b>url</b> {url}</div>
      <FetchFeedButton id={id}/>
    </div>
  )
}