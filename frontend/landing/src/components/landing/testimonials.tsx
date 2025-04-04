import { useEffect, useState } from "react";

interface Testimonial {
  id: number;
  quote: string;
  name: string;
  role: string;
  initial: string;
}

// Componente de Depoimentos
const Testimonials: React.FC = () => {
  const testimonials: Testimonial[] = [
    {
      id: 1,
      quote: "O UniTime transformou completamente a forma como organizo minhas avaliações. Agora tenho tudo centralizado e não perco mais nenhum prazo importante.",
      name: "Maria Silva",
      role: "Estudante de Engenharia",
      initial: "M"
    },
    {
      id: 2,
      quote: "Como professor, o UniTime me ajuda a planejar melhor meu calendário de aulas e provas, evitando sobrecarregar os alunos com entregas simultâneas.",
      name: "Prof. Carlos Oliveira",
      role: "Docente de Medicina",
      initial: "C"
    },
    {
      id: 3,
      quote: "A interface intuitiva e as notificações foram essenciais para melhorar meu desempenho acadêmico. Recomendo para todos os universitários.",
      name: "João Pereira",
      role: "Estudante de Direito",
      initial: "J"
    }
  ];
  
  const [activeTestimonial, setActiveTestimonial] = useState(0);
  
  useEffect(() => {
    const interval = setInterval(() => {
      setActiveTestimonial((prev) => (prev + 1) % testimonials.length);
    }, 5000);
    
    return () => clearInterval(interval);
  }, [testimonials.length]);
  
  return (
    <section id="testimonials" className="py-20 relative overflow-hidden">
      <div className="absolute top-0 left-0 w-full h-full overflow-hidden">
        <div className="absolute top-20 left-20 w-64 h-64 bg-yellow-600/5 rounded-full filter blur-3xl"></div>
        <div className="absolute bottom-20 right-20 w-64 h-64 bg-yellow-500/5 rounded-full filter blur-3xl"></div>
      </div>
      
      <div className="container mx-auto px-4 relative z-10">
        <div className="text-center mb-16">
          <h2 className="text-3xl md:text-4xl font-light text-white mb-4">O que dizem nossos <span className="text-yellow-500 font-medium">usuários</span></h2>
          <p className="text-gray-400 max-w-2xl mx-auto">Veja como o UniTime tem ajudado estudantes e professores a otimizar seu tempo acadêmico.</p>
        </div>
        
        <div className="max-w-4xl mx-auto">
          <div className="relative">
            {testimonials.map((testimonial, index) => (
              <div 
                key={testimonial.id}
                className={`bg-gray-800/50 rounded-xl p-8 border border-yellow-500/10 shadow-[0_0_20px_rgba(234,179,8,0.1)] transition-all duration-500 ${
                  index === activeTestimonial ? 'opacity-100 translate-y-0' : 'opacity-0 absolute top-0 left-0 translate-y-8'
                }`}
              >
                <div className="flex items-start mb-6">
                  <div className="bg-yellow-600 w-12 h-12 rounded-full flex items-center justify-center text-white font-bold text-xl mr-4">
                    {testimonial.initial}
                  </div>
                  <div>
                    <h4 className="text-white font-medium text-lg">{testimonial.name}</h4>
                    <p className="text-gray-400 text-sm">{testimonial.role}</p>
                  </div>
                </div>
                <p className="text-gray-300 italic">{testimonial.quote}</p>
              </div>
            ))}
          </div>
          
          <div className="flex justify-center mt-8">
            {testimonials.map((_, index) => (
              <button
                key={index}
                onClick={() => setActiveTestimonial(index)}
                className={`w-3 h-3 rounded-full mx-1 transition-all ${
                  index === activeTestimonial ? 'bg-yellow-500 w-8' : 'bg-gray-600 hover:bg-gray-500'
                }`}
                aria-label={`Depoimento ${index + 1}`}
              />
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default Testimonials;