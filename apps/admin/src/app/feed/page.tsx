'use client'
import { container } from '@/styled-system/patterns'
import { addFeed, type FeedRequestData } from './actions'
import { TextInput } from '@/components/TextInput'
import { useForm } from 'react-hook-form'
import { useTransition } from 'react'

// see https://github.com/react-hook-form/react-hook-form/issues/10391
export default function Page() {
  const { register, handleSubmit, formState: { errors } } = useForm<FeedRequestData>()
  const [isPending, startTransition] = useTransition()

  const onSubmit = handleSubmit(data => {
    startTransition(() => { addFeed(data) })
  })

  return (
    <>
      <div className={container()}>
        <h1>Add Feed</h1>
        {isPending ? 'pending': ''}
        <form onSubmit={onSubmit}>
          <TextInput label='name' regist={register('name')} />
          <TextInput label='url' regist={register('url')} />
          <button type='submit'>submit</button>
        </form>
      </div>
    </>
  )
}
