import { Calendar} from 'lucide-react';
import 'react'

function ComingSoon() {
  return (
    <div className="min-h-screen bg-gray-50 font-sans flex flex-col">
      {/* Header */}
      <header className="bg-white shadow-sm w-full">
        <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-20">
            <div className="flex items-center">
              <a  href="#"className="text-2xl font-light tracking-tight">
                <span className="font-medium">Uni</span>
                <span className="text-yellow-600">Time</span>
              </a>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-grow flex items-center justify-center">
        <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 text-center py-16">
          <div className="mb-8 flex justify-center">
            <div className="w-24 h-24 rounded-full bg-yellow-600 flex items-center justify-center">
              <Calendar size={48} className="text-white" />
            </div>
          </div>
          
          <h1 className="text-4xl sm:text-5xl font-light mb-6">
            <span className="font-medium">Site em </span>
            <span className="text-yellow-600">Construção</span>
          </h1>
          
          <p className="text-gray-600 text-xl mb-10 max-w-xl mx-auto">
            Estamos trabalhando para trazer a melhor experiência de gestão acadêmica para você.
          </p>

          {/* <div className="bg-white p-8 rounded-lg shadow-lg mb-10">
            <h2 className="text-2xl font-medium mb-4">Seja notificado quando lançarmos</h2>
            <p className="text-gray-600 mb-6">
              Deixe seu email para receber atualizações sobre o lançamento do UniTime.
            </p>
            
            <form className="flex flex-col sm:flex-row gap-4 max-w-md mx-auto">
              <input 
                type="email" 
                placeholder="Seu email" 
                className="flex-grow p-3 border border-gray-300 rounded-md focus:ring-2 focus:ring-yellow-500 focus:border-transparent"
              />
              <button 
                type="submit" 
                className="bg-yellow-600 hover:bg-yellow-700 text-white px-6 py-3 rounded-md transition duration-300"
              >
                Notifique-me
              </button>
            </form>
          </div> */}
        </div>
      </main>

      {/* Footer */}
      <footer className="bg-gray-900 text-gray-400 py-8">
        <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col md:flex-row justify-center items-center">
            <div className="mb-4 md:mb-0">
              <div className="text-xl font-light text-white">
                <span className="font-medium">Uni</span>
                <span className="text-yellow-500">Time</span>
              </div>
            </div>
          </div>
          
          <div className="mt-4 text-sm text-center ">
            <p>&copy; {new Date().getFullYear()} UniTime. Todos os direitos reservados.</p>
          </div>
        </div>
      </footer>
    </div>
  )
}

export default ComingSoon;
