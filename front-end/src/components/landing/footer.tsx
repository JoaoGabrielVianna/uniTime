import { Github } from "lucide-react";

// Componente Footer
const Footer: React.FC = () => {
  return (
    <footer className="bg-gray-900 text-gray-400 py-12">
      <div className="container mx-auto px-4">
        <div className="flex flex-col md:flex-row justify-between items-center">
          <div className="mb-6 md:mb-0">
            <div className="text-2xl font-light tracking-tight text-white relative">
              <span className="font-medium">Uni</span>
              <span className="text-yellow-600">Time</span>
              <div className="absolute -top-1 -right-3 w-2 h-2 bg-yellow-600 rounded-full"></div>
            </div>
            <p className="mt-2 max-w-xs">Simplificando a gestão de tempo acadêmico para instituições de ensino, professores e alunos.</p>
          </div>

          {/* <div className="grid grid-cols-2 md:grid-cols-3 gap-8 text-center md:text-left">
            <div>
              <h4 className="text-white font-medium mb-3">Produto</h4>
              <ul className="space-y-2">
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Recursos</a></li>
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Preços</a></li>
                <li><a href="#" className="hover:text-yellow-500 transition-colors">FAQ</a></li>
              </ul>
            </div>

            <div>
              <h4 className="text-white font-medium mb-3">Empresa</h4>
              <ul className="space-y-2">
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Sobre</a></li>
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Blog</a></li>
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Carreiras</a></li>
              </ul>
            </div>

            <div>
              <h4 className="text-white font-medium mb-3">Legal</h4>
              <ul className="space-y-2">
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Termos</a></li>
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Privacidade</a></li>
                <li><a href="#" className="hover:text-yellow-500 transition-colors">Cookies</a></li>
              </ul>
            </div>
          </div> */}
        </div>

        <div className="border-t border-gray-800 mt-12 pt-8 flex flex-col md:flex-row justify-between items-center">
          <p>&copy; 2025 UniTime. Todos os direitos reservados.</p>
          <div className="flex space-x-4 mt-4 md:mt-0">
            <a href="https://github.com/JoaoGabrielVianna/uniTime" target="_blank" className="hover:text-yellow-500 transition-colors">
              <Github size={18} />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;