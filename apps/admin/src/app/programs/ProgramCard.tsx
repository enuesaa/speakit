import { css } from '@/styled-system/css'
import { PlayStartButton } from './PlayStartButton';
import { ConvertButton } from './ConvertButton';

type Props = {
  title: string;
  converted: boolean;
  id: string;
}
export const ProgramCard = ({ title, converted, id }: Props) => {
  const styles = {
    main: css({
      border: 'solid 1px rgba(255,255,255,0.3)',
      padding: '3',
      borderRadius: '2px',
      color: 'indigo.200',
      margin: '10px 0',
      display: 'flex',
    }),
    left: css({
      flex: '1 1 auto',
    }),
    right: css({
      flex: '0 0 100px',
    })
  }

  return (
    <div className={styles.main}>
      <div className={styles.left}>
        {title}
      </div>
      <div className={styles.right}>
        {converted ? (<PlayStartButton id={id} />) : (<ConvertButton id={id} />)}
      </div>
    </div>
  )
}
