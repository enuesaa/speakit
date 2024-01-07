'use client'
import { usePostfeedsidfetchHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { MouseEventHandler } from 'react'

export const FetchFeedButton = ({ id }: { id: string }) => {
  const fetchFeeds = usePostfeedsidfetchHook()

  const handleFeedFetch: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    fetchFeeds(id, {})
  }

  const styles = {
    main: css({
      display: 'inline-block',
      padding: '5px 10px',
      borderRadius: '10px',
      background: 'orange.400',
      cursor: 'pointer',
      margin: '2',
      fontWeight: 'bold',
      _hover: {
        background: 'orange.500',
      },
    }),
  }

  return (
    <button onClick={handleFeedFetch} className={styles.main}>fetch</button>
  )
}
