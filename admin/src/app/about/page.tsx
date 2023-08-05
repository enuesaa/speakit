import { css } from '@/styled-system/css'
import { container } from '@/styled-system/patterns'
import { addNote } from './actions'

export default async function Page() {
  const styles = {
    main: css({
      fontSize: '7xl',
      color: 'violet.700',
      _hover: {
        color: 'red.300',
      },
    })
  }

  return (
    <>
      <div className={container()}>
        <h1 className={styles.main}>
          This is about page.
        </h1>
        <form action={addNote}>
          <textarea name='name'></textarea>
          <button type='submit'>submit</button>
        </form>
      </div>
    </>
  )
}

// https://zenn.dev/sumiren/articles/664c86a28ec573
// nextjs 13 では fetch が cache されるらしく、そのまま Page() の中で fetch() すれば良いみたい
// export function generateStaticParams() {
//   return [
//     { slug: 'a' },
//   ]
// }
