'use client'
import { PageTitle } from '@/components/PageTitle'
import { useGetapiprograms, usePostapiprogramsidconvertHook } from '@/lib/api'
import { css } from '@/styled-system/css'
import { MouseEventHandler } from 'react'
import { GiCycle } from 'react-icons/gi'
import { AiFillPlayCircle } from 'react-icons/ai'

export default function Page() {
  const { data, isLoading } = useGetapiprograms()
  const styles = {
    item: css({
      color: 'indigo.200',
      margin: '10px 0',
      '& button': {
        color: '#ff6633',
        margin: '0 10px',
        cursor: 'pointer',
      }
    })
  }

  return (
    <>
      <PageTitle title='Programs' />
      {data?.items?.map((v,i) => (
        <div key={i} className={styles.item}>
          {v.data?.title}
          {v.data?.converted ? (<PlayStartButton id={v.id ?? ''} />) : (<ConvertButton id={v.id ?? ''} />)}
        </div>
      ))}
    </>
  )
}

const ConvertButton = ({ id }: {id: string}) => {
  const convert = usePostapiprogramsidconvertHook()

  const handleConvert: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    convert(id, {})
  }

  return (
    <button onClick={handleConvert}><GiCycle /></button>
  )
}

const PlayStartButton = ({ id }: {id: string}) => {
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
