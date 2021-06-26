import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>go-sync</title>
        <meta name="description" content="go-sync" />
        <link rel="icon" href="favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          go-sync
        </h1>
      </main>

      <footer className={styles.footer}>
        <span className={styles.logo}>
          Bruno Shiroma - {}
        </span>
      </footer>
    </div>
  )
}
