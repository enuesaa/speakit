import styles from './ProgramCard.css'
import { PlayStartButton } from './PlayStartButton'
import { ConvertButton } from './ConvertButton'
import { DeleteProgramButton } from './DeleteProgramButton'

type Props = {
  title: string
  converted: boolean
  id: string
}
export const ProgramCard = ({ title, converted, id }: Props) => {
  return (
    <div className={styles.main}>
      <div className={styles.left}>{title}</div>
      <div className={styles.right}>{converted ? <PlayStartButton id={id} /> : <ConvertButton id={id} />}</div>
      <div className={styles.delete}>
        <DeleteProgramButton id={id} />
      </div>
    </div>
  )
}
