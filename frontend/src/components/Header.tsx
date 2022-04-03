import styles from '../styles/components/Header.module.css'

const Header = () => {
  return (
    <header className={styles.headerBlock}>
      <h1 className={styles.title}>Kawada Tech Blog</h1>
    </header>
  )
}

export default Header;