// App.tsx
import React, { useState } from 'react';
import { LogIn, X } from 'lucide-react';
import Header from '../components/landing/header';
import Hero from '../components/landing/hero';
import Features from '../components/landing/features';
// import Testimonials from '../components/landing/testimonials';
import Contact from '../components/landing/contact';
import Footer from '../components/landing/footer';

// Componente de Login Modal
const LoginModal: React.FC<{ isOpen: boolean; onClose: () => void }> = ({ isOpen, onClose }) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/70 backdrop-blur-sm animate-fadeIn">
      <div className="bg-gray-900 rounded-xl p-6 w-full max-w-md border border-yellow-500/20 shadow-[0_0_30px_rgba(234,179,8,0.2)]">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-white">Login</h2>
          <button onClick={onClose} className="text-gray-400 hover:text-white">
            <X size={24} />
          </button>
        </div>

        <form className="space-y-6">
          <div>
            <label htmlFor="email" className="block text-gray-300 mb-2">Email</label>
            <input
              type="email"
              id="email"
              className="w-full bg-gray-800 border border-gray-700 focus:border-yellow-500 rounded-lg p-3 text-white outline-none transition-colors"
              placeholder="seu@email.com"
            />
          </div>

          <div>
            <label htmlFor="password" className="block text-gray-300 mb-2">Senha</label>
            <input
              type="password"
              id="password"
              className="w-full bg-gray-800 border border-gray-700 focus:border-yellow-500 rounded-lg p-3 text-white outline-none transition-colors"
              placeholder="••••••••"
            />
          </div>

          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <input
                type="checkbox"
                id="remember"
                className="w-4 h-4 bg-gray-800 border-gray-700 rounded focus:ring-yellow-500"
              />
              <label htmlFor="remember" className="ml-2 text-sm text-gray-300">Lembrar-me</label>
            </div>
            <a href="#" className="text-sm text-yellow-500 hover:text-yellow-400">Esqueceu a senha?</a>
          </div>

          <button
            type="submit"
            className="w-full bg-yellow-600 hover:bg-yellow-700 text-white py-3 rounded-lg font-medium transition-all shadow-[0_0_15px_rgba(234,179,8,0.3)] hover:shadow-[0_0_20px_rgba(234,179,8,0.4)]"
          >
            Entrar
          </button>

          <div className="text-center text-gray-400 text-sm">
            Não tem uma conta? <a href="#" className="text-yellow-500 hover:text-yellow-400">Cadastre-se</a>
          </div>
        </form>
      </div>
    </div>
  );
};

// App Principal
const LandingPage: React.FC = () => {
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);

  return (
    <div className="min-h-screen bg-gray-900 text-white">
      <Header />
      <main>
        <Hero />
        <Features />
        {/* <Testimonials /> */}
        <Contact />
      </main>
      <Footer />
      <LoginModal isOpen={isLoginModalOpen} onClose={() => setIsLoginModalOpen(false)} />

      {/* Botão de login flutuante */}
      <button
        className="fixed bottom-8 right-8 bg-gray-600 text-gray-300 p-4 rounded-full shadow-lg transition-colors shadow-[0_0_15px_rgba(234,179,8,0.3)] group  cursor-not-allowed z-40"
      // onClick={() => setIsLoginModalOpen(true)}
      >
        <LogIn size={24} />
        <span className="absolute -top-10 -left-20 right-0  bg-gray-800 text-yellow-500 text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity">
          Em breve disponível
        </span>
      </button>

    
    </div>
  );
};

// Exporte o componente App
export default LandingPage;