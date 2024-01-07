'use client'
import { css } from '@/styled-system/css'
import { MouseEventHandler } from 'react'
import { BsFillPlayFill } from 'react-icons/bs'

export const PlayStartButton = ({ id }: {id: string}) => {
  const handlePlayStart: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const res = await fetch(`http://localhost:3000/api/programs/${id}/audio`)
    const body = await res.arrayBuffer()
    const audioContext = new AudioContext()
    const audioBuffer = await audioContext.decodeAudioData(body)

    const source = audioContext.createBufferSource();
    source.buffer = audioBuffer
    source.connect(audioContext.destination)
    source.start();
  }

  const styles = {
    main: css({
      background: 'orange.700',
      fontWeight: 'bold',
      color: '#fafafa',
      padding: '3',
      cursor: 'pointer',
      borderRadius: '5px',
      _hover: {
        background: 'orange.600',
      },
    }),
  }

  return (
    <button onClick={handlePlayStart} className={styles.main}><BsFillPlayFill /></button>
  )
}
