'use client'
import { PageTitle } from '@/components/PageTitle'
import { useGetprograms } from '@/lib/api'
import { css } from '@/styled-system/css'
import { PlayStartButton } from './PlayStartButton'
import { ConvertButton } from './ConvertButton'

export default function Page() {
  const { data, isLoading } = useGetprograms()
  const styles = {
    item: css({
      color: 'indigo.200',
      margin: '10px 0',
      '& button': {
        color: '#ff6633',
        margin: '0 10px',
        cursor: 'pointer',
      }
    })
  }

  return (
    <>
      <PageTitle title='Programs' />
      {data?.items?.map((v,i) => (
        <div key={i} className={styles.item}>
          {v.data?.title}
          {v.data?.converted ? (<PlayStartButton id={v.id ?? ''} />) : (<ConvertButton id={v.id ?? ''} />)}
        </div>
      ))}
    </>
  )
}
