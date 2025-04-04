import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import LandingPage from './pages/landingpage.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import PageTest from './pages/page_test.tsx'

// Configuração das rotas
const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />, // Aqui mostra a LandingPage na raiz "/"
  },
  {
    path: '/teste',
    element: <PageTest/>
  }
]);


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
