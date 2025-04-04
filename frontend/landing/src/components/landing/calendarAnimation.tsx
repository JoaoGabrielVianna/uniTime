import { Calendar } from "lucide-react";
import { useEffect, useState } from "react";

// Componente de Calendário Interativo
const CalendarAnimation: React.FC = () => {
  const [currentDay, setCurrentDay] = useState(new Date().getDate());
  const [dots, setDots] = useState<number[]>([]);
  
  useEffect(() => {
    // Gerar eventos aleatórios para o calendário
    const randomDots = Array.from({length: 8}, () => Math.floor(Math.random() * 31) + 1);
    setDots(randomDots);
    
    // Animar o calendário
    const interval = setInterval(() => {
      setCurrentDay(prev => prev === 31 ? 1 : prev + 1);
    }, 2000);
    
    return () => clearInterval(interval);
  }, []);
  
  const days = Array.from({length: 31}, (_, i) => i + 1);
  
  return (
    <div className="w-full max-w-xs mx-auto bg-gray-800/50 backdrop-blur-md rounded-xl p-4 shadow-[0_0_25px_rgba(234,179,8,0.15)] border border-yellow-600/20">
      <div className="mb-4 flex justify-between items-center">
        <h3 className="text-yellow-500 font-semibold">Março 2025</h3>
        <Calendar size={20} className="text-yellow-500" />
      </div>
      <div className="grid grid-cols-7 gap-1 text-xs text-center mb-2">
        <div className="text-gray-500">D</div>
        <div className="text-gray-500">S</div>
        <div className="text-gray-500">T</div>
        <div className="text-gray-500">Q</div>
        <div className="text-gray-500">Q</div>
        <div className="text-gray-500">S</div>
        <div className="text-gray-500">S</div>
      </div>
      <div className="grid grid-cols-7 gap-1">
        {Array(6).fill(null).map((_, i) => (
          <div key={`empty-${i}`} className="h-8"></div>
        ))}
        {days.map(day => (
          <div 
            key={day} 
            className={`h-8 w-8 flex items-center justify-center rounded-full text-sm ${
              currentDay === day 
                ? 'bg-yellow-600 text-white shadow-[0_0_10px_rgba(234,179,8,0.7)]' 
                : 'text-gray-300 hover:bg-gray-700 cursor-pointer transition-colors'
            } ${dots.includes(day) ? 'relative' : ''}`}
          >
            {day}
            {dots.includes(day) && (
              <div className="absolute -bottom-1 left-1/2 transform -translate-x-1/2 w-1 h-1 bg-yellow-500 rounded-full"></div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};

export default CalendarAnimation;