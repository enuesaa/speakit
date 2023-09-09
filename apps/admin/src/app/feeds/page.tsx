'use client'
import { PageTitle } from '@/components/PageTitle'
import { useGetapifeeds } from '@/lib/api'

export default function Page() {
  const { data, isLoading } = useGetapifeeds()
  console.log(data)

  return (
    <>
      <PageTitle title='feeds' />
    </>
  )
}
