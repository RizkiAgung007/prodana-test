export const Footer = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="fixed bottom-0 left-0 w-full bg-white/80 backdrop-blur-md border-t border-slate-200 py-4 z-50">
      <div className="container mx-auto px-4 md:px-8">
        <div className="flex flex-col md:flex-row justify-between items-center gap-4">
          <div className="flex items-center gap-3">
            <div className="w-8 h-8 bg-slate-900 rounded-lg flex items-center justify-center shadow-sm">
              <span className="text-white font-bold text-xs">P</span>
            </div>
            <div className="flex flex-col">
              <span className="text-xs font-bold text-slate-900 leading-tight">
                Prodana System
              </span>
              <span className="text-[10px] text-slate-500 font-medium">
                © {currentYear} Rizki Agung Dermawan
              </span>
            </div>
          </div>

          <div className="hidden sm:flex flex-wrap justify-center gap-2">
            {["Golang", "React", "Gemini AI"].map((tech) => (
              <div
                key={tech}
                className="flex items-center px-2 py-0.5 bg-slate-50 border border-slate-200 rounded-full"
              >
                <div
                  className={`w-1 h-1 rounded-full mr-1.5 ${tech === "Golang" ? "bg-blue-500" : tech === "React" ? "bg-cyan-400" : "bg-indigo-500"}`}
                ></div>
                <span className="text-[9px] font-bold text-slate-600 uppercase tracking-wider">
                  {tech}
                </span>
              </div>
            ))}
          </div>

          <div className="flex items-center gap-2 bg-green-50/50 border border-green-100 px-2 py-1 rounded-md">
            <span className="relative flex h-1.5 w-1.5">
              <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
              <span className="relative inline-flex rounded-full h-1.5 w-1.5 bg-green-500"></span>
            </span>
            <span className="text-[10px] font-semibold text-green-700 uppercase">
              System Active
            </span>
          </div>
        </div>
      </div>
    </footer>
  );
};
