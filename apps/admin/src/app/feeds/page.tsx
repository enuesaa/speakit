'use client'
import { useGetapifeeds } from '@/lib/api'
import { container } from '@/styled-system/patterns'

export default function Page() {
  const { data, isLoading } = useGetapifeeds()
  console.log(data)

  return (
    <>
      <div className={container()}>
        <h1>List Feeds</h1>
        {/* <form onSubmit={onSubmit}>
          <TextInput label='name' regist={register('name')} />
          <TextInput label='url' regist={register('url')} />
          <button type='submit'>submit</button>
        </form> */}
      </div>
    </>
  )
}
