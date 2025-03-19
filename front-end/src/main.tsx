import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import LandingPage from './pages/landingpage.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

// Configuração das rotas
const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />, // Aqui mostra a LandingPage na raiz "/"
  },
]);


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
