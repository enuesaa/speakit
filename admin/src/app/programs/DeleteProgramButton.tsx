'use client'
import { useDeleteprogramsidHook } from '@/lib/api'
import { useQueryClient } from '@tanstack/react-query'
import { MouseEventHandler } from 'react'
import { FaTrash } from 'react-icons/fa'
import styles from './DeleteProgramButton.css'

type Props = {
  id: string
}
export const DeleteProgramButton = ({ id }: Props) => {
  const queryClient = useQueryClient()

  const deleteProgram = useDeleteprogramsidHook()
  const handleDelete: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await deleteProgram(id)
    await queryClient.invalidateQueries({ queryKey: ['/programs'] })
  }

  return (
    <button onClick={handleDelete} className={styles.main}>
      <FaTrash />
    </button>
  )
}
