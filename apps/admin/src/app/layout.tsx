import { Metadata } from 'next'
import { QueryProvider } from './query'
import './global.css'

type Props = {
  children: React.ReactNode,
}
export default function RootLayout({ children }: Props) {
  return (
    <>
      <html lang='ja'>
        <body>
          <QueryProvider>
            {children}
          </QueryProvider>
        </body>
      </html>
    </>
  )
}
 
export const metadata: Metadata = {
  title: 'my-nextjs-template',
  description: 'my-nextjs-template',
}
