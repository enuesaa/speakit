'use client'
import { useDeleteprogramsidHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { useQueryClient } from '@tanstack/react-query'
import { MouseEventHandler } from 'react'
import { FaTrash } from 'react-icons/fa'

type Props = {
  id: string;
}
export const DeleteProgramButton = ({ id }: Props) => {
  const queryClient = useQueryClient()

  const deleteProgram = useDeleteprogramsidHook()
  const handleDelete: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await deleteProgram(id)
    await queryClient.invalidateQueries({queryKey: ['/programs']})
  }

  const styles = {
    main: css({
      fontWeight: 'bold',
      color: 'indigo.500',
      padding: '3',
      cursor: 'pointer',
      _hover: {
        color: 'orange.500',
      },
    }),
  }

  return (
    <button onClick={handleDelete} className={styles.main}><FaTrash /></button>
  )
}