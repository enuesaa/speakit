'use client'
import { MouseEventHandler } from 'react'
import { AiFillPlayCircle } from 'react-icons/ai'

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

  return (<button onClick={handlePlayStart}><AiFillPlayCircle /></button>)
}