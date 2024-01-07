'use client'
import { PageTitle } from '@/components/PageTitle'
import { useGetfeeds } from '@/lib/api'
import { FeedCard } from './FeedCard'

export default function Page() {
  const { data, isLoading } = useGetfeeds()

  return (
    <>
      <PageTitle title='Feeds' />
      {data?.items?.map((v,i) => (
        <FeedCard key={i} name={v.data?.name ?? ''} url={v.data?.url ?? ''} id={v.id ?? ''} />
      ))}
    </>
  )
}
