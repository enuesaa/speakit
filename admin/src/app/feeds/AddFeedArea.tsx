'use client'
import { usePostfeedsHook } from '@/lib/api'
import { FormEventHandler } from 'react'

export const AddFeedArea = () => {
  const addfeed = usePostfeedsHook()

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const name = e.currentTarget.nam.value
    const url = e.currentTarget.url.value

    addfeed({ name, url })
  }

  return (
    <form onSubmit={handleSubmit}>
      <input type='name' name='nam' data-1p-ignore />
      <input type='url' name='url' />
      <button type='submit'>submit</button>
    </form>
  )
}
