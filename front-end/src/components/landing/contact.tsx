import { Github, Mail, Phone } from "lucide-react";

// Componente de Contato
const Contact: React.FC = () => {
  return (
    <section id="contact" className="py-20 bg-gray-900/30 backdrop-blur-sm">
      <div className="container mx-auto px-4">
        <div className="text-center mb-16">
          <h2 className="text-3xl md:text-4xl font-light text-white mb-4">Entre em <span className="text-yellow-500 font-medium">contato</span></h2>
          <p className="text-gray-400 max-w-2xl mx-auto">Estamos sempre disponíveis para ajudar. Envie-nos uma mensagem e responderemos o mais breve possível.</p>
        </div>

        <div className="flex flex-col lg:flex-row gap-12">
          <div className="lg:w-1/2">
            <form className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label htmlFor="name" className="block text-gray-300 mb-2">Nome</label>
                  <input
                    type="text"
                    id="name"
                    className="w-full bg-gray-800 border border-gray-700 focus:border-yellow-500 rounded-lg p-3 text-white outline-none transition-colors"
                    placeholder="Seu nome"
                  />
                </div>
                <div>
                  <label htmlFor="email" className="block text-gray-300 mb-2">Email</label>
                  <input
                    type="email"
                    id="email"
                    className="w-full bg-gray-800 border border-gray-700 focus:border-yellow-500 rounded-lg p-3 text-white outline-none transition-colors"
                    placeholder="seu@email.com"
                  />
                </div>
              </div>

              <div>
                <label htmlFor="subject" className="block text-gray-300 mb-2">Assunto</label>
                <input
                  type="text"
                  id="subject"
                  className="w-full bg-gray-800 border border-gray-700 focus:border-yellow-500 rounded-lg p-3 text-white outline-none transition-colors"
                  placeholder="Assunto da mensagem"
                />
              </div>

              <div>
                <label htmlFor="message" className="block text-gray-300 mb-2">Mensagem</label>
                <textarea
                  id="message"
                  rows={6}
                  className="w-full bg-gray-800 border border-gray-700 focus:border-yellow-500 rounded-lg p-3 text-white outline-none transition-colors resize-none"
                  placeholder="Digite sua mensagem aqui..."
                ></textarea>
              </div>

              <button
                type="submit"
                className="bg-gray-600 text-gray-300 px-6 py-3 rounded-lg font-medium transition-all flex items-center justify-center group relative cursor-not-allowed"
              >
                  <span className="absolute top-14 left-0 right-0 bg-gray-800 text-yellow-500 text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity">
                  Em breve disponível
                </span>
                Enviar Mensagem
              </button>
            </form>
          </div>

          <div className="lg:w-1/2">
            <div className="bg-gray-800/50 rounded-xl p-8 border border-yellow-500/10 h-full">
              <h3 className="text-2xl font-light text-white mb-6">Informações de <span className="text-yellow-500 font-medium">Contato</span></h3>

              <div className="space-y-6">
                <div className="flex items-start">
                  <div className="bg-gray-700/50 rounded-full w-12 h-12 flex items-center justify-center mr-4 text-yellow-500">
                    <Mail size={24} />
                  </div>
                  <div>
                    <h4 className="text-white font-medium">Email</h4>
                    <p className="text-gray-400">joaogabrielvianna05@gmail.com</p>
                  </div>
                </div>

                <div className="flex items-start">
                  <div className="bg-gray-700/50 rounded-full w-12 h-12 flex items-center justify-center mr-4 text-yellow-500">
                    <Phone size={24} />
                  </div>
                  <div>
                    <h4 className="text-white font-medium">Telefone</h4>
                    <p className="text-gray-400">+55 (67) 99985-0112</p>
                  </div>
                </div>


              </div>

              <div className="mt-8">
                <h4 className="text-white font-medium mb-4">Siga-nos</h4>
                <div className="flex space-x-4">
                  <a href="https://github.com/JoaoGabrielVianna/uniTime" target="_blank" className="bg-gray-700/50 rounded-full w-10 h-10 flex items-center justify-center text-yellow-500 hover:bg-yellow-600/20 transition-colors">
                    <Github size={20} />
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Contact;