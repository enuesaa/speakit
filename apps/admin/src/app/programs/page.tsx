'use client'
import { PageTitle } from '@/components/PageTitle'
import { useGetapiprograms, usePostapiconvertHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { MouseEventHandler } from 'react'
import { GiCycle } from 'react-icons/gi'

export default function Page() {
  const { data, isLoading } = useGetapiprograms()
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
          {v.data?.converted ? (<></>) : (<ConvertButton id={v.id ?? ''} />)}
        </div>
      ))}
    </>
  )
}

const ConvertButton = ({ id }: {id: string}) => {
  const convert = usePostapiconvertHook()

  const handleConvert: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    convert({ id })
  }

  return (
    <button onClick={handleConvert}><GiCycle /></button>
  )
}
