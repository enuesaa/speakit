'use client'
import { UseFormRegisterReturn } from 'react-hook-form'

type Props = {
  label: string
  defaultValue?: string
  regist: UseFormRegisterReturn<string>
}
export const TextInput = ({ label, regist, defaultValue }: Props) => {
  return (
    <label>
      {label}
      <input type='text' defaultValue={defaultValue ?? ''} {...regist} />
    </label>
  )
}
