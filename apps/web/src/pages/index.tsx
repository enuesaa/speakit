import { useEventListener } from 'usehooks-ts'

export default function Page() {

  useEventListener('keydown', (e) => {
    if (e.key === 'ArrowRight') {
      console.log('right')
    }
    if (e.key === 'ArrowLeft') {
      console.log('left')
    }
  });

  return (
    <div style={{ color: '#fafafa' }}>
      a
    </div>
  )
}
