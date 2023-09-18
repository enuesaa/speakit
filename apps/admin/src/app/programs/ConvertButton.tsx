'use client'
import { usePostprogramsidconvertHook } from '@/lib/api'
import { MouseEventHandler } from 'react'
import { GiCycle } from 'react-icons/gi'

export const ConvertButton = ({ id }: {id: string}) => {
  const convert = usePostprogramsidconvertHook()

  const handleConvert: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    convert(id, {})
  }

  return (
    <button onClick={handleConvert}><GiCycle /></button>
  )
}
