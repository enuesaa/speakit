'use client'
import { css } from '@/styled-system/css'
import { container } from '@/styled-system/patterns'
import { addFeed } from './actions'
import { TextInput } from '@/components/TextInput'
import { useForm } from 'react-hook-form'

type FormData = {
  name: string;
  url: string;
}
export default function Page() {
  const styles = {
    main: css({
      fontSize: '7xl',
      color: 'violet.700',
      _hover: {
        color: 'red.300',
      },
    })
  }
  const { register, handleSubmit } = useForm<FormData>()

  return (
    <>
      <div className={container()}>
        <h1 className={styles.main}>
          Add Feed
        </h1>
        <form action={addFeed}>
          {/* <TextInput label='name' regist={register('name')} />
          <TextInput label='url' regist={register('url')} /> */}
          <textarea name='name'></textarea>
          <textarea name='url'></textarea>
          <button type='submit'>submit</button>
        </form>
      </div>
    </>
  )
}
