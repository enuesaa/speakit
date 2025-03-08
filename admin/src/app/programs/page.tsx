'use client'
import { PageTitle } from '../PageTitle'
import { useGetprograms } from '@/lib/api'
import { ProgramCard } from './ProgramCard'

export default function Page() {
  const { data, isLoading } = useGetprograms()

  return (
    <>
      <PageTitle title='Programs' />
      {data?.items?.map((v, i) => (
        <ProgramCard key={i} id={v.id ?? ''} title={v.data?.title ?? ''} converted={v.data?.converted ?? false} />
      ))}
    </>
  )
}
