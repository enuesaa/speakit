'use client'
import { PageTitle } from '../PageTitle'
import { useGetfeeds } from '@/lib/api'
import { FeedCard } from './FeedCard'
import { AddFeedArea } from './AddFeedArea'

export default function Page() {
  const { data, isLoading } = useGetfeeds()

  return (
    <>
      <PageTitle title='Feeds' />
      <AddFeedArea />
      {data?.items?.map((v,i) => (
        <FeedCard key={i} name={v.data?.name ?? ''} url={v.data?.url ?? ''} id={v.id ?? ''} />
      ))}
    </>
  )
}
