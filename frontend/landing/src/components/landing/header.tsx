import { LogIn, Menu, X } from "lucide-react";
import { useEffect, useState } from "react";

// Componente do Cabeçalho
const Header: React.FC = () => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 20);
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  return (
    <header className={`fixed w-full z-50 transition-all duration-300 ${isScrolled ? 'bg-gray-900/90 backdrop-blur-md shadow-lg' : 'bg-transparent'}`}>
      <div className="container mx-auto px-4 py-4">
        <div className="flex justify-between items-center">
          <div className="flex items-center">
            <div className="text-2xl font-light tracking-tight text-white relative">
              <span className="font-medium">Uni</span>
              <span className="text-yellow-600">Time</span>
              <div className="absolute -top-1 -right-3 w-2 h-2 bg-yellow-600 rounded-full animate-pulse"></div>
            </div>
          </div>

          <nav className="hidden md:flex items-center space-x-8">
            <a href="#home" className="text-gray-300 hover:text-yellow-500 transition-colors">Home</a>
            <a href="#features" className="text-gray-300 hover:text-yellow-500 transition-colors">Funcionalidades</a>
            {/* <a href="#testimonials" className="text-gray-300 hover:text-yellow-500 transition-colors">Depoimentos</a> */}
            <a href="#contact" className="text-gray-300 hover:text-yellow-500 transition-colors">Contato</a>
            <button className="bg-gray-600 text-gray-300 px-6 py-3 rounded-lg font-medium transition-all flex items-center justify-center group relative cursor-not-allowed">
              <span className="flex items-center">
                <LogIn size={18} className="mr-2" />
                Login
              </span>
              <span className="absolute top-14 left-0 right-0 bg-gray-800 text-yellow-500 text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity">
                Em breve disponível
              </span>
            </button>
          </nav>

          <div className="md:hidden flex items-center">
            <button
              onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
              className="text-gray-200 focus:outline-none"
            >
              {isMobileMenuOpen ? <X size={24} /> : <Menu size={24} />}
            </button>
          </div>
        </div>

        {/* Mobile Menu */}
        {isMobileMenuOpen && (
          <div className="md:hidden mt-4 bg-gray-800/80 backdrop-blur-md rounded-lg p-4 animate-fadeIn">
            <div className="flex flex-col space-y-4">
              <a href="#home" className="text-gray-300 hover:text-yellow-500 transition-colors py-2">Home</a>
              <a href="#features" className="text-gray-300 hover:text-yellow-500 transition-colors py-2">Funcionalidades</a>
              <a href="#testimonials" className="text-gray-300 hover:text-yellow-500 transition-colors py-2">Depoimentos</a>
              <a href="#contact" className="text-gray-300 hover:text-yellow-500 transition-colors py-2">Contato</a>
              <button className="bg-gray-700 text-yellow-500 px-4 py-2 rounded-md flex items-center justify-center font-medium transition-all border border-yellow-500/30">
                <LogIn size={18} className="mr-2" />
                Login
              </button>
            </div>
          </div>
        )}
      </div>
    </header>
  );
};

export default Header;