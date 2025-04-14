import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './styles/index.css'
import LoginPage from './pages/login'
import { ThemeProvider } from './components/theme-provider'


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider defaultTheme='dark' storageKey='vite-ui-theme'>
      <LoginPage />
    </ThemeProvider>
  </StrictMode>,
)
