'use client'
import { usePostprogramsidconvertHook } from '@/lib/api'
import { useQueryClient } from '@tanstack/react-query'
import { MouseEventHandler } from 'react'
import styles from './ConvertButton.css'

export const ConvertButton = ({ id }: { id: string }) => {
  const convert = usePostprogramsidconvertHook()
  const queryClient = useQueryClient()

  const handleConvert: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await convert(id, {})
    await queryClient.invalidateQueries({ queryKey: ['/programs'] })
  }

  return (
    <button onClick={handleConvert} className={styles.main}>
      convert
    </button>
  )
}
