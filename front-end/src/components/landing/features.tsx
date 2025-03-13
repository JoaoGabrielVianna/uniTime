import { Bell, Calendar, Clock, Users } from "lucide-react";

// Tipos
interface Feature {
  id: number;
  icon: React.ReactNode;
  title: string;
  description: string;
}


// Componente de Funcionalidades
const Features: React.FC = () => {
  const features: Feature[] = [
    {
      id: 1,
      icon: <Calendar size={32} className="text-yellow-500" />,
      title: "Cadastro e visualização de eventos",
      description: "Crie, visualize e gerencie datas de provas e eventos acadêmicos de forma centralizada e organizada."
    },
    {
      id: 2,
      icon: <Bell size={32} className="text-yellow-500" />,
      title: "Notificações automáticas",
      description: "Receba lembretes personalizados sobre prazos de entrega, provas e outros eventos importantes."
    },
    {
      id: 3,
      icon: <Users size={32} className="text-yellow-500" />,
      title: "Acesso para professores e alunos",
      description: "Interface adaptada para diferentes perfis, facilitando a comunicação entre docentes e discentes."
    },
    {
      id: 4,
      icon: <Clock size={32} className="text-yellow-500" />,
      title: "Gestão de horários",
      description: "Organize seu tempo acadêmico com facilidade e evite sobreposição de compromissos."
    }
  ];
  
  return (
    <section id="features" className="py-20 bg-gray-900/30 backdrop-blur-sm">
      <div className="container mx-auto px-4">
        <div className="text-center mb-16">
          <h2 className="text-3xl md:text-4xl font-light text-white mb-4">Funcionalidades <span className="text-yellow-500 font-medium">Principais</span></h2>
          <p className="text-gray-400 max-w-2xl mx-auto">Explore as ferramentas que tornam o UniTime a solução ideal para gerenciamento acadêmico.</p>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {features.map((feature) => (
            <div 
              key={feature.id} 
              className="bg-gray-800/50 rounded-xl p-6 border border-yellow-500/10 hover:border-yellow-500/30 transition-all hover:shadow-[0_0_20px_rgba(234,179,8,0.15)] group"
            >
              <div className="mb-4 bg-gray-700/50 rounded-full w-16 h-16 flex items-center justify-center group-hover:bg-yellow-600/20 transition-colors">
                {feature.icon}
              </div>
              <h3 className="text-xl font-medium text-white mb-2">{feature.title}</h3>
              <p className="text-gray-400">{feature.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Features