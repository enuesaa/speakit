'use client'
import { useDeleteprogramsidHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { MouseEventHandler } from 'react'
import { FaTrash } from 'react-icons/fa'

type Props = {
  id: string;
}
export const DeleteProgramButton = ({ id }: Props) => {
  const deleteProgram = useDeleteprogramsidHook()
  const handleDelete: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    deleteProgram(id)
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