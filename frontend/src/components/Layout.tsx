import Menu from "./Menu";

const Layout: React.FC = ({ children }) => {
  return (
    <main className='flex'>
      <div className="w-full">
        {children}
      </div>
      <Menu />
    </main>    
  )
}

export default Layout;