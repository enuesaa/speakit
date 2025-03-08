'use client'
import { usePostfeedsidfetchHook } from '@/lib/api'
import { MouseEventHandler } from 'react'
import styles from './FetchFeedButton.css'

export const FetchFeedButton = ({ id }: { id: string }) => {
  const fetchFeeds = usePostfeedsidfetchHook()

  const handleFeedFetch: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    fetchFeeds(id, {})
  }

  return (
    <button onClick={handleFeedFetch} className={styles.main}>
      fetch
    </button>
  )
}
