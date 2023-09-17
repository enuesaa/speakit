'use client'
import { PageTitle } from '@/components/PageTitle'
import { useGetapiprograms } from '@/lib/api'
import { css } from '@/styled-system/css'

export default function Page() {
  const { data, isLoading } = useGetapiprograms()
  const styles = {
    item: css({
      color: 'indigo.200',
      margin: '10px 0',
    })
  }

  return (
    <>
      <PageTitle title='Programs' />
      {data?.items?.map((v,i) => (
        <div key={i} className={styles.item}>{v.data?.title}</div>
      ))}
    </>
  )
}