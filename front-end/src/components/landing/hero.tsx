import { ChevronRight, Clock } from "lucide-react";
import CalendarAnimation from "./calendarAnimation";

// Componente Hero
const Hero: React.FC = () => {
  return (
    <section id="home" className="pt-32 pb-20 relative overflow-hidden">
      <div className="absolute top-0 right-0 w-full h-full bg-gradient-to-br from-yellow-900/20 via-transparent to-gray-900/10"></div>
      <div className="absolute top-0 left-0 w-full h-full overflow-hidden">
        <div className="absolute top-10 left-10 w-32 h-32 bg-yellow-600/10 rounded-full filter blur-3xl"></div>
        <div className="absolute bottom-10 right-10 w-40 h-40 bg-yellow-500/10 rounded-full filter blur-3xl"></div>
      </div>

      <div className="container mx-auto px-4 relative z-10">
        <div className="flex flex-col md:flex-row items-center justify-between gap-12">
          <div className="md:w-1/2 text-center md:text-left">
            <h1 className="text-4xl md:text-5xl lg:text-6xl font-light leading-tight text-white mb-6">
              Gerencie suas <span className="text-yellow-500 font-medium">datas acadêmicas</span> com facilidade
            </h1>
            <p className="text-gray-300 text-lg md:text-xl mb-8 max-w-lg">
              O UniTime simplifica a organização de provas, eventos e compromissos acadêmicos,
              conectando professores e alunos em um ambiente intuitivo e eficiente.
            </p>
            <div className="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 justify-center md:justify-start">
              <button
                className="bg-gray-600 text-gray-300 px-6 py-3 rounded-lg font-medium transition-all flex items-center justify-center group relative cursor-not-allowed"
                disabled
              >
                <span className="flex items-center">
                  Experimente Agora
                  <ChevronRight size={18} className="ml-2 opacity-50" />
                </span>
                <span className="absolute -top-10 left-0 right-0 bg-gray-800 text-yellow-500 text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity">
                  Em breve disponível
                </span>
              </button>
              <button className="bg-transparent border border-yellow-500/40 hover:border-yellow-500 text-yellow-500 px-6 py-3 rounded-lg font-medium transition-all flex items-center justify-center">
                Saiba Mais
              </button>
            </div>
          </div>

          <div className="md:w-1/2 flex justify-center">
            <div className="relative">
              <CalendarAnimation />
              <div className="absolute -top-4 -right-4 w-12 h-12 bg-yellow-600/20 backdrop-blur-md rounded-full flex items-center justify-center text-yellow-400 shadow-[0_0_15px_rgba(234,179,8,0.3)]">
                <Clock size={24} className="animate-pulse" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Hero;