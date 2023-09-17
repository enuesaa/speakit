'use client'
import { PageTitle } from '@/components/PageTitle'
import { useGetapifeeds, usePostapifetchHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { MouseEventHandler } from 'react'

export default function Page() {
  const { data, isLoading } = useGetapifeeds()
  const styles = {
    item: css({
      color: 'indigo.200',
      margin: '10px 0',
    })
  }


  return (
    <>
      <PageTitle title='Feeds' />
      {data?.items?.map((v,i) => (
        <div key={i} className={styles.item}>
          {v.data?.name}
          {v.data?.url}
          <FetchFeedButton id={v.id ?? ''}/>
        </div>
      ))}
    </>
  )
}

const FetchFeedButton = ({ id }: { id: string }) => {
  const fetchFeeds = usePostapifetchHook()

  const handleFeedFetch: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    fetchFeeds({ id })
  }

  return (
    <button onClick={handleFeedFetch}>fetch</button>
  )
}
