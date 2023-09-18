'use client'
import { usePostprogramsidconvertHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { MouseEventHandler } from 'react'
import { GiCycle } from 'react-icons/gi'

export const ConvertButton = ({ id }: {id: string}) => {
  const convert = usePostprogramsidconvertHook()

  const handleConvert: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    convert(id, {})
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
