import UserNavbar from '../Navbar/UserNavbar'
import Head from 'next/head'
import TopNavbar from '../Navbar/TopNavbar'

type LayoutProps = {
    title?: string
    contained?: boolean
    children: React.ReactNode
}

export default function Layout({ title, contained, children }: LayoutProps) {
    return (
        <div className="flex flex-col w-screen h-screen text-gray-700 bg-gray-50">
            <Head>
                <title>{`${title} - Billstore` || 'Billstore'}</title>
            </Head>

            <TopNavbar />
            <UserNavbar />

            <main className={`${contained ? 'container m-auto' : ''} mt-3 flex-1 overflow-auto`}>
                <div>{children}</div>
            </main>
        </div>
    )
}
