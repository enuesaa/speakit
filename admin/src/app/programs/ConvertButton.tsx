'use client'
import { usePostprogramsidconvertHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { useQueryClient } from '@tanstack/react-query'
import { MouseEventHandler } from 'react'

export const ConvertButton = ({ id }: {id: string}) => {
  const convert = usePostprogramsidconvertHook()
  const queryClient = useQueryClient()

  const handleConvert: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await convert(id, {})
    await queryClient.invalidateQueries({queryKey: ['/programs']})
  }

  const styles = {
    main: css({
      background: 'indigo.900',
      fontWeight: 'bold',
      color: 'indigo.200',
      padding: '3',
      cursor: 'pointer',
      borderRadius: '5px',
      _hover: {
        background: 'indigo.700',
      },
    }),
  }

  return (
    <button onClick={handleConvert} className={styles.main}>convert</button>
  )
}
